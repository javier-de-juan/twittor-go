package handlers

import (
	"github.com/gorilla/mux"
	"github.com/javier-de-juan/twittor-go/middlew"
	"github.com/javier-de-juan/twittor-go/routers"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func Handle() {
	router  := mux.NewRouter()

	router.HandleFunc("/register", middlew.IsDbConnected(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.IsDbConnected(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlew.IsDbConnected(middlew.IsValidJWT(routers.Profile))).Methods("GET")
	router.HandleFunc("/profile", middlew.IsDbConnected(middlew.IsValidJWT(routers.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.IsDbConnected(middlew.IsValidJWT(routers.Tweet))).Methods("POST")
	router.HandleFunc("/tweet", middlew.IsDbConnected(middlew.IsValidJWT(routers.GetTweets))).Methods("GET")
	router.HandleFunc("/tweet", middlew.IsDbConnected(middlew.IsValidJWT(routers.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/avatar", middlew.IsDbConnected(middlew.IsValidJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/avatar", middlew.IsDbConnected(middlew.IsValidJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/banner", middlew.IsDbConnected(middlew.IsValidJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/banner", middlew.IsDbConnected(middlew.IsValidJWT(routers.GetBanner))).Methods("GET")

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
