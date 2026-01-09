package routes

import (
	controller "ecommerceapi/controllers"

	"github.com/gorilla/mux"
)

func OrderRoutes(r *mux.Router) {
	r.HandleFunc("/orders", controller.GetOrders).Methods("GET")
	r.HandleFunc("/orders", controller.CreateOrder).Methods("POST")
	r.HandleFunc("/orders{id}", controller.UpdateOrder).Methods("PUT")
	r.HandleFunc("/orders{id}", controller.DeleteOrder).Methods("DELETE")
}
