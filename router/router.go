package router

import (
	"github.com/gorilla/mux"

	"github.com/Fedkasin/users-api/controllers"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	Router.HandleFunc("/users", controllers.GetAllUsersController).Methods("GET")
	Router.HandleFunc("/users", controllers.CreateUserController).Methods("POST")

	Router.HandleFunc("/users/{id}", controllers.GetUserByIdController).Methods("GET")
	Router.HandleFunc("/users/{id}", controllers.UpdateUserByIdController).Methods("PUT")
	Router.HandleFunc("/users/{id}", controllers.RemoveUserByIdController).Methods("DELETE")
}


