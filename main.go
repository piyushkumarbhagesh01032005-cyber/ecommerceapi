package main

import (
	"ecommerceapi/database"
	"ecommerceapi/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	database.Connect()

	r := mux.NewRouter()

	routes.OrderItemRoutes(r)
	routes.OrderRoutes(r)
	routes.ProductRoutes(r)
	routes.CartItemRoutes(r)
	routes.CartRoutes(r)
	routes.UserRoutes(r)

	fmt.Println("Server started at http://localhost:8085")
	log.Fatal(http.ListenAndServe(":8085", r))
}
