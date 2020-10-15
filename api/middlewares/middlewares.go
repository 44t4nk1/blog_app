package middlewares

import (
	"errors"
	"net/http"

	"github.com/44t4nk1/blog_app/api/responses"

	"github.com/44t4nk1/blog_app/api/auth"
)

//SetMiddlewareJSON ...
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

//SetMiddlewareAuthentication ...
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorised"))
			return
		}
		next(w, r)
	}
}
