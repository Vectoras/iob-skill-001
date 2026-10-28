module github.com/Vectoras/hello-014/apps/api

go 1.27.3

require (
	github.com/Vectoras/hello-014/packages/shared-types v0.0.0
	github.com/caarlos0/env/v11 v11.4.1
	github.com/danielgtaylor/huma/v2 v2.40.0
	github.com/go-chi/chi/v5 v5.5.0
)

replace github.com/Vectoras/hello-014/packages/shared-types => ../../packages/shared-types
