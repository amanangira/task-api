package container

import (
	"task/app"
	"task/controller"
	"task/repository"
	"task/service"
)

type API struct {
	TaskController *controller.TaskController
}

var APIContainer API

//TODO - can use DI library like google/wire instead.

func InitControllers(api app.IAPI) {
	// Repository
	taskRepository := repository.NewTaskRepository(api.GetDBClient())

	// Services
	taskService := service.NewTaskService(taskRepository)

	// controllers
	userController := controller.NewTaskController(taskService)

	APIContainer = API{
		TaskController: userController,
	}
}
