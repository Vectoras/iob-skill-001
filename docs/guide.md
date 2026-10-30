# Guide

Which technology is used for which purpose, and the conventions that connect them. No reasoning here: the why is in `reference.md`.

Marks used: _(not tried)_ means the rule was decided without hands-on experience; _(checked by Claude)_ means Claude verified it, which is evidence, not experience. **Not decided** means it is chosen per project; the list is at the end.

## 0. Stance

- A **monorepo** is the centre of every project: **pnpm workspaces** + **Turborepo**.
- The **frontend** uses the JS ecosystem, in **TypeScript, never plain JavaScript**.
- The **backend** uses the Go ecosystem.
- **Code quality** tooling is configured once at the root. **Lefthook** gives fast feedback locally; **CI is the gate**.
- **Type safety between frontend and backend** is required.
- **Development always happens inside containers.** Work is done in one editing container; each app container runs only its app. A quick run on the host stays possible for someone who only wants to try the project, but it is not a way to develop (§3).
- **Versions**: latest stable at creation (for the JS runtime, the newest LTS), then pinned to exact versions. Updates are deliberate, as their own change.

## 1. Workspace & orchestration

| Purpose                           | Technology / convention                                                    |
| --------------------------------- | -------------------------------------------------------------------------- |
| Package management, local linking | pnpm workspaces, `workspace:*`; nothing hoisted to the root (`publicHoistPattern` at its default, `[]`) |
| Task running, ordering, caching   | Turborepo; `$schema` in `turbo.json` points at the installed package's schema |
| Apps                              | one or more frontend apps, one or more Go apps, or both; each ecosystem's tooling (image, configs, scripts, hooks) is set up with its first app, so a project holds only the ecosystems it uses |
| Layout                            | `apps/` (deployables), `packages/` (shared libraries), `e2e/` (end-to-end test packages); the only three workspace globs; plus `tools/`, the Go dev tools module, outside the workspace |
| Package naming                    | unscoped; a scope only once packages are numerous enough to group          |
| Shared packages                   | export their source (types, generated code, stylesheets included) and are **not built**; only apps build |
| Go apps/packages in the workspace | in the same `apps/` and `packages/` dirs; own `go.mod` + a manifest-only `package.json` whose scripts call Go, and whose only dependencies are the local Go modules it requires, as workspace packages (so the task runner sees them); the Go sync check fails on one missing _(checked by Claude; not tried)_ |
| Linking local Go modules          | `go.work` at the root + `require`/`replace` in each module; module path mirrors the repo URL |
| Dependency versions (JS)          | exact versions, written once in the pnpm catalog (`catalogMode: strict`, `saveExact: true`); packages refer to it with `catalog:`; hand edits to a manifest are not caught; a scaffolder's versions are moved into the catalog as exact versions, within the scaffolder's own ranges (`catalogPrune: true`, then remove, add with the scaffolder's range, and make exact) _(checked by Claude; not tried)_; workspace dependencies (`workspace:*`) are written into the manifest by hand |
| Release age                       | nothing younger than 3 days is installed (`minimumReleaseAge: 4320` in the workspace file); the same for toolchains (JS runtime, package manager, Go, Go linter), checked by hand, and for Go modules, by the time the Go module proxy first saw the version (`scripts/go-release-age.sh`, from the module index) |
| Dependency versions (Go)          | `go.mod` + committed `go.sum`; after any dependency change `go work sync`, then `GOWORK=off go mod tidy` in each changed module (the sync skips `go.sum`); a read-only script in the pre-push hook fails on any module that would build differently alone; CI builds and tests each module with `GOWORK=off` _(checked by Claude; not tried)_ |
| Lockfile                          | committed; a plain install never changes it (frozen in the maintenance service and in CI) |

**Script names**, the same in every package, in both ecosystems:

- `build`, `typecheck`, `generate`
- `dev:lean` / `dev:full`: in every app, for a run on the host; the same commands Compose runs in lean and full mode (§3)
- `storybook:dev` / `storybook:build`: in every package with a Storybook, shared packages included. Storybook is not an app, so it is outside the `dev:*` naming; `storybook:dev` is the same command its Compose service runs
- `lint:fix` / `lint:check`, `format:fix` / `format:check`
- `test:unit` / `test:e2e`
- `openapi`: in every API, writes its OpenAPI description (§5)
- A script that splits is named `<category>:<variant>`. There is no bare `dev`, `lint`, `format` or `test`.
- A package defines only the scripts that apply to it. Go packages have no `typecheck`; shared packages have no `dev:*`.
- The root `dev:lean` / `dev:full` run `turbo watch generate dev:<mode>`: the generator's watcher and every app, in one command.
- The root `package.json` scripts only delegate to Turborepo.

**Maintenance scripts** (whole-repo chores, not package work):

- In `scripts/` at the root; not a workspace package. Work belonging to one package is a `package.json` script there, run through Turborepo.
- Called by path (`./scripts/<name>.sh`), never through root `package.json` entries.
- Bash, with `node` one-liners for JSON. They stay small glue: they exist because versions have no single source of truth, not to hold logic.
- `check-*` is read-only: exits 0 when all is fine and 1 when not, prints `ok` / `DIFF` per item and the fixing command, and leaves the tree exactly as it found it. `update-*` writes, never automatically.
- Run inside the editing container, from anywhere in the repo; each has a header comment (what, exit codes, how to run).
- A `check-*` that is fast and depends only on the repo's files may run in the pre-push hook; one that depends on the local environment runs by hand.

**Task rules**:

- Tasks are independent unless one really reads another's output. No `^build` by default.
- The one standing dependency: tasks that read generated code depend on `generate` (§6).
- Read-only tasks (checks, tests, builds) are cached. Fixers are not.
- **In the containers, Turborepo starts no app**: Compose does (§3). `dev:*` are Turborepo tasks only for the host run: persistent, uncached, depending on `^generate`.
- `storybook:dev` is the same kind of task, but the root `dev:*` don't start it: on the host it is started by hand, per package. `storybook:build` is cached, depends on `^generate`, and declares its output folder.
- A task that writes files declares them as outputs.
- A task that reads config from outside its package lists those root files as inputs.
- Every env var an app reads is declared in the task config.

## 2. Code quality

| Purpose              | JS / TS / Svelte                                   | Go                                         |
| -------------------- | -------------------------------------------------- | ------------------------------------------ |
| Formatting           | Prettier: defaults + the framework's plugin        | golangci-lint running `gofmt`              |
| Linting              | ESLint (flat config): recommended presets, plus a tool's own plugin where its initializer adds one; formatter-compat config last, not type-aware | golangci-lint: the standard linter set, as a starting point; parallel runners allowed |
| Type checking        | tsc, svelte-check                                  | the Go compiler (in `build`)               |
| Type-checking config | one base `tsconfig` at the root, `strict`, extended by every package | n/a                      |

- **Config lives at the repo root only. Scripts live in every package** and run through Turborepo.
- Root files (configs, README) belong to no package, so they get their own root-level format scripts.
- Lint and format take their ignores from the root `.gitignore`.
- A Go package's format scripts also run Prettier on its non-Go files.

**Git hooks (Lefthook)**, installed as a root dev dependency (its install script must be allowed):

- **Pre-commit**: fix the staged files in sequence (lint fix, then format), and re-stage them. A linter that works on whole packages runs once per package with staged files.
- **Pre-push**: verify through Turborepo: lint, format and type check, for packages and root files; plus the read-only Go sync check (§1). **No tests**; they run in CI (§9).

**Repo hygiene**:

- One `.gitignore`, at the root. Per-package ones from scaffolding are merged into it and deleted.
- `.editorconfig` at the root, agreeing with the formatters.
- LF line endings on every OS, through `.gitattributes`.

## 3. Development environment

**Three kinds of service**, all in one Compose file:

| Kind            | What it does                                                                                       | Image            |
| --------------- | -------------------------------------------------------------------------------------------------- | ---------------- |
| **Editing**     | one service; the editor attaches to it, git and the hooks run in it; runs nothing else             | the all image    |
| **Maintenance** | one service: install, Go tools (from the tools module, with a Go app), the test browser (with a frontend app), first `generate`, ready marker, then `pnpm turbo watch generate` | the all image    |
| **App**         | one per app, plus one per Storybook (which has a service, though it is not an app); runs that program and nothing else | its ecosystem's image |

| Purpose                    | Technology / convention                                                        |
| -------------------------- | ------------------------------------------------------------------------------ |
| Containerised dev          | Docker + Compose; no container for libraries                                   |
| Images                     | one per ecosystem the project uses (the JS runtime's slim official image; Go's official image) plus one **all image** with every toolchain the project uses (the JS runtime's official image, always, since the workspace runs on it; Go and the Go linter added with the first Go app; the test browser's system libraries with the first frontend app, picked by a pinned version of the browser test tool, changed with the OS release, not kept in step with the catalog) |
| Image files                | `Dockerfile.<name>` at the repo root (they sort together)                     |
| Operating system           | the same release in every image: Debian, the newest release common to every base image and supported by the browser test tool; written once as a build argument, named in every tag |
| User                       | non-root, the same user id in every image; tools outside the mounted path      |
| Dependency install         | only in the maintenance service; app services wait until it is healthy         |
| Starting an app            | from Compose, calling the app's program directly, with no wrapper (no shell, no package manager, no `go tool`): JS programs from the app's `node_modules/.bin`, Go tools from `$GOPATH/bin` |
| Run on the host            | kept possible for a quick try, not for development: install by hand (packages and the Go tools, with Go's bin folder on the path), then the root `dev:lean` or `dev:full`; nothing is pinned there |
| Go reload                  | wgo builds a binary into a folder inside the app and runs it; that folder is excluded from the watch |
| Modes                      | **lean** (the Compose file alone) and **full** (plus an override file that switches every app to its debug command; Storybook's command stays the same); debuggers are opt-in |
| Debugging                  | from the editing service, by service name; debugger ports not published        |
| Process init               | `init: true` on every service                                                  |
| Source                     | the whole repo bind-mounted into every container, at the same path             |
| Caches                     | `.cache/` at the repo root, gitignored: `pnpm-store`, `go`, `go-build`, `golangci-lint`, `playwright` (the test browser); no named volumes |
| Network                    | a fixed port per app; every dev server listens on all interfaces; an internal service's port is published on loopback only (§5) |
| Toolchain versions         | exact versions written once in the Compose file as build arguments with no defaults; the **argument names are a contract**; the browser test tool's version there only picks the all image's system libraries |
| Go toolchain               | `GOTOOLCHAIN=local` in every image with Go                                     |
| Version scripts            | a read-only report of every version written in more than one place, against the Compose file, run by hand; one update script per tool; scripts never write the Compose file and change other places only through the tools' own commands (maintenance scripts, §1) |
| Package manager            | installed directly in the images that run it, not through Corepack; also named in the root `package.json`, which the task runner requires, with a range limited to its major version _(checked by Claude)_ and set to **fail** when its version doesn't satisfy that range; the exact version is only in the Compose file |
| Package manager settings   | every setting it fills in differently with and without `CI` is written down in the workspace file |
| Editing                    | the editor attaches to the running editing service through a helper script (project root by default, other folders relative to it); Compose alone starts and stops services; no `devcontainer.json` |
| Editor files in git        | extension recommendations, launch, task and settings files                     |
| Push access                | SSH agent forwarding                                                           |
| End-to-end tests           | run from the editing service with `test:e2e`, against the running stack by service name; no service or image of their own _(not tried)_ |

**Watched tasks** (the maintenance service's watch mode):

- A watched task writes only into its own output folder, apart from the sources, and git ignores that folder.
- Never watch a task that edits sources.
- Check every new watched task once by hand: one edit gives exactly one run, and idle gives none.
- If a watched task gets `inputs`, the list starts with `$TURBO_DEFAULT$`.

**Start commands, kept in step**:

- Each app's start command is written twice: in the Compose files (lean, and full in the override file) and as the app's `dev:lean` / `dev:full` scripts. Each Storybook's command likewise: in the Compose file and as its package's `storybook:dev`.
- **A change in one place is made in the other, in the same commit.** A comment beside the commands in the Compose files says so.
- Intended differences: in the full-mode scripts the debuggers listen on loopback only, and JS debuggers on a random port (port `0`): on one host a fixed port clashes between JS apps, and all interfaces would open the debuggers to the local network. A Go app's scripts set its port themselves (`PORT=<port>`), which Compose sets in the service's environment: on one host the Go apps would otherwise share one `PORT`.
- Each app's editor launch file has a host configuration beside the container one: JS apps attach by picking the process, Go apps on loopback.

**Env vars and secrets**:

- Compose injects the variables; apps never load `.env` files themselves. Exception: a framework or tool that reads them by default (SvelteKit in dev and at build) is left to do so, not turned off.
- Every process that runs an app's code (its dev server, Storybook, tests) gets that app's variables.
- Non-secret values are written in the Compose file.
- Secrets go in a gitignored `.env` beside the app, with a committed `.env.example`. No root `.env`.
- What is a secret in production is a secret in dev too.

## 4. Frontend

| Purpose                  | Technology / convention                                                          |
| ------------------------ | -------------------------------------------------------------------------------- |
| Web app                  | SvelteKit (Svelte, Vite)                                                         |
| Adapter                  | follows the deployment target, set at creation; the auto adapter only while the target is unknown |
| Rendering                | server-side by default; prerender every page that needs no API data; browser-only needs a reason |
| Server code              | web-standard features only, no Node-only modules; heavier work belongs in a Go app |
| Calling a private API    | always through the frontend's server, never from the browser; no CORS, no proxy; every frontend app calls every private API (§5) |
| Page data                | a server `load`                                                                  |
| User actions             | form actions; frontend endpoints only for what doesn't fit a form                |
| Live updates             | a ticketed direct channel that carries signals, not data; the browser re-fetches through the frontend's server _(not tried)_ |
| Env vars                 | the framework's env modules; private by default, read at runtime by default      |
| UI components            | chosen per project, at creation: a shared UI package in `packages/`, which owns its styling setup, or directly in the frontend app |
| How components are built | chosen per project, after creation: hand-written, a library, or a combination; own components use plain CSS, no Tailwind |
| Component workshop       | Storybook, one per app/library                                                   |
| Sessions                 | the frontend's server holds the session and passes the user's identity to the private APIs; the rest of authentication is **not decided** |

## 5. Backend

| Purpose                  | Technology / convention                                                          |
| ------------------------ | -------------------------------------------------------------------------------- |
| Routing, middleware      | chi                                                                              |
| Validation, OpenAPI docs | huma (chi adapter)                                                               |
| API description files    | every API (private and public) writes its OpenAPI description from the code, with no server started, by its `openapi` script (root: `pnpm openapi`, every API): `generated/openapi.yaml` (OpenAPI 3.1) and `generated/openapi-3.0.yaml` (3.0), gitignored; a separate script, not part of `generate`; a private API's description declares its secret header |
| Service layout           | `cmd/<name>/main.go` from the start; everything else in `internal/`, no `pkg/`   |
| Inside `internal/`       | start flat; then one package per feature with the layers as files, plus a few technical packages |
| Go app roles             | each Go app is one of: a **private API**, called only by the frontend apps' servers, which add a shared secret in its own header (`X-Api-Secret`) on every call but the health check, never in `Authorization`; a **public API**, with no secret (CORS, rate limits and authentication **not decided**); an **internal service**, called only by other Go apps, its port published in dev on loopback only (how it is called is **not decided**). One Go app is a private API when the project has a frontend app, a public API otherwise _(not tried)_ |
| Shared Go code           | the types packages when the apps call for them (§6); any other shared module only on second use, under `packages/` |
| Config                   | one struct filled once at startup with `caarlos0/env`, failing on a missing or invalid variable, passed explicitly |
| Dev reload               | wgo, started by Compose (§3)                                                     |
| Debugger                 | Delve, from the tools module; in full mode only                                  |
| Dev tools                | one tools module for the repo (`tools/go.mod`, only `tool` lines, no code), **outside** `go.work`; every command there runs with `GOWORK=off`; installed into `$GOPATH/bin` by the maintenance service and called by name everywhere _(checked by Claude)_ |
| Linter install           | golangci-lint's released binary, installed in the all image (its authors don't support building it from source) |
| Database                 | PostgreSQL; its own Compose service in dev, with a named volume _(not tried)_    |
| Migrations, query layer  | **not decided**                                                                  |
| Logging, error response shape | **not decided**                                                             |

## 6. Type safety across ecosystems

| Purpose                                  | Technology / convention                                           |
| ---------------------------------------- | ----------------------------------------------------------------- |
| Shared types                             | three packages, none depending on another: `go-types` (between Go apps; exists with two or more Go apps), `ts-types` (between frontend apps, hand-written; two or more frontend apps), `shared-types` (between Go and TS; at least one Go app and one frontend app); each is created when the apps call for it, even empty |
| Where a type goes                        | the narrowest package covering its users; a type very likely to be shared soon goes in the right package from the start; a type that gains users stays where it is until moved, and is moved, never imported or re-exported from another types package |
| Go → TS types                            | tygo, in `shared-types` (Go structs are the source); the package is both a Go module and a TS package |
| Generated code                           | gitignored; produced by the `generate` task; the package exports it directly, with no compile step |
| Who depends on `generate`                | `typecheck`, `build`, `test:*`, `storybook:*`, and `dev:*` on the host, in every package: on their dependencies' `generate`, direct or through other packages (`generate` itself depends on `^generate`), and on their own package's |
| Generated code in dev                    | the maintenance service generates once before the apps start, then regenerates on every change (§3) |
| Runtime validation of API responses      | Zod schemas, hand-written in `shared-types` beside the generated types, checked against them (`satisfies z.ZodType<T>`) |
| Other Zod schemas                        | as types: in the narrowest package covering their users          |
| Runtime helpers (TS)                     | a hand-written shared TS package, only on second use (as shared Go code) |
| API client                               | server-only: one generic `apiFetch` core (address, secret, status check, Zod parse), made once per private API, + a one-line function per endpoint; the core is a shared package, `api-client`, once two or more frontend apps call a private API, otherwise in the app |
| API errors                               | a typed `ApiError`; `error()` for page loads, one shared toast helper for actions |

## 7. Testing

| Purpose                       | Technology / convention                                               |
| ----------------------------- | --------------------------------------------------------------------- |
| Unit and component tests (JS) | Vitest (`test:unit`)                                                  |
| Story tests, accessibility    | Storybook addon-vitest, addon-a11y                                    |
| Go tests                      | `go test` (built in, no extra tooling), as `test:unit`                |
| End-to-end tests              | Playwright, in a separate package per frontend app: `e2e/<app>/`, named `<app>-e2e` (`test:e2e`) _(not tried)_ |
| How e2e tests get a stack     | they run against a stack Compose already started; they don't launch servers |
| Where tests run               | in CI, not in git hooks                                               |
| At creation                   | every runner set up (unit, story, Go, e2e), so the first real test needs no setup |

## 8. Production deployment _(not tried)_

| Purpose                         | Technology / convention                                          |
| ------------------------------- | ---------------------------------------------------------------- |
| Frontend on a hosting platform  | the platform builds it from the repo, with the platform's adapter (the likely case) |
| Frontend as a container         | `turbo prune --docker` + multi-stage Dockerfile, with the Node adapter |
| Go apps                         | multi-stage Go build to a minimal runtime image (no `prune`); where they are hosted is **not decided** |
| Secrets                         | set in the host's or platform's own settings, never in the repo or an image |
| Build/deploy only what changed  | Turborepo git-diff `--filter`                                    |

## 9. Continuous integration _(not tried)_

| Purpose           | Technology / convention                                                        |
| ----------------- | ------------------------------------------------------------------------------ |
| Role              | the gate: the main branch only accepts pull requests whose run passed          |
| Platform          | GitHub Actions                                                                 |
| What runs         | everything pre-push runs (the Go sync check included), plus `build`, `test:unit` and `test:e2e`; tasks through Turborepo |
| Go modules        | each built and tested on its own, with `GOWORK=off`                            |
| Maintenance scripts | never one that writes; read-only ones are chosen one by one              |
| On pull requests  | only the packages affected and their dependents; the main branch runs everything |
| Environment       | inside the all image, built from the same Compose file as dev                  |
| End-to-end        | Compose starts the stack with a production build of the frontend, then `test:e2e` runs in the all image |
| Cache between runs | **not decided**                                                               |

## Not decided

Chosen per project until there is experience to decide from:

- **Persistence** beyond "PostgreSQL in its own service": migration tool and practices, query layer, database tests, seed data, production hosting.
- **Backend conventions**: logging, error response shape.
- **Authentication and sessions**: how users sign in, where sessions are stored, how a private API trusts the identity it receives, authorisation.
- **Public APIs**: CORS, rate limits, authentication.
- **Calls to internal Go services**: the protocol and the contract.
- **CI**: cache between runs, dependency update tooling.
- **Toolchain versions**: a single source of truth with one bump command.
- **Where the Go apps are hosted** in production.
- **UI component libraries**: which ones, if any, beside hand-written components.
