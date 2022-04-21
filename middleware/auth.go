package middleware

import (
	"context"
	"fmt"
	"net/http"
)

func Auth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := request.Header.Get("authorization")
		if token == "" {
			token = request.URL.Query().Get("token")
		}
		if token == "" {
			http.Error(writer, "Need auth", 401)
			fmt.Println("Need auth for request from ", request.RemoteAddr)
		} else {
			fmt.Println("token auth = ", token)
			//next.ServeHTTP(writer, request)
			next.ServeHTTP(writer, request.WithContext(context.WithValue(request.Context(), "token", token)))
		}
	})
}
