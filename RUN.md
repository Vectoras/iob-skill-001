# Run 012: frontend only

- **Date**: 2026-10-28. **Base**: an empty directory, with an empty build cache (every image built from scratch). **Skill version**: `run/012`: apps optional (one or more Go apps, SvelteKit apps or both), each added through the project's `add-package`; the types packages their mix calls for.
- **Inputs**: project `hello-012`; no Go app; one SvelteKit app `web` on 5173 (Storybook 6007; 6006 kept for a UI package), components in the app, deployment target `node`.
- **Versions found**: Node 24.22.1, pnpm 12.15.0 (12.16.0 and 12.16.1 held back by the 3-day rule), Debian trixie, Playwright 1.64.1, sv 1.1.1, create-storybook 10.6.3 (10.6.4 and 10.6.5 held back).
- **Not in this branch**: `CLAUDE.md` and `.claude/skills/`, which the creation writes from its templates.

## Steps

One commit per step that writes files. Steps 1–3 write nothing; steps 6 (Go apps) and 7 (UI package) don't apply to this shape; step 10 needed no fixes, so it has no commit.

| Step | Commit message | Check |
|---|---|---|
| 4 | root files and images | pass |
| 5 | root dependencies and the hooks | pass |
| 8 | web (with the TS side and its E2E package) | pass |
| 9 | the project's docs | pass |

## Checklist

Run during the creation (step 10), then graded again afterwards (full mode included).

| # | Check | Result |
|---|---|---|
| 1 | Stack starts lean | pass |
| 2 | Every program answers from the host | pass |
| 3 | Each container runs only its program | pass |
| 4 | Env vars reach the real processes | skipped: no private API |
| 5 | One pnpm version | pass |
| 6 | No reinstall between containers | pass |
| 7 | Generator watcher once per edit | skipped: no `shared-types` |
| 8 | (merged into 15) | — |
| 9 | Unit and story tests | pass, uncached |
| 10 | End-to-end | pass, uncached |
| 11 | Cached outputs restored; generated code comes first | pass (build); the generated-code part skipped: no `shared-types` |
| 12 | Git ignores what is generated | pass |
| 13 | Go layout | Go part skipped: no Go app; the version check passes |
| 14 | Full mode, and back to lean | pass |
| 15 | Hooks | pass |
| 16 | OpenAPI descriptions | skipped: no API |

## Verdict

**Pass**: the project needed no fixes and no step needed a workaround.
