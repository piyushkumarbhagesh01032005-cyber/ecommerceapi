package routes

import (
	controller "ecommerceapi/controllers"

	"github.com/gorilla/mux"
)

func OrderItemRoutes(r *mux.Router) {
	r.HandleFunc("/orderitems", controller.GetOrderItems).Methods("GET")
	r.HandleFunc("/orderitems", controller.CreateOrderItem).Methods("POST")
	r.HandleFunc("/orderitems{id}", controller.UpdateOrderItem).Methods("PUT")
	r.HandleFunc("/orderitems{id}", controller.DeleteOrderItem).Methods("DELETE")

}
