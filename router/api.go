package router

import (
	"github.com/go-chi/chi/v5"
	"task/app"
	"task/app/container"
	"task/middleware"
)

func (a api) Init(r *chi.Mux, apiInstance app.IAPI) {
	r.Use(middleware.ApplyBasicAuthorizer(apiInstance))

	r.Post("/task", container.APIContainer.TaskController.Create)
}
