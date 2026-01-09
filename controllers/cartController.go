package controllers

import (
	"ecommerceapi/database"
	"ecommerceapi/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var carts []models.Cart

	result := database.DB.Find(&carts)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(carts)
}

func CreateCart(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var cart models.Cart
	json.NewDecoder(r.Body).Decode(&cart)

	result := database.DB.Create(&cart)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cart)

}

func UpdateCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var cart models.Cart

	result := database.DB.First(&cart, id)

	if result.Error != nil {
		http.Error(w, "cart is not found", http.StatusInternalServerError)

		return
	}
	var input models.Cart
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Model(&cart).Updates(map[string]interface{}{
		"user_id": input.UserID,
	})

	json.NewEncoder(w).Encode(&cart)

}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid cart ID", http.StatusBadRequest)
		return
	}

	result := database.DB.Delete(&models.Cart{}, id)
	if result.RowsAffected == 0 {
		http.Error(w, "cart not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "cart deleted successfully",
	})

}
