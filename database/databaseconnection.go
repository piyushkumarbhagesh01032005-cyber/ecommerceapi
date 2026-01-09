package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ecommerceapi/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=piyush@6395 dbname=ecommerce port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.Cart{},
		&models.CartItem{},
	)

	fmt.Println("Database connected successfully!")
}
