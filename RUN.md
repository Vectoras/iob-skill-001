# Run 020: two of each

- **Date**: 2026-10-30. **Base**: an empty directory; images built by this run, some layers from the earlier runs' build cache (the cache was cleared before run 017). **Skill version**: `run/020` (`run/019`'s plus the fixes of findings 130–136).
- **Inputs**: project `hello-020`; two Go apps, both private APIs: `accounts` (8080) and `orders` (8081); two SvelteKit apps: `shop` (5173, Storybook 6007, target unknown: `adapter-auto`) and `admin` (5174, Storybook 6008, `node`); components in the shared UI package `packages/ui-components` (Storybook 6006); repo URL the default.
- **Versions found**: Node 24.22.1 (26.13.0 and 24.22.2 held back by the 3-day rule), pnpm 12.16.0 (12.16.1 held back), Debian trixie, Go 1.27.3, golangci-lint 2.15.0, Playwright 1.64.1, sv 1.1.1, create-storybook 10.6.4 (10.6.5 held back), chi v5.5.0, huma v2.40.0, env v11.4.1, wgo v0.7.2, Delve v1.27.2, tygo v0.2.21.
- **Packages the apps called for**: `go-types`, `shared-types`, `ts-types`, `api-client` (the first app's client moved into it).
- **Not in this branch**: `CLAUDE.md` and `.claude/skills/`, which the creation writes from its templates.

## Steps

One commit per step that writes files, steps 6–8 once per app or package. Steps 1–3 write nothing; step 10 changed nothing, so it has no commit.

| Step | Commit message | Check |
|---|---|---|
| 4 | root files and images | pass |
| 5 | root dependencies and the hooks | pass |
| 6 | accounts (with the Go side) | pass |
| 6 | orders (with `go-types`) | pass |
| 7 | ui-components (with the TS side) | pass |
| 8 | shop (with `shared-types`, the wiring and its E2E package) | pass |
| 8 | admin (with `ts-types`, `api-client` and its E2E package) | pass |
| 9 | the project's docs (reorders README rows) | pass |

## Checklist

Run during the creation (step 10), then graded again afterwards (full mode included). This shape has every need, so nothing is skipped.

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
| 11 | Cached outputs restored; generated code comes first | pass, Go apps' builds included |
| 12 | Git ignores what is generated | pass |
| 13 | Go layout | pass; golangci-lint's cache in `.cache/` through Turborepo |
| 14 | Full mode, and back to lean | pass |
| 15 | Hooks | pass |
| 16 | OpenAPI descriptions | pass |

Grading also covered the runbook's stop, take down and bring back (17 of 17 pass) and the Go modules' 3-day release age.

## Go lint's load

Four Go modules (`shared-types`' lint also runs ESLint), golangci-lint's cache cleaned before each; Go's build cache stayed warm. Sampled about every 2 seconds, so the peaks are approximate. No container limit; the Docker VM has 7.75 GiB.

| Concurrency | Peak memory | Peak CPU | Time |
|---|---|---|---|
| Turborepo's default | 1.31 GiB | 787% | 5.23 s |
| 1 | 868 MiB | 649% | 7.81 s |

## Verdict

**Pass**: the project needed no fix and no workaround, and no command was refused.
