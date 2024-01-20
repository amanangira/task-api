package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"task/persistence/models"
	"task/service/mocks"
	"testing"
	"time"
)

func TestTaskCreate(t *testing.T) {
	t.Run("Blank Request", func(t *testing.T) {
		controller := NewTaskController(mocks.NewITaskService(t))
		httpRecorder := httptest.NewRecorder()
		body := bytes.NewBufferString(``)
		request, err := http.NewRequest(http.MethodPost, "/api/task", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		controller.Create(httpRecorder, request)
		result := httpRecorder.Result()
		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("error: expected http status code 400, got %d", result.StatusCode)
		}
	})

	t.Run("Blank Body Request", func(t *testing.T) {
		controller := NewTaskController(mocks.NewITaskService(t))
		httpRecorder := httptest.NewRecorder()
		body := bytes.NewBufferString(`{}`)
		request, err := http.NewRequest(http.MethodPost, "/api/task", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		controller.Create(httpRecorder, request)
		result := httpRecorder.Result()
		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("error: expected http status code 400, got %d", result.StatusCode)
		}
	})

	t.Run("Blank Body Request", func(t *testing.T) {
		taskServiceMock := mocks.NewITaskService(t)
		controller := NewTaskController(taskServiceMock)
		httpRecorder := httptest.NewRecorder()
		dueAt := time.Unix(1515151515, 0)
		body := bytes.NewBufferString(fmt.Sprintf(`{
    "title" : "Test title",
    "description" : "test description",
    "priority" : "p0",
    "due_at": "%s"}`, dueAt.Format(time.RFC3339)))
		request, err := http.NewRequest(http.MethodPost, "/api/task", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		taskServiceMock.On("Create", request.Context(), models.Task{
			Title:       "Test title",
			Description: "test description",
			Priority:    models.P0TaskPriority,
			DueAt:       dueAt,
		}).Return("46908780-c9d0-4f63-9cda-f0f858418074", nil)

		controller.Create(httpRecorder, request)
		result := httpRecorder.Result()
		if result.StatusCode != http.StatusOK {
			t.Errorf("error: expected http status code 200, got %d", result.StatusCode)
		}
	})
}
