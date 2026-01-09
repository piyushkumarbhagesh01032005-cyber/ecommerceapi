package routes

import (
	controller "ecommerceapi/controllers"

	"github.com/gorilla/mux"
)

func CartItemRoutes(r *mux.Router) {
	r.HandleFunc("/carts", controller.GetCartItems).Methods("GET")
	r.HandleFunc("/carts", controller.CreateCartItem).Methods("POST")
	r.HandleFunc("/carts{id}", controller.UpdateCartItem).Methods("PUT")
	r.HandleFunc("/carts{id}", controller.DeleteCartItem).Methods("DELETE")
}
