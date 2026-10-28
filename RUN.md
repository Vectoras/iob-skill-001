# Run 014: defaults

- **Date**: 2026-10-28. **Base**: an empty directory; images built by this run (the build cache was cleared before the round's first run). **Skill version**: `run/014` (run 013's version plus fixes 86, 87, 89, 90).
- **Inputs**: project `hello-014`, every other default: one Go app `api` (a private API, 8080) and one SvelteKit app `web` (5173, Storybook 6007, deployment target unknown: `adapter-auto`), components in the shared UI package `packages/ui-components` (Storybook 6006). Repo URL the default: no `origin`, so `https://github.com/Vectoras/hello-014`.
- **Versions found**: Node 24.22.1, pnpm 12.15.0 (12.16.0 and 12.16.1 held back by the 3-day rule), Debian trixie, Go 1.27.3, golangci-lint 2.15.0, Playwright 1.64.1, sv 1.1.1, create-storybook 10.6.3 (10.6.4 and 10.6.5 held back).
- **Not in this branch**: `CLAUDE.md` and `.claude/skills/`, which the creation writes from its templates.

## Steps

One commit per step that writes files. Steps 1–3 write nothing; step 10 needed no fixes, so it has no commit.

| Step | Commit message | Check |
|---|---|---|
| 4 | root files and images | pass |
| 5 | root dependencies and the hooks | pass |
| 6 | api (with the Go side) | pass |
| 7 | ui-components (with the TS side) | pass |
| 8 | web (with `shared-types`, the wiring to `api` and its E2E package) | pass |
| 9 | the project's docs (the README's "Not set up yet" lacked "UI component libraries") | pass |

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

**Pass**: the project needed no fixes and no step needed a workaround.
