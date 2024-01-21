package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"task/app"
	"task/app/container"
	"task/controller"
	"task/middleware"
)

func (a api) Init(r *chi.Mux, apiInstance app.IAPI) {
	r.Use(middleware.ApplyBasicAuthorizer(apiInstance))

	specificTaskPath := fmt.Sprintf("/task/{%s}", controller.TaskIDRequestKey)
	r.Post("/task", container.APIContainer.TaskController.Create)
	r.Get("/task", container.APIContainer.TaskController.ListTasks)
	r.Delete(
		specificTaskPath,
		container.APIContainer.TaskController.DeleteTask)
	r.Get(
		specificTaskPath,
		container.APIContainer.TaskController.GetTask)
	r.Patch(
		specificTaskPath,
		container.APIContainer.TaskController.UpdateTask)
}
