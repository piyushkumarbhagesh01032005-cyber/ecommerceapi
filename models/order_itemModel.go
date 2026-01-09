package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Product   Product `gorm:"foreginKey:ProductID"`
}
