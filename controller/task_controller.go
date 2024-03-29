package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"task/app"
	"task/app/app_error"
	"task/persistence/models"
	"task/service"
)

// TODO - add validation for incoming requests

const TaskIDRequestKey = "taskID"

type ITaskController interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	ListTasks(w http.ResponseWriter, r *http.Request)
}

type TaskController struct {
	taskService service.ITaskService
}

func NewTaskController(
	taskService service.ITaskService) *TaskController {
	return &TaskController{
		taskService: taskService,
	}
}

func (t TaskController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input models.Task
	decodeErr := json.NewDecoder(r.Body).Decode(&input)
	if decodeErr != nil {
		app.BadRequest(
			w,
			app_error.NewError(decodeErr, http.StatusBadRequest, "invalid input").Log().Error(),
		)

		return
	}

	if err := validateCreatePayload(input); err != nil {
		app.BadRequest(
			w,
			err.Error(),
		)

		return
	}

	id, err := t.taskService.Create(ctx, input)
	if err != nil {
		app_error.NewError(err, http.StatusInternalServerError, "").Log().HttpError(w)
		return
	}

	app.WriteJSON(w, http.StatusOK, struct {
		ID string
	}{
		ID: id,
	})
}

func (t TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		TaskID string `json:"taskID"`
	}

	decodeErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if decodeErr != nil || requestBody.TaskID == "" {
		app.BadRequest(
			w,
			app_error.NewError(decodeErr, http.StatusBadRequest, "invalid input").Log().Error(),
		)

		return
	}
	taskID := requestBody.TaskID
	task, err := t.taskService.GetTask(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	app.WriteJSON(
		w,
		http.StatusOK,
		task)
}

func (t TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var input models.Task
	decodeErr := json.NewDecoder(r.Body).Decode(&input)
	if decodeErr != nil {
		app.BadRequest(
			w,
			app_error.NewError(decodeErr, http.StatusBadRequest, "invalid input").Log().Error(),
		)

		return
	}

	if err := validateCreatePayload(input); err != nil {
		app.BadRequest(
			w,
			err.Error(),
		)

		return
	}

	err := t.taskService.UpdateTask(input.ID, input)
	if err != nil {
		app_error.NewError(err, http.StatusInternalServerError, "").Log().HttpError(w)
		return
	}

	app.WriteJSON(w, http.StatusOK, nil)
}

func (t TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		TaskID string `json:"taskID"`
	}

	decodeErr := json.NewDecoder(r.Body).Decode(&requestBody)
	if decodeErr != nil || requestBody.TaskID == "" {
		app.BadRequest(
			w,
			app_error.NewError(decodeErr, http.StatusBadRequest, "invalid input").Log().Error(),
		)

		return
	}
	taskID := requestBody.TaskID

	err := t.taskService.DeleteTask(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	app.WriteJSON(
		w,
		http.StatusOK,
		nil)
}

func (t TaskController) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := t.taskService.ListTasks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	app.WriteJSON(
		w,
		http.StatusOK,
		tasks)
}

func validateCreatePayload(input models.Task) error {
	if input.Title == "" {
		return fmt.Errorf("invalid title, expected valid string got %s", input.Title)
	}

	if input.Description == "" {
		return fmt.Errorf("invalid description, expected valid string got %s", input.Description)
	}

	if input.Priority == "" {
		return fmt.Errorf("invalid priority, expected valid string got %s", input.Priority)
	}

	if !input.Priority.Valid() {
		return fmt.Errorf("invalid priority, must be one of %s, got %s", strings.Join(models.ListTaskPriority(), ","), input.Priority)
	}

	return nil
}
