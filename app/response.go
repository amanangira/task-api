package app

import (
	"encoding/json"
	"log"
	"net/http"
)

type IDResponse struct {
	ID string `json:"id"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"message,omitempty"`
}

func WriteJSON(w http.ResponseWriter, httpCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	var encodeErr error
	if data != nil {
		if encodeErr = json.NewEncoder(w).Encode(data); encodeErr != nil {
			panic(encodeErr)
		}
	}
}

// BadRequest - helper for creating 400 response
// TODO - absorb app_error and simplify implementation
func BadRequest(w http.ResponseWriter, message string) {
	bodyBytes, err := json.Marshal(ErrorResponse{
		ErrorMessage: message,
	})

	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusBadRequest)
	if _, err2 := w.Write(bodyBytes); err2 != nil {
		log.Println(err2)
	}

	return
}
