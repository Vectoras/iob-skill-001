# Run 016: frontend only

- **Date**: 2026-10-30. **Base**: an empty directory; images built by this run, the build cache cleared before it. **Skill version**: `run/016`.
- **Inputs**: project `hello-016`; no Go app; one SvelteKit app, `web` (5173, Storybook 6007, target `node`); components in the app.
- **Versions found**: Node 24.22.1 (26.13.0 and 24.22.2 held back by the 3-day rule), pnpm 12.16.0 (12.16.1 held back), Debian trixie, Playwright 1.64.1, sv 1.1.1, create-storybook 10.6.4 (10.6.5 held back).
- **Packages the apps called for**: none.
- **Not in this branch**: `CLAUDE.md` and `.claude/skills/`, which the creation writes from its templates.

## Steps

One commit per step that writes files. Steps 1–3 write nothing; steps 6 and 7 don't apply; steps 9 and 10 changed nothing, so they have no commit.

| Step | Commit message | Check |
|---|---|---|
| 4 | root files and images | pass |
| 5 | root dependencies and the hooks | pass |
| 8 | web (with the TS side and its E2E package) | pass with leftovers: a safety check refused the deletion of the scaffolders' example files, so they stayed on disk, untracked |

## Checklist

Run during the creation (step 10), then graded again afterwards (full mode included). Every run below had the leftover files on disk.

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
| 10 | End-to-end | pass, uncached |
| 11 | Cached outputs restored | pass; the generated-code part skipped: no `shared-types` |
| 12 | Git ignores what is generated | **fail**: 6 untracked paths, `apps/web/.gitignore`, `apps/web/.npmrc` |
| 13 | Go layout | skipped: no Go app; `check-versions.sh` passes |
| 14 | Full mode, and back to lean | pass |
| 15 | Hooks | pass |
| 16 | OpenAPI descriptions | skipped: no API |

Grading also covered the runbook's stop, take down and bring back: 14 of 15 pass; the other is the same untracked files.

## Verdict

**Fail**: check 12 fails, and the committed tree was never tested without the leftover files.
