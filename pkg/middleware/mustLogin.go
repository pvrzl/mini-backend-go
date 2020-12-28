package middleware

import (
	"net/http"
)

// HTTP middleware setting a value on the request context
func MustLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader != "dummy token" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
