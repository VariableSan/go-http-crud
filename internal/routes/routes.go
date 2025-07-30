package routes

import (
	"github.com/VariableSan/go-http-crud/internal/app"
	"github.com/go-chi/chi"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	return r
}
