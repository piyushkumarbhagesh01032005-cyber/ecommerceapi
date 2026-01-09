package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID        int         `json:"user_id"`
	TotalAmount   float64     `json:"total_amount"`
	PaymentStatus string      `json:"payment_status"`
	OrderItems    []OrderItem `json:"order_items"`
}
