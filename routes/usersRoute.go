package routes

import (
	controller "ecommerceapi/controllers"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/users", controller.CreateUser).Methods("POST")
}
