package middleware

import (
	"context"
	"net/http"
)

type contextKey string

var IsHxRequest = contextKey("class")

func Layout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if header hx-request is true
		ctx := r.Context()
		if r.Header.Get("hx-request") == "true" {
			ctx = context.WithValue(r.Context(), IsHxRequest, "false")
		} else {
			ctx = context.WithValue(r.Context(), IsHxRequest, "true")
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
