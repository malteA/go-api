package car

import (
	"strconv"

	"github.com/maltea/go-api/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// Car struct
type Car struct {
	gorm.Model
	Brand string `json:"brand"`
	Make  string `json:"make"`
}

// GetCars
func GetCars(c *fiber.Ctx) {
	db := database.DBConn
	var cars []Car
	db.Find(&cars)
	c.JSON(cars)
}

// GetCar
func GetCar(c *fiber.Ctx) {
	id := c.Params("id")
	if id == "" {
		c.SendStatus(400)
		return
	}
	db := database.DBConn
	var car Car
	db.Find(&car, id)
	c.JSON(car)
}

func CreateCar(c *fiber.Ctx) {
	var car Car
	err := c.BodyParser(&car)
	if err != nil {
		c.SendStatus(400)
		return
	}
	db := database.DBConn
	db.Create(&car)
	c.JSON(car)
}

// UpdateCar
func UpdateCar(c *fiber.Ctx) {
	id := c.Params("id")
	_, parseErr := strconv.Atoi(id)
	if (id == "") || (parseErr != nil) {
		c.SendStatus(400)
		return
	}

	var carUp Car
	err := c.BodyParser(&carUp)
	if err != nil {
		c.SendStatus(400)
		return
	}

	var car Car
	db := database.DBConn
	db.Find(&car, id)
	db.Model(&car).Update(&carUp)
	c.JSON(carUp)
}

// DeleteCar
func DeleteCar(c *fiber.Ctx) {
	id := c.Params("id")
	if id == "" {
		c.SendStatus(400)
		return
	}

	db := database.DBConn
	var car Car
	db.Find(&car, id)

	db.Delete(&car, id)
	c.SendStatus(200)
}
