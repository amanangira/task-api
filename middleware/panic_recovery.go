package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

func ApplyPanicRecovery(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Print(err)
				log.Print(string(debug.Stack()))
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal server error"))
			}
		}()

		handle.ServeHTTP(w, r)
	})
}
