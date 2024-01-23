package handler

import (
	"errors"

	"github.com/bhill77/todo-golang/config"
	"github.com/bhill77/todo-golang/database"
	"github.com/bhill77/todo-golang/helper"
	"github.com/bhill77/todo-golang/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	conf := config.GetConfig()
	db = database.NewConnection(conf)
}

func IndexHandler(c *fiber.Ctx) error {
	var todos []model.Todo
	db.Find(&todos)
	return c.JSON(todos)
}

func CreateTodoHandler(c *fiber.Ctx) error {
	var payload model.Todo
	err := c.BodyParser(&payload)

	if err != nil {
		c.Status(400).JSON("bad request")
	}

	db.Create(&payload)

	return c.JSON(payload)
}

func DetailTodoHandler(c *fiber.Ctx) error {
	id := helper.StrToUint(c.Params("id"))

	var todo model.Todo
	todo.ID = id

	db.First(&todo)
	return c.JSON(todo)
}

func UpdateTodoHandler(c *fiber.Ctx) error {
	id := helper.StrToUint(c.Params("id"))
	var todo model.Todo
	todo.ID = id

	result := db.First(&todo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON("Invalid ID")
	}

	var payload model.Todo
	err1 := c.BodyParser(&payload)

	if err1 != nil {
		c.Status(400).JSON("bad request")
	}

	todo.Label = payload.Label
	todo.Finish = payload.Finish

	db.Save(&todo)

	return c.JSON(todo)
}

func DeleteTodoHandler(c *fiber.Ctx) error {
	id := helper.StrToUint(c.Params("id"))
	var todo model.Todo
	todo.ID = id

	result := db.First(&todo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON("Invalid ID")
	}

	db.Delete(&todo)

	return c.JSON("data berhasil dihapus")
}
