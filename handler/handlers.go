package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/CristianArboleda/gotwittor/middleware"
	"github.com/CristianArboleda/gotwittor/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*RoutesHandler: build the server listener*/
func RoutesHandler() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDB(routes.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routes.Login)).Methods("POST")
	router.HandleFunc("/profile", middleware.CheckDB(middleware.CheckJWT(routes.GetProfile))).Methods("GET")
	router.HandleFunc("/profile", middleware.CheckDB(middleware.CheckJWT(routes.UpdatePerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middleware.CheckDB(middleware.CheckJWT(routes.SaveTweet))).Methods("POST")
	router.HandleFunc("/tweet", middleware.CheckDB(middleware.CheckJWT(routes.GetTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	//Give all permissions
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
