package middleware

import (
	"net/http"
	"task/app"
)

type IMiddleware interface {
	Apply(a app.IAPI) func(http.Handler) http.Handler
}
