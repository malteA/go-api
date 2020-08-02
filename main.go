package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/maltea/go-api/car"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/car", car.GetCars)
	app.Get("/api/car/:id", car.GetCar)
	app.Post("/api/car", car.CreateCar)
	app.Put("/api/car/:id", car.UpdateCar)
	app.Delete("/api/car/:id", car.DeleteCar)
}

func main() {
	// iinit router
	app := fiber.New()

	car.SetupMockdData()

	setupRoutes(app)

	log.Fatal(app.Listen(3000))
}
