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
func TestGetTask(t *testing.T) {
	t.Run("Blank Request", func(t *testing.T) {
		controller := NewTaskController(mocks.NewITaskService(t))
		httpRecorder := httptest.NewRecorder()
		body := bytes.NewBufferString(``)
		request, err := http.NewRequest(http.MethodGet, "/api/task/taskID", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		controller.GetTask(httpRecorder, request)
		result := httpRecorder.Result()
		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("error: expected http status code 400, got %d", result.StatusCode)
		}
	})

	t.Run("Blank Body Request", func(t *testing.T) {
		controller := NewTaskController(mocks.NewITaskService(t))
		httpRecorder := httptest.NewRecorder()
		body := bytes.NewBufferString(`{}`)
		request, err := http.NewRequest(http.MethodGet, "/api/task/taskID", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		controller.GetTask(httpRecorder, request)
		result := httpRecorder.Result()
		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("error: expected http status code 400, got %d", result.StatusCode)
		}
	})

	t.Run("Correct Body Request", func(t *testing.T) {
		taskServiceMock := mocks.NewITaskService(t)
		controller := NewTaskController(mocks.NewITaskService(t))
		taskID := "123"
		body := bytes.NewBufferString(`{"taskID": "123"}`)
		httpRecorder := httptest.NewRecorder()

		// Mock Task returned by the service
		mockTask := models.Task{
			ID:          taskID,
			Title:       "Test Task",
			Description: "Test Description",
			Priority:    "LowPriority",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DueAt:       time.Now().Add(time.Hour),
		}

		taskServiceMock.On("GetTask", taskID).Return(mockTask, nil)

		request, err := http.NewRequest(http.MethodGet, "/api/task/taskID", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		controller.GetTask(httpRecorder, request)

		result := httpRecorder.Result()
		if result.StatusCode != http.StatusOK {
			t.Errorf("error: expected http status code 200, got %d", result.StatusCode)
		}

		// Assert that the expectations were met
		taskServiceMock.AssertExpectations(t)
	})

}

func TestUpdateTask(t *testing.T) {
	t.Run("Blank Request", func(t *testing.T) {
		controller := NewTaskController(mocks.NewITaskService(t))
		httpRecorder := httptest.NewRecorder()
		body := bytes.NewBufferString(``)
		request, err := http.NewRequest(http.MethodPatch, "/api/task/taskID", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		controller.UpdateTask(httpRecorder, request)
		result := httpRecorder.Result()
		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("error: expected http status code 400, got %d", result.StatusCode)
		}
	})

	t.Run("Blank Body Request", func(t *testing.T) {
		controller := NewTaskController(mocks.NewITaskService(t))
		httpRecorder := httptest.NewRecorder()
		body := bytes.NewBufferString(`{}`)
		request, err := http.NewRequest(http.MethodPatch, "/api/task/taskID", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		controller.UpdateTask(httpRecorder, request)
		result := httpRecorder.Result()
		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("error: expected http status code 400, got %d", result.StatusCode)
		}
	})

	t.Run("Correct Body Request", func(t *testing.T) {
		taskServiceMock := mocks.NewITaskService(t)
		controller := NewTaskController(mocks.NewITaskService(t))
		dueAt := time.Unix(1515151515, 0)
		taskID := "123"
		body := bytes.NewBufferString(fmt.Sprintf(`{
    "title" : "Test title",
	"taskId": "123",
    "description" : "test description",
    "priority" : "p0",
    "due_at": "%s"}`, dueAt.Format(time.RFC3339)))
		httpRecorder := httptest.NewRecorder()

		taskServiceMock.On("UpdateTask", taskID, models.Task{
			Title:       "Test title",
			ID:          "123",
			Description: "test description",
			Priority:    models.P0TaskPriority,
			DueAt:       dueAt,
		}).Return(nil)
		request, err := http.NewRequest(http.MethodPatch, "/api/task/taskID", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		controller.UpdateTask(httpRecorder, request)

		result := httpRecorder.Result()
		if result.StatusCode != http.StatusOK {
			t.Errorf("error: expected http status code 200, got %d", result.StatusCode)
		}

		// Assert that the expectations were met
		taskServiceMock.AssertExpectations(t)
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("Blank Request", func(t *testing.T) {
		controller := NewTaskController(mocks.NewITaskService(t))
		httpRecorder := httptest.NewRecorder()
		body := bytes.NewBufferString(``)
		request, err := http.NewRequest(http.MethodDelete, "/api/task/taskID", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		controller.DeleteTask(httpRecorder, request)
		result := httpRecorder.Result()
		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("error: expected http status code 400, got %d", result.StatusCode)
		}
	})

	t.Run("Blank Body Request", func(t *testing.T) {
		controller := NewTaskController(mocks.NewITaskService(t))
		httpRecorder := httptest.NewRecorder()
		body := bytes.NewBufferString(`{}`)
		request, err := http.NewRequest(http.MethodDelete, "/api/task/taskID", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		controller.DeleteTask(httpRecorder, request)
		result := httpRecorder.Result()
		if result.StatusCode != http.StatusBadRequest {
			t.Errorf("error: expected http status code 400, got %d", result.StatusCode)
		}
	})

	t.Run("Correct Body Request", func(t *testing.T) {
		taskServiceMock := mocks.NewITaskService(t)
		controller := NewTaskController(mocks.NewITaskService(t))
		taskID := "123"
		body := bytes.NewBufferString(`{"taskID": "123"}`)
		httpRecorder := httptest.NewRecorder()

		taskServiceMock.On("DeleteTask", taskID).Return(nil)
		request, err := http.NewRequest(http.MethodDelete, "/api/task/taskID", body)
		if err != nil {
			t.Errorf("error creating request: %s", err.Error())
		}

		controller.DeleteTask(httpRecorder, request)

		result := httpRecorder.Result()
		if result.StatusCode != http.StatusOK {
			t.Errorf("error: expected http status code 200, got %d", result.StatusCode)
		}

		// Assert that the expectations were met
		taskServiceMock.AssertExpectations(t)
	})
}
