# Run 015: two of each

- **Date**: 2026-10-28. **Base**: an empty directory; images built by this run, some layers from the earlier runs' build cache (the cache was cleared before the round's first run). **Skill version**: `run/015` (run 014's version plus fixes 93–97).
- **Inputs**: project `hello-015`; two Go apps, both private APIs: `accounts` (8080) and `orders` (8081); two SvelteKit apps: `shop` (5173, Storybook 6007, target unknown: `adapter-auto`) and `admin` (5174, Storybook 6008, `node`); components in the shared UI package `packages/ui-components` (Storybook 6006). Repo URL the default: `https://github.com/Vectoras/hello-015`.
- **Versions found**: Node 24.22.1, pnpm 12.15.0 (12.16.0 and 12.16.1 held back by the 3-day rule), Debian trixie, Go 1.27.3, golangci-lint 2.15.0, Playwright 1.64.1, sv 1.1.1, create-storybook 10.6.3 (10.6.4 and 10.6.5 held back).
- **Packages the apps called for**: `go-types`, `shared-types`, `ts-types`, `api-client` (the first app's client moved into it).
- **Not in this branch**: `CLAUDE.md` and `.claude/skills/`, which the creation writes from its templates.

## Steps

One commit per step that writes files, steps 6–8 once per app. Steps 1–3 write nothing; steps 9 and 10 changed nothing, so they have no commit.

| Step | Commit message | Check |
|---|---|---|
| 4 | root files and images | pass |
| 5 | root dependencies and the hooks | pass |
| 6 | accounts (with the Go side) | pass |
| 6 | orders (with `go-types`) | pass |
| 7 | ui-components (with the TS side) | pass |
| 8 | shop (with `shared-types`, the wiring and its E2E package) | pass after a fix: `.golangci.yml` gets `run.allow-serial-runners: true` (golangci-lint failed now and then with "parallel golangci-lint is running" when Turborepo linted several Go modules at once) |
| 8 | admin (with `ts-types`, `api-client` and its E2E package) | pass after a workaround: the `api-client` move's wording missed the `./api-client` imports in `src/lib/server/` |

## Checklist

Run during the creation (step 10), then graded again afterwards (full mode included).

| # | Check | Result |
|---|---|---|
| 1 | Stack starts lean | pass |
| 2 | Every program answers from the host | pass |
| 3 | Each container runs only its program | pass |
| 4 | Env vars reach the real processes | pass |
| 5 | One pnpm version | pass |
| 6 | No reinstall between containers | pass |
| 7 | Generator watcher once per edit | pass |
| 8 | (merged into 15) | — |
| 9 | Unit and story tests | pass, uncached |
| 10 | End-to-end | pass, uncached |
| 11 | Cached outputs restored; generated code comes first | pass |
| 12 | Git ignores what is generated | pass |
| 13 | Go layout | pass |
| 14 | Full mode, and back to lean | pass |
| 15 | Hooks | pass |
| 16 | OpenAPI descriptions | pass |

## Verdict

**Fail**: the checks pass, but the project needed one fix (golangci-lint's lock with several Go modules) and one step a workaround (the `api-client` move).
