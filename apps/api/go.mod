module github.com/Vectoras/iob-skill-001/apps/api

go 1.27.2

require (
	github.com/Vectoras/iob-skill-001/packages/types v0.0.0
	github.com/caarlos0/env/v11 v11.4.1
	github.com/danielgtaylor/huma/v2 v2.39.1
	github.com/go-chi/chi/v5 v5.5.0
)

replace github.com/Vectoras/iob-skill-001/packages/types => ../../packages/types
