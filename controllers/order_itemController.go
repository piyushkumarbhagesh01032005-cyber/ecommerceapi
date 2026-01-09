package controllers

import (
	"ecommerceapi/database"
	"ecommerceapi/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetOrderItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order_item []models.OrderItem

	result := database.DB.Find(&order_item)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(order_item)
}

func CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order_item models.OrderItem
	json.NewDecoder(r.Body).Decode(&order_item)

	result := database.DB.Create(&order_item)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order_item)

}

func UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var order_item models.OrderItem
	result := database.DB.First(&order_item, id)
	if result.Error != nil {
		http.Error(w, "Order_item Not Found", http.StatusNotFound)

		return
	}

	var input models.OrderItem
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Model(&order_item).Updates(map[string]interface{}{
		"product_id": input.ProductID,
		"quantity":   input.Quantity,
		"price":      input.Price,
	})

	json.NewEncoder(w).Encode(&order_item)

}

func DeleteOrderItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid Order_items ID", http.StatusBadRequest)
		return
	}

	result := database.DB.Delete(&models.OrderItem{}, id)
	if result.RowsAffected == 0 {
		http.Error(w, "Order_items not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Order_items deleted successfully",
	})

}
