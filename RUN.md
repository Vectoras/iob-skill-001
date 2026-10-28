# Run 013: backend only

- **Date**: 2026-10-28. **Base**: an empty directory; images built by this run (the build cache was cleared before the round's first run). **Skill version**: `run/013` (run 012's version plus fixes 79–85).
- **Inputs**: project `hello-013`; two Go apps, no SvelteKit app: `catalog`, a public API on 8080, and `worker`, an internal service on 8081 (published on `127.0.0.1` only). Repo URL the default: no `origin`, so `https://github.com/Vectoras/hello-013`.
- **Versions found**: Node 24.22.1, pnpm 12.15.0 (12.16.0 and 12.16.1 held back by the 3-day rule), Debian trixie, Go 1.27.3, golangci-lint 2.15.0.
- **Not in this branch**: `CLAUDE.md` and `.claude/skills/`, which the creation writes from its templates.

## Steps

One commit per step that writes files. Steps 1–3 write nothing; steps 7 and 8 don't apply to this shape; steps 9 and 10 changed nothing, so they have no commit.

| Step | Commit message | Check |
|---|---|---|
| 4 | root files and images | pass |
| 5 | root dependencies and the hooks | pass |
| 6 | catalog (with the Go side) | pass |
| 6 | worker (with `go-types`) | pass |

## Checklist

Run during the creation (step 10), then graded again afterwards (full mode included).

| # | Check | Result |
|---|---|---|
| 1 | Stack starts lean | pass |
| 2 | Every program answers from the host | pass; `worker` on `127.0.0.1` only |
| 3 | Each container runs only its program | pass |
| 4 | Env vars reach the real processes | skipped: no SvelteKit app |
| 5 | One pnpm version | pass |
| 6 | No reinstall between containers | pass |
| 7 | Generator watcher once per edit | skipped: no `shared-types` |
| 8 | (merged into 15) | — |
| 9 | Unit tests | pass, uncached |
| 10 | End-to-end | skipped: no SvelteKit app |
| 11 | Cached outputs restored; generated code comes first | skipped: no SvelteKit app, no `shared-types` |
| 12 | Git ignores what is generated | pass |
| 13 | Go layout | pass |
| 14 | Full mode, and back to lean | pass |
| 15 | Hooks | pass |
| 16 | OpenAPI descriptions | pass, but the check's own command is wrong: `head -1` expects the `openapi:` line first, and the generator writes the keys in alphabetical order (`components:` first) |

## Verdict

**Pass**: the project needed no fixes and no step needed a workaround. Check 16's command needs fixing in the skill.
