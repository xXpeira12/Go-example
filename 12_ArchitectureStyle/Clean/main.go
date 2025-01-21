package main

import (
	"CleanExample/adapters"
	"CleanExample/entities"
	"CleanExample/usecases"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	app := fiber.New()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&entities.Order{})

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := usecases.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)
	
	app.Post("/order", orderHandler.CreateOrder)

	app.Listen(":8000")
}