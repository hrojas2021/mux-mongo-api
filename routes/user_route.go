package routes

import (
	"mux-mongo-api/controllers"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/users", controllers.CreateUser()).Methods("POST")
	router.HandleFunc("/users/{userId}", controllers.GetAUser()).Methods("GET")
	router.HandleFunc("/users/{userId}", controllers.EditAUser()).Methods("PUT")
	router.HandleFunc("/users/{userId}", controllers.DeleteAUser()).Methods("DELETE")
	router.HandleFunc("/users", controllers.GetAllUser()).Methods("GET")
}
