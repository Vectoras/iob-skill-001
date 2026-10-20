# Run 003: defaults

- **Date**: 2026-10-20. **Base**: an empty directory. **Skill version**: `run/003`.
- **Inputs**: project `hello-003`, repo `git@github.com:Vectoras/iob-skill-001.git`, every default: apps `api` and `web`, ports 8080, 5173, 6006, 6007, frontend target unknown (`adapter-auto`), components in the shared UI package `packages/ui-components`.
- **Versions found**: Node 24.22.0, pnpm 12.14.0, Go 1.27.2 (1.27.3 held back by the 3-day rule), golangci-lint 2.15.0, Playwright 1.64.1, Debian trixie, sv 1.1.1, create-storybook 10.6.2 (10.6.3 held back).
- **Not in this branch**: `CLAUDE.md` and `.claude/skills/`, which the creation writes from its templates.

## Steps

One commit per step. Steps 1–3 (inputs, host, versions) write nothing, and step 13 needed no fixes, so it has no commit either.

| Step | Commit message | Check |
|---|---|---|
| 4 | root files and images | pass |
| 5 | root dependencies and the hooks | pass |
| 6 | Go tools module, types package and API | pass |
| 7 | the UI package | pass |
| 8 | the web app | pass |
| 9 | the shared TS package | pass |
| 10 | the Hello world chain in the web app | pass |
| 11 | the e2e package | pass |
| 12 | the project's docs | pass |

## Checklist

Run during the creation (step 13; checks 9 and 10 uncached), then graded again afterwards.

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
| 9 | Unit and story tests | pass, uncached; both packages' browser tests ran in parallel |
| 10 | End-to-end | pass |
| 11 | Cached outputs restored | pass |
| 12 | Git ignores what is generated | pass |
| 13 | Go layout | pass |
| 14 | Full mode | pass once the check waited for the programs. Run right after the switch, it saw nothing yet: the debug build of the API takes a few seconds longer. Also, the check's probe of the API's debugger stays open about 2 minutes |
| 15 | Hooks | pass |

## Verdict

**Pass**: the project needed no fixes. Check 14 needs a wait before it looks.
