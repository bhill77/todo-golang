package main

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:123456@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	type Todo struct {
		gorm.Model
		Label  string
		Finish bool
	}

	db.AutoMigrate(&Todo{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Get("/todo", func(c *fiber.Ctx) error {
		var todos []Todo
		db.Find(&todos)
		return c.JSON(todos)
	})

	app.Post("/todo", func(c *fiber.Ctx) error {
		var payload Todo
		err := c.BodyParser(&payload)

		if err != nil {
			c.Status(400).JSON("bad request")
		}

		db.Create(&payload)

		return c.JSON(payload)
	})

	app.Get("/todo/:id", func(c *fiber.Ctx) error {
		strId := c.Params("id")

		id, _ := strconv.ParseUint(strId, 10, 32)
		var todo Todo
		todo.ID = uint(id)

		db.First(&todo)
		return c.JSON(todo)
	})

	app.Put("/todo/:id", func(c *fiber.Ctx) error {
		strId := c.Params("id")

		id, _ := strconv.ParseUint(strId, 10, 32)
		var todo Todo
		todo.ID = uint(id)

		result := db.First(&todo)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON("Invalid ID")
		}

		var payload Todo
		err1 := c.BodyParser(&payload)

		if err1 != nil {
			c.Status(400).JSON("bad request")
		}

		todo.Label = payload.Label
		todo.Finish = payload.Finish

		db.Save(&todo)

		return c.JSON(todo)
	})

	app.Delete("/todo/:id", func(c *fiber.Ctx) error {
		strId := c.Params("id")

		id, _ := strconv.ParseUint(strId, 10, 32)
		var todo Todo
		todo.ID = uint(id)

		result := db.First(&todo)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON("Data tidak ditemukan")
		}

		db.Delete(&todo)

		return c.JSON("data berhasil dihapus")
	})

	app.Listen(":3000")
}
