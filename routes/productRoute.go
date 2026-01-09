package routes

import (
	controllers "ecommerceapi/controllers"

	"github.com/gorilla/mux"
)

func ProductRoutes(r *mux.Router) {
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")

}
