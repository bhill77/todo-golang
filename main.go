package main

import (
	"fmt"

	"github.com/bhill77/todo-golang/config"
	"github.com/bhill77/todo-golang/database"
	"github.com/bhill77/todo-golang/handler"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Label  string
	Finish bool
}

var db *gorm.DB

func main() {

	conf := config.GetConfig()
	db := database.NewConnection(conf)

	// run auto migrate
	db.AutoMigrate(&Todo{})

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Get("/todo", handler.IndexHandler)

	app.Post("/todo", handler.CreateTodoHandler)

	app.Get("/todo/:id", handler.DetailTodoHandler)

	app.Put("/todo/:id", handler.UpdateTodoHandler)

	app.Delete("/todo/:id", handler.DeleteTodoHandler)

	app.Listen(fmt.Sprintf(":%d", conf.Port))
}
