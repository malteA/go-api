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

// GetCars godoc
// @Summary List cars
// @Description get cars
// @Tags cars
// @Produce json
// @Success 200 {array} Car
// @Router /car [get]
func GetCars(c *fiber.Ctx) {
	db := database.DBConn
	var cars []Car
	db.Find(&cars)
	c.JSON(cars)
}

// GetCar gooc
// @Summary Gets all cars
// @Description get car by ID
// @Tags cars
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Success 200 {object} Car
// @Failure 400
// @Router /car/{id} [get]
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

// CreateCar gooc
// @Summary Create a car
// @Description create a json car
// @Tags cars
// @Accept json
// @Produce json
// @Param car body Car true "Add car"
// @Success 200 {object} Car
// @Failure 400
// @Router /car [post]
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

// UpdateCar gooc
// @Summary Update a car
// @Description update a json car
// @Tags cars
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Param car body Car true "Add car"
// @Success 200 {object} Car
// @Failure 400
// @Router /car/{id} [put]
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

// DeleteCar gooc
// @Summary Delete a car
// @Description delete a car by ID
// @Tags cars
// @Accept json
// @Produce json
// @Param id path string true "Car ID"
// @Success 200
// @Failure 400
// @Router /car/{id} [delete]
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
