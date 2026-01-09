package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerceapi/database"
	"ecommerceapi/models"

	"github.com/gorilla/mux"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var orders []models.Order
	result := database.DB.Find(&orders)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var order models.Order
	json.NewDecoder(r.Body).Decode(&order)

	result := database.DB.Create(&order)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var order models.Order
	result := database.DB.First(&order, id)
	if result.Error != nil {
		http.Error(w, "Order Not Found", http.StatusNotFound)

		return
	}

	var input models.Order

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	database.DB.Model(&order).Updates(map[string]interface{}{
		"user_id":        input.UserID,
		"total_amount":   input.TotalAmount,
		"payment_status": input.PaymentStatus,
	})
	json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid Order ID", http.StatusBadRequest)
		return
	}

	result := database.DB.Delete(&models.Order{}, id)

	if result.RowsAffected == 0 {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Order deleted successfully",
	})
}
