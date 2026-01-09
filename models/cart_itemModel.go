package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    int     `json:"cart_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
