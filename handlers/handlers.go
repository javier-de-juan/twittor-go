package handlers

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handle() {
	router  := mux.NewRouter()
	PORT    := getPort()
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

func getPort() string {
	PORT := os.Getenv("TWITTOR_PORT")

	if PORT == "" {
		PORT = "8080"
	}

	return PORT
}
