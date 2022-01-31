package handlers

import (
	"github.com/gorilla/mux"
	"github.com/javier-de-juan/twittor-go/middlew"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handle() {
	router  := mux.NewRouter()

	router.HandleFunc("/registro", middlew.IsDbConnected(routers.Registro)).Methods("POST")

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
