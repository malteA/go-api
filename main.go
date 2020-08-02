package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/maltea/go-api/car"
	"github.com/maltea/go-api/database"
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

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	log.Fatal(app.Listen(3000))
}
