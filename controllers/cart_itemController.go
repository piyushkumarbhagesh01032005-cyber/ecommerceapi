package controllers

import (
	"ecommerceapi/database"
	"ecommerceapi/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCartItems(w http.ResponseWriter, r *http.Request) {
	var cart_Item []models.CartItem

	result := database.DB.Find(&cart_Item)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cart_Item)
}

func CreateCartItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cart_Item models.CartItem

	json.NewDecoder(r.Body).Decode(&cart_Item)

	result := database.DB.Create(&cart_Item)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return

	}

	json.NewEncoder(w).Encode(&cart_Item)

}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var cart_Item models.CartItem
	result := database.DB.First(&cart_Item, id)

	if result.Error != nil {
		http.Error(w, "cartitems not found", http.StatusInternalServerError)

		return
	}

	var input models.CartItem
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database.DB.Model(&cart_Item).Updates(map[string]interface{}{
		"product_id": input.ProductID,
		"quantity":   input.Quantity,
	})

	json.NewEncoder(w).Encode(&cart_Item)

}

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, "Invalid cart_item ID", http.StatusBadRequest)
		return
	}

	result := database.DB.Delete(&models.Cart{}, id)
	if result.RowsAffected == 0 {
		http.Error(w, "cart_item not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "cart_item deleted successfully",
	})

}
