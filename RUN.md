# Run 008: components in the web app, Node adapter

- **Date**: 2026-10-20. **Base**: an empty directory. **Skill version**: `run/008`.
- **Inputs**: project `hello-008`, repo `https://github.com/Vectoras/iob-skill-001.git`, apps `api` and `web`, ports `api` 8090 and `web` 5180 (not the defaults), `web`'s Storybook 6006, frontend target `node` (`adapter-node`), components directly in the web app (no UI package).
- **Versions found**: Node 24.22.0, pnpm 12.14.0, Go 1.27.2 (1.27.3 held back by the 3-day rule), golangci-lint 2.15.0, Playwright 1.64.1, Debian trixie, sv 1.1.1, create-storybook 10.6.2 (10.6.3 held back).
- **Not in this branch**: `CLAUDE.md` and `.claude/skills/`, which the creation writes from its templates.

## Steps

One commit per step. Steps 1–3 (inputs, host, versions) write nothing, step 7 (the UI package) doesn't apply, and step 13 needed no fixes.

| Step | Commit message | Check |
|---|---|---|
| 4 | root files and images | pass (the apps' folders had to be created before their `.env` files) |
| 5 | root dependencies and the hooks | pass |
| 6 | Go tools module, types package, API | pass |
| 8 | the web app | pass |
| 9 | the shared TS package | pass |
| 10 | the Hello world chain in the web app | pass |
| 11 | the e2e package | pass |
| 12 | the project's docs | pass |

## Checklist

Run during the creation (step 13; checks 9 and 10 uncached), then graded again afterwards (including full mode).

| # | Check | Result |
|---|---|---|
| 1 | Stack starts lean | pass |
| 2 | Every program answers from the host | pass |
| 3 | Each container runs only its program | pass |
| 4 | Env vars reach the real processes | pass |
| 5 | One pnpm version | pass |
| 6 | No reinstall between containers | pass |
| 7 | Generator watcher once per edit | pass |
| 8 | Lint, format, type check | pass |
| 9 | Unit and story tests | pass, uncached |
| 10 | End-to-end | pass, uncached |
| 11 | Cached outputs restored | pass |
| 12 | Git ignores what is generated | pass |
| 13 | Go layout | pass |
| 14 | Full mode | pass |
| 15 | Hooks | pass |

## Verdict

**Pass**: the project needed no fixes.

Found outside the checklist: `turbo.json`'s `^generate` reaches only direct dependencies, so `turbo run typecheck --filter web` doesn't generate the types `web` reaches through `shared`; with them deleted, it fails (`Cannot find module 'types'`). In dev the maintenance service hides it.
