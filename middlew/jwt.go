package middlew

import (
	"github.com/javier-de-juan/twittor-go/routers"
	"net/http"
)

func IsValidJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _, _, err := routers.IsValidRequestToken(request.Header.Get("Authorization"))

		if err != nil {
			http.Error(writer, "Token not valid" + err.Error(), http.StatusUnauthorized)
		}

		next.ServeHTTP(writer, request)
	}
}