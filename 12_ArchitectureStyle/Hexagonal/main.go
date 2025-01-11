package main

import (
	"HexagonalExample/core"
	"HexagonalExample/adapters"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

func main() {
	app := fiber.New()
	
	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&core.Order{})

	// Secondary Port
	orderRepo := adapters.NewGormOrderRepository(db)
	// ส่ง Secondary Port เข้าไปใน Application Core
	// Primary Port
	orderService := core.NewOrderService(orderRepo)
	// Primary Adapter
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	app.Listen(":8080")
}