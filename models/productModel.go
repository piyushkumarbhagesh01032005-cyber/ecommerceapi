package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `jsom:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
