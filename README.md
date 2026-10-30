# hello-020

A monorepo in pnpm workspaces and Turborepo, developed inside containers. Its apps and packages are listed below.

## Structure

| Folder            | What it holds                                                                                                            |
| ----------------- | ------------------------------------------------------------------------------------------------------------------------ |
| `scripts/`        | Maintenance scripts (versions)                                                                                           |
| `docs/guide.md`   | The rules this project follows                                                                                           |
| `.claude/skills/` | Project skills for Claude Code: `add-package` (adding an app or package), `runbook` (day-to-day operations of the stack) |
| `.vscode/`        | Editor settings for the attached container (Prettier as the formatter)                                                   |

## Prerequisites

- Docker with Compose. Nothing else: every toolchain lives in the images, and development happens inside the containers.
- VS Code, attached to the running editing container (`./code-attach.sh`, which needs `code` on the host).
- Files created by the containers belong to user id 1000 on the host. On Linux, work as that user.

Only want to try it, without Docker? See [Quick run without Docker](#quick-run-without-docker).

## Getting started

```bash
docker compose up -d
```

This builds the images and starts every service. The first start installs everything and takes a few minutes. Each later start still contacts the npm registry, to check the lockfile against the 3-day release rule (about 20 seconds at most, with nothing downloaded).

### Services

| Service       | Image            | What it does                                                                                       |
| ------------- | ---------------- | -------------------------------------------------------------------------------------------------- |
| `editing`     | `Dockerfile.all` | Runs nothing. The editor attaches here; git, the hooks, checks, tests and builds run here          |
| `maintenance` | `Dockerfile.all` | Installs dependencies and what the apps' tools need, generates code, then regenerates it on change |

The app services wait for `maintenance` to be healthy. The whole repo is bind-mounted into every container at `/monorepo`. Compose starts each program directly, with no `pnpm run` in between; Turborepo runs the tasks (checks, tests, builds, code generation), not the apps.

### Lean and full mode

```bash
docker compose up -d                                          # lean: apps only
docker compose -f compose.yaml -f compose.full.yaml up -d     # full: apps with their debuggers on
```

In full mode each app listens for a debugger, inside the Compose network only; attach from the editor (attached to `editing`) with the "(Docker)" configuration in each app's `.vscode/launch.json`. The ports:

### Working in the repo

Attach the editor (`./code-attach.sh`, or `./code-attach.sh apps/<app>` for one folder), then in its terminal:

```bash
pnpm turbo run lint:check format:check typecheck   # what the pre-push hook runs
pnpm turbo run test:unit                           # unit tests
pnpm turbo run build --filter <app>
```

Commits are made in the editing service, where the hooks run: pre-commit fixes the staged files, pre-push checks (no tests: those are for CI).

### Changing how an app starts

Each app's start command is written twice: in `compose.yaml` (its full-mode version in `compose.full.yaml`), and as the app's `dev:lean` / `dev:full` scripts, which the host run uses. **When you change one, change the other in the same commit.** The intended differences:

- in the full-mode scripts, the debuggers listen on `127.0.0.1`

### Env vars and secrets

Compose injects every variable. Non-secret values are in `compose.yaml`; secrets in each app's gitignored `.env`, listed with dummy values in its committed `.env.example`. There is no root `.env`.

### Tool versions

Every toolchain version is written once, in the build arguments at the top of `compose.yaml`. Change one there by hand, rebuild (`docker compose up -d --build`), then run the matching script inside `editing`:

| Script                             | What it does                                                                                  |
| ---------------------------------- | --------------------------------------------------------------------------------------------- |
| `./scripts/check-versions.sh`      | Changes nothing. Reports everything out of step with `compose.yaml` and which script fixes it |
| `./scripts/update-pnpm-version.sh` | Sets the pnpm range in the root `package.json` to a new major; then restart `maintenance`     |

## Hello world

Each piece of wiring has a "Hello world" that proves it works, each with a comment saying so (`git grep -l 'Hello world'`). Delete them when the real work starts.

## Not set up yet

Chosen per project when the need comes (see "Not decided" in `docs/guide.md`):

- Persistence (PostgreSQL is the database; migrations, query layer and the rest are open)
- Logging and error response shapes
- Authentication and sessions
- CI (the rules for it are in `docs/guide.md` §9)
- Toolchain versions: one source of truth with one bump command (for now each tool has its own `scripts/update-*.sh`, and `scripts/check-versions.sh` checks they agree)

## Quick run without Docker

For a quick look only, not for development. You need Node and pnpm installed (and Go, with a Go app); their versions are not checked (the containers' are at the top of `compose.yaml`).

```bash
pnpm install
pnpm dev:lean     # or: pnpm dev:full, with the debuggers on
```

One command generates code (and regenerates it on change), then starts every app. In full mode the debuggers listen on `127.0.0.1` only; attach with the "(host)" configurations.

Don't switch between this and the containers in the same clone: both use the same `node_modules`, and each side reinstalls it for itself.
