# Run 018: backend only

- **Date**: 2026-10-30. **Base**: an empty directory; images built by this run, some layers from the earlier runs' build cache (the cache was cleared before run 017). **Skill version**: `run/018` (the same as `run/017`).
- **Inputs**: project `hello-018`; two Go apps: `catalog` (public API, 8080) and `worker` (internal service, 8081); no SvelteKit app; repo URL the default.
- **Versions found**: Node 24.22.1 (26.13.0 and 24.22.2 held back by the 3-day rule), pnpm 12.16.0 (12.16.1 held back), Debian trixie, Go 1.27.3, golangci-lint 2.15.0, chi v5.5.0, huma v2.40.0, env v11.4.1, wgo v0.7.2, Delve v1.27.2 (no Go module held back).
- **Packages the apps called for**: `go-types`.
- **Not in this branch**: `CLAUDE.md` and `.claude/skills/`, which the creation writes from its templates.

## Steps

One commit per step that writes files, step 6 once per app. Steps 1–3 write nothing; steps 7 and 8 don't apply; steps 9 and 10 changed nothing, so they have no commit.

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
| 2 | Every program answers from the host | pass |
| 3 | Each container runs only its program | pass |
| 4 | Env vars reach the real processes | skipped: no app wired to a private API |
| 5 | One pnpm version | pass |
| 6 | No reinstall between containers | pass |
| 7 | Generator watcher once per edit | skipped: no `shared-types` |
| 8 | (merged into 15) | — |
| 9 | Unit and story tests | pass, uncached |
| 10 | End-to-end | skipped: no SvelteKit app |
| 11 | Cached outputs restored | skipped: no SvelteKit app, no `shared-types` |
| 12 | Git ignores what is generated | pass |
| 13 | Go layout | pass |
| 14 | Full mode, and back to lean | pass |
| 15 | Hooks | pass |
| 16 | OpenAPI descriptions | pass |

Grading also covered the runbook's stop, take down and bring back (13 of 13 pass) and the Go modules' 3-day release age (`go-release-age.sh` read 151 index pages and printed the chi version the apps use).

## Go lint's load

Three Go modules, with golangci-lint's cache cleaned first. No container limit is set; the Docker VM has 7.75 GiB and 10 CPUs. Sampled about every 2 seconds, so the peaks are lower bounds, on modules of Hello world size.

| Concurrency | Peak memory | Peak CPU | Time |
|---|---|---|---|
| Turborepo's default | 1.05 GiB | 923% | 3.45 s |
| 1 | 593 MiB | 495% | 5.05 s |

## Verdict

**Pass**: the project needed no fix and no workaround.
