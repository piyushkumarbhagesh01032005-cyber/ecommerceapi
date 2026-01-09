package routes

import (
	controller "ecommerceapi/controllers"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	r.HandleFunc("/carts", controller.GetCarts).Methods("GET")
	r.HandleFunc("/carts", controller.CreateCart).Methods("POST")
	r.HandleFunc("/carts{id}", controller.UpdateCart).Methods("PUT")
	r.HandleFunc("/carts{id}", controller.DeleteCart).Methods("DELETE")
}
