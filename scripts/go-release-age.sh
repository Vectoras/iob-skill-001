#!/usr/bin/env bash
# Prints, for each Go module given, the newest release at least 3 days old: the
# version to pass to `go get <module>@<version>` (or `go get -tool <package>@<version>`
# for a tool in the module). Changes nothing.
#
# The age is the time the Go module proxy first saw the version, read from the module
# index (index.golang.org), which the proxy sets and nobody can change. A version's
# `.info` time is its commit's, which the committer sets. The index can't be asked
# about one module, so the script reads every version published in the last 3 days
# (about 150 pages, a couple of minutes) and holds back the given modules' versions
# among them. Give every module at once: the scan is the same for one or ten.
#
# Run it inside the editing container, from anywhere in the repo:
#   ./scripts/go-release-age.sh github.com/go-chi/chi/v5 github.com/danielgtaylor/huma/v2

set -euo pipefail

if [[ $# -eq 0 ]]; then
  echo "Usage: $0 <module path>..." >&2
  exit 1
fi

min_age_days=3
cutoff=$(date -u -d "${min_age_days} days ago" +%Y-%m-%dT%H:%M:%SZ)
young=$(mktemp)
trap 'rm -f "$young"' EXIT

# The proxy's URLs write an uppercase letter as "!" and its lowercase.
escape() { sed -E 's/[A-Z]/!\L&/g' <<<"$1"; }

# Release versions only (no prerelease), newest first.
declare -A versions
for module in "$@"; do
  list=$(curl -fsS "https://proxy.golang.org/$(escape "$module")/@v/list" | grep -vE -- '-' | sort -rV || true)
  if [[ -z $list ]]; then
    echo "${module}: no release found on proxy.golang.org" >&2
    exit 1
  fi
  versions[$module]=$list
done

# The index lists the versions in the order the proxy first saw them, at most 2000 per
# page; a page's last timestamp is where the next one starts. A short page is the end.
patterns=$(printf '"Path":"%s",\n' "$@")
since=$cutoff
pages=0
while :; do
  page=$(curl -fsS "https://index.golang.org/index?since=${since}&limit=2000")
  pages=$((pages + 1))
  grep -F -- "$patterns" <<<"$page" >>"$young" || true
  [[ $(wc -l <<<"$page") -lt 2000 ]] && break
  since=$(tail -n 1 <<<"$page" | sed -E 's/.*"Timestamp":"([^"]+)".*/\1/')
done
echo "Read ${pages} pages of the module index (since ${cutoff})." >&2

for module in "$@"; do
  held=()
  chosen=""
  while read -r version; do
    if grep -qF -- "\"Path\":\"${module}\",\"Version\":\"${version}\"" "$young"; then
      held+=("$version")
    else
      chosen=$version
      break
    fi
  done <<<"${versions[$module]}"
  if [[ -z $chosen ]]; then
    echo "${module}: every release is younger than ${min_age_days} days" >&2
    exit 1
  fi
  if [[ ${#held[@]} -gt 0 ]]; then
    echo "${module}: held back ${held[*]} (younger than ${min_age_days} days)" >&2
  fi
  echo "${module} ${chosen}"
done
