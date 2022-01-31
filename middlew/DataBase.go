package middlew

import (
	"github.com/javier-de-juan/twittor-go/bd"
	"net/http"
)

func IsDbConnected(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if ! bd.IsConnected() {
			http.Error(writer, "Database is not connected", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(writer, request)
	}
}
