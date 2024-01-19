package middleware

import (
	"context"
	"gitlab.com/tozd/go/errors"
	"net/http"
	"task/app"
	"task/app/app_error"
)

const authorizationHeader = "Authorization"
const contextUserIDKey = "user_id"

func AuthorizedUserID(r *http.Request) (string, error) {
	// TODO - add logic
	//idTokenString := stripBearerToken(r.Header.Get(authorizationHeader))

	return "", nil
}

//TODO - Consider decoupling from IAPI and instead implement a middleware interface that can be implemented by a struct
// that would hold any dependencies for that particular middleware following isolation and decouple design

func ApplyBasicAuthorizer(api app.IAPI) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dbUserID, err := AuthorizedUserID(r)
			if err != nil {
				app_error.NewError(errors.New("invalid unauthorised"), http.StatusInternalServerError, "").Log().HttpError(w)

				return
			}
			// TODO - query DB if user exists and load required details for the request context

			ctx := r.Context()
			if ctx.Value(contextUserIDKey) != dbUserID {
				ctxWithUserSession := context.WithValue(ctx, contextUserIDKey, dbUserID)
				r = r.WithContext(ctxWithUserSession)
			}

			next.ServeHTTP(w, r)
		})
	}
}
