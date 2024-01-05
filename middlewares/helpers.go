package middlewares

import (
	"net/http"

	"github.com/gorilla/context"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, ok := r.BasicAuth()

		if !ok {
			errMsg := "Authentication error!"
			http.Error(w, errMsg, http.StatusForbidden)
			return
		}
		context.Set(r, "isAdmin", true)

		next(w, r)
	}
}
