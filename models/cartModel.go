package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    int        `json:"user_id"`
	CartItems []CartItem `json:"cart_items"`
}
