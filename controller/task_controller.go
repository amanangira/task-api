package controller

import (
	"encoding/json"
	"net/http"
	"task/app"
	"task/app/app_error"
	"task/persistence/models"
	"task/service"
)

// TODO - add validation for incoming requests

type TaskController struct {
	taskService service.ITaskService
}

func NewTaskController(
	taskService service.ITaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (p TaskController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input models.Task

	// TODO - decode to view equivalent and run validations on it
	decodeErr := json.NewDecoder(r.Body).Decode(&input)
	if decodeErr != nil {
		app_error.NewError(decodeErr, http.StatusBadRequest, "invalid input").Log().HttpError(w)
		return
	}

	err := p.taskService.Create(ctx, &input)
	if err != nil {
		app_error.NewError(err, http.StatusInternalServerError, "").Log().HttpError(w)
		return
	}

	app.WriteJSON(w, http.StatusOK, struct {
		ID string
	}{
		ID: input.ID,
	})
}
