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

// RoutesHandler : build the server listener
func RoutesHandler() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDB(routes.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routes.Login)).Methods("POST")

	router.HandleFunc("/profile", middleware.CheckDB(middleware.CheckJWT(routes.GetProfile))).Methods("GET")
	router.HandleFunc("/profile", middleware.CheckDB(middleware.CheckJWT(routes.UpdateProfile))).Methods("PUT")
	router.HandleFunc("/profiles", middleware.CheckDB(middleware.CheckJWT(routes.GetProfilesByFilters))).Methods("GET")

	router.HandleFunc("/tweet", middleware.CheckDB(middleware.CheckJWT(routes.SaveTweet))).Methods("POST")
	router.HandleFunc("/tweet", middleware.CheckDB(middleware.CheckJWT(routes.GetTweets))).Methods("GET")
	router.HandleFunc("/tweet", middleware.CheckDB(middleware.CheckJWT(routes.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/followerstweets", middleware.CheckDB(middleware.CheckJWT(routes.GetFollowersTweets))).Methods("GET")

	router.HandleFunc("/avatar", middleware.CheckDB(middleware.CheckJWT(routes.GetAvatar))).Methods("GET")
	router.HandleFunc("/avatar", middleware.CheckDB(middleware.CheckJWT(routes.AddAvatar))).Methods("POST")
	router.HandleFunc("/banner", middleware.CheckDB(middleware.CheckJWT(routes.GetBanner))).Methods("GET")
	router.HandleFunc("/banner", middleware.CheckDB(middleware.CheckJWT(routes.AddBanner))).Methods("POST")

	router.HandleFunc("/relation", middleware.CheckDB(middleware.CheckJWT(routes.SaveRelation))).Methods("POST")
	router.HandleFunc("/relation", middleware.CheckDB(middleware.CheckJWT(routes.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/relation", middleware.CheckDB(middleware.CheckJWT(routes.GetRelation))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	//Give all permissions
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
