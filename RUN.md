# Run 010: components in the web app, Node adapter

- **Date**: 2026-10-23. **Base**: an empty directory (Docker's build cache warm from run 009: the round's clean build is run 009's). **Skill version**: `run/010`.
- **Inputs**: project `hello-010`, repo `https://github.com/Vectoras/iob-skill-001.git`, apps `api` and `web`, ports `api` 8090 and `web` 5180 (not the defaults), `web`'s Storybook 6006, frontend target `node` (`adapter-node`), components directly in the web app (no UI package).
- **Versions found**: Node 24.22.0 (24.22.1 held back by the 3-day rule), pnpm 12.14.0 (12.15.0 held back), Go 1.27.3, golangci-lint 2.15.0, Playwright 1.64.1, Debian trixie, sv 1.1.1, create-storybook 10.6.2 (10.6.3 held back).
- **Not in this branch**: `CLAUDE.md` and `.claude/skills/`, which the creation writes from its templates.

## Steps

One commit per step. Steps 1–3 (inputs, host, versions) write nothing, step 7 (the UI package) doesn't apply, and step 13 needed no fixes.

| Step | Commit message | Check |
|---|---|---|
| 4 | root files and images | pass |
| 5 | root dependencies and the hooks | pass |
| 6 | Go tools module, types package, API | pass |
| 8 | the web app | pass |
| 9 | the shared TS package | pass |
| 10 | the Hello world chain in the web app | pass |
| 11 | the e2e package | pass |
| 12 | the project's docs | pass (the README's "Not set up yet" list: the skill's wording contradicted the guide; the run kept the template's list) |

## Checklist

Run during the creation (step 13), then graded again afterwards (full mode included).

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
| 11 | Cached outputs restored (`build/` included); generated code comes first | pass |
| 12 | Git ignores what is generated | pass |
| 13 | Go layout | pass |
| 14 | Full mode, and back to lean | pass |
| 15 | Hooks | pass |

## Verdict

**Pass**: the project needed no fixes.
