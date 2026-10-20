# Run 001: defaults

- **Date**: 2026-10-20. **Base**: an empty directory. **Skill version**: `run/001`.
- **Inputs**: project `hello-001`, repo `git@github.com:Vectoras/iob-skill-001.git`, every default: apps `api` and `web`, ports 8080, 5173, 6006, 6007, frontend target unknown (`adapter-auto`), components in the shared UI package `packages/ui-components`.
- **Versions found**: Node 24.22.0, pnpm 12.14.0, Go 1.27.2 (1.27.3 held back by the 3-day rule), golangci-lint 2.15.0, Playwright 1.64.1, Debian trixie, sv 1.1.1, create-storybook 10.6.2.
- **Not in this branch**: `CLAUDE.md` and `.claude/skills/`, which the creation writes from its templates.

## Steps

One commit per step; steps 1–3 (inputs, host, versions) write nothing.

| Step | Commit message | Check |
|---|---|---|
| 4 | root files and images | pass |
| 5 | root dependencies and the hooks | pass |
| 6 | Go tools module, types package and the API | pass |
| 7 | the UI package | pass after two fixes: a lint error in the scaffolded `vite.config.ts` (a triple-slash reference); `svelte-kit sync` failing with no route in the package |
| 8 | the web app | pass |
| 9 | the shared TS package | pass |
| 10 | the Hello world chain in the web app | pass after one fix: the page's links failed `svelte/no-navigation-without-resolve` |
| 11 | the e2e package | pass |
| 12 | the project's docs | pass |
| 13 | fixes from the checks | see below |

Also seen during the creation: on Docker Desktop for Linux, the editing container saw a stale `.git/index` after a `git add` on the host, so staging moved into the container too.

## Checklist

Run during the creation (step 13), then graded again afterwards against a forced (uncached) run.

| # | Check | Result |
|---|---|---|
| 1 | Stack starts lean | pass |
| 2 | Every program answers from the host | **failed first**: the home page threw `env_invalid` (`API_URL` missing). Storybook and the unit tests, which run SvelteKit's Vite plugin without `API_URL`, rewrite the dev server's generated env module. Fixed by giving them `API_URL` too; then pass |
| 3 | Each container runs only its program | pass |
| 4 | Env vars reach the real processes | pass |
| 5 | One pnpm version | pass |
| 6 | No reinstall between containers | pass |
| 7 | Generator watcher once per edit | pass |
| 8 | Lint, format, type check | pass |
| 9 | Unit and story tests | **fail** when run uncached: the two packages' browser tests start in parallel and both take Vitest's browser port (`Port 63315 is already in use`); each passes alone |
| 10 | End-to-end | pass |
| 11 | Cached outputs restored | pass |
| 12 | Git ignores what is generated | pass |
| 13 | Go layout | pass |
| 14 | Full mode | pass |
| 15 | Hooks | **failed first**: pre-push's format check read `go.work.sum` while the Go sync check deleted it (parallel jobs). Fixed by ignoring it in Prettier; then pass |

## Verdict

**Fail**: check 9 fails, and checks 2 and 15 needed fixes during the creation.
