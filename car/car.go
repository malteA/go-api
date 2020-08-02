package car

import (
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber"
)

// Car struct
type Car struct {
	ID    string `json:id`
	Brand string `json:brand`
	Make  string `json:make`
}

var cars []Car

func SetupMockdData() {
	// mock data
	cars = append(cars, Car{ID: "1", Brand: "Tesla", Make: "Model S"})
	cars = append(cars, Car{ID: "2", Brand: "VW", Make: "Golf"})
}

func GetCars(c *fiber.Ctx) {
	c.JSON(cars)
}

func GetCar(c *fiber.Ctx) {
	id := c.Params("id")
	for _, item := range cars {
		if item.ID == id {
			c.JSON(item)
			return
		}
	}
}

func CreateCar(c *fiber.Ctx) {
	var car Car
	_ = c.BodyParser(&car)
	car.ID = strconv.Itoa(rand.Intn(1000))
	cars = append(cars, car)
	c.JSON(car)
}

func UpdateCar(c *fiber.Ctx) {
	id := c.Params("id")
	for index, item := range cars {
		if item.ID == id {
			cars = append(cars[:index], cars[index+1:]...)
			var car Car
			_ = c.BodyParser(car)
			car.ID = id
			cars = append(cars, car)
			c.JSON(car)
			return
		}
	}
}

func DeleteCar(c *fiber.Ctx) {
	id := c.Params("id")
	for index, item := range cars {
		if item.ID == id {
			cars = append(cars[:index], cars[index+1:]...)
			break
		}
	}
	c.JSON(cars)
}
