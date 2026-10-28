# hello-015

A monorepo in pnpm workspaces and Turborepo, developed inside containers. Its apps and packages are listed below.

## Structure

| Folder           | What it holds                                                           |
| ---------------- | ----------------------------------------------------------------------- |
| `scripts/`       | Maintenance scripts (versions)                                          |
| `docs/guide.md`  | The rules this project follows                                          |
| `apps/accounts/` | Go private API (chi + huma)                                             |
| `tools/`         | Go dev tools (wgo, Delve, tygo): a module with no code, outside go.work |

## Prerequisites

- Docker with Compose. Nothing else: every toolchain lives in the images, and development happens inside the containers.
- VS Code, attached to the running editing container (`./code-attach.sh`, which needs `code` on the host).
- Files created by the containers belong to user id 1000 on the host. On Linux, work as that user.

Only want to try it, without Docker? See [Quick run without Docker](#quick-run-without-docker).

## Getting started

A fresh clone first needs each app's secrets file (gitignored), from its committed example:

```bash
for f in apps/*/.env.example; do cp "$f" "${f%.example}"; done
```

A secret shared between apps has the same name in each of their files (a private API's `<API>_SECRET`, in the API and in every frontend app) and must have the same value; replace `change-me` with a random value (`head -c 32 /dev/urandom | base64`) if the stack is reachable by anyone else. Then:

```bash
docker compose up -d
```

This builds the images and starts every service. The first start installs everything and takes a few minutes.

### Services

| Service       | Image            | What it does                                                                                                      |
| ------------- | ---------------- | ----------------------------------------------------------------------------------------------------------------- |
| `editing`     | `Dockerfile.all` | Runs nothing. The editor attaches here; git, the hooks, checks, tests and builds run here                         |
| `maintenance` | `Dockerfile.all` | Installs dependencies and what the apps' tools need, generates code, then regenerates it on change                |
| `accounts`    | `Dockerfile.go`  | The Go private API on port 8080, rebuilt and restarted on change; every route but `/healthz` needs `X-Api-Secret` |

The app services wait for `maintenance` to be healthy. The whole repo is bind-mounted into every container at `/monorepo`. Compose starts each program directly, with no `pnpm run` in between; Turborepo runs the tasks (checks, tests, builds, code generation), not the apps.

### Lean and full mode

```bash
docker compose up -d                                          # lean: apps only
docker compose -f compose.yaml -f compose.full.yaml up -d     # full: apps with their debuggers on
```

In full mode each JS app listens for a debugger on port 9229 and each Go app on 2345, inside the Compose network only. Attach from the editor (attached to `editing`) with the "(Docker)" configuration in each app's `.vscode/launch.json`.

### Working in the repo

Attach the editor (`./code-attach.sh`, or `./code-attach.sh apps/<app>` for one folder), then in its terminal:

```bash
pnpm turbo run lint:check format:check typecheck   # what the pre-push hook runs
pnpm turbo run test:unit                           # unit and story tests
pnpm turbo run test:e2e                            # end-to-end, against the running stack
pnpm turbo run build --filter <app>
```

Commits are made in the editing service, where the hooks run: pre-commit fixes the staged files, pre-push checks (no tests: those are for CI).

### OpenAPI descriptions

`pnpm openapi` writes each API's OpenAPI description to `apps/<api>/generated/openapi.yaml` (OpenAPI 3.1) and `openapi-3.0.yaml` (3.0), for Postman and other API clients. They are gitignored and built from the code: run it again after a change. A private API's description asks for its `X-Api-Secret`.

### Changing how an app starts

Each app's start command is written twice: in `compose.yaml` (its full-mode version in `compose.full.yaml`), and as the app's `dev:lean` / `dev:full` scripts, which the host run uses. Each Storybook's command likewise, as its package's `storybook:dev`. **When you change one, change the other in the same commit.** The intended differences: in the full-mode scripts the debuggers listen on `127.0.0.1`, the JS ones on a random port; a Go app's scripts set its `PORT` themselves.

### Env vars and secrets

Compose injects every variable. Non-secret values are in `compose.yaml`; secrets in each app's gitignored `.env`, listed with dummy values in its committed `.env.example`. There is no root `.env`.

### Go dev tools

The Go dev tools (wgo, Delve, and tygo once `packages/shared-types` exists) are listed in `tools/go.mod`, a module with no code, kept out of `go.work` so the tools' dependencies never change the apps' versions. `maintenance` installs them into `$GOPATH/bin`, and everything calls them by name. Every command in `tools/` needs `GOWORK=off`.

### Go modules in sync

Inside `go.work`, Go builds every module with the highest version any module asks for; a module built on its own (a production image) uses only its own `go.mod`. `./scripts/check-go-sync.sh` reports any module that would build differently on its own, with the fix (`go work sync`, then `GOWORK=off go mod tidy` in each module it names). It also reports a local module a `go.mod` replaces but its `package.json` doesn't list as a workspace dependency: Turborepo reads only `package.json`, and would replay the module's cached checks and tests after a change in the module it uses. It changes nothing, and runs in the pre-push hook.

### Tool versions

Every toolchain version is written once, in the build arguments at the top of `compose.yaml`. Change one there by hand, rebuild (`docker compose up -d --build`), then run the matching script inside `editing`:

| Script                             | What it does                                                                                  |
| ---------------------------------- | --------------------------------------------------------------------------------------------- |
| `./scripts/check-versions.sh`      | Changes nothing. Reports everything out of step with `compose.yaml` and which script fixes it |
| `./scripts/update-pnpm-version.sh` | Sets the pnpm range in the root `package.json` to a new major; then restart `maintenance`     |
| `./scripts/update-go-version.sh`   | Sets the `go` line of `go.work` and every `go.mod`                                            |

## Hello world

Each piece of wiring has a "Hello world" that proves it works: a route, a type, a schema, a page, a component, a test, each with a comment saying so (`grep -rl 'Hello world' apps packages e2e`). Delete them when the real work starts.

## Not set up yet

Chosen per project when the need comes (see "Not decided" in `docs/guide.md`):

- Persistence (PostgreSQL is the database; migrations, query layer and the rest are open)
- Logging and error response shapes
- Authentication and sessions
- CI (the rules for it are in `docs/guide.md` §9)
- Where the Go apps are hosted in production
- Toolchain versions: one source of truth with one bump command (for now each tool has its own `scripts/update-*.sh`, and `scripts/check-versions.sh` checks they agree)

## Quick run without Docker

For a quick look only, not for development. You need Node and pnpm installed (and Go, with a Go app); their versions are not checked (the containers' are at the top of `compose.yaml`).

```bash
pnpm install
(cd tools && GOWORK=off go install tool)    # the Go dev tools
export PATH="$PATH:$(go env GOPATH)/bin"    # if Go's bin folder isn't on your PATH yet
for f in apps/*/.env.example; do cp -n "$f" "${f%.example}"; done
set -a; for f in apps/*/.env; do . "$f"; done; set +a    # every app's secrets
pnpm dev:lean     # or: pnpm dev:full, with the debuggers on
```

One command generates code (and regenerates it on change), then starts every app. A Storybook runs on its own: `pnpm turbo run storybook:dev --filter <package>`. In full mode the debuggers listen on `127.0.0.1` only; attach with the "(host)" configurations.

Don't switch between this and the containers in the same clone: both use the same `node_modules`, and each side reinstalls it for itself.
