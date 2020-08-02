package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/maltea/go-api/car"
	"github.com/maltea/go-api/database"
	_ "github.com/maltea/go-api/docs"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/car", car.GetCars)
	app.Get("/api/car/:id", car.GetCar)
	app.Post("/api/car", car.CreateCar)
	app.Put("/api/car/:id", car.UpdateCar)
	app.Delete("/api/car/:id", car.DeleteCar)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "cars.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	log.Println("Database connection opened")

	database.DBConn.AutoMigrate(&car.Car{})
	log.Println("Automigration completed")
}

// @title Swagger Car API
// @version 1.0
// @description This is a sample golang server with fiber and gorm

// @host localhost:3000
// @basePath /api
func main() {
	app := fiber.New()

	app.Use("/swagger", swagger.Handler)

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	log.Fatal(app.Listen(3000))
}
