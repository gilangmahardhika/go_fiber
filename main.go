package main

import (
	"github.com/gilangmahardhika/golang-web/controllers/bookcontroller"
	"github.com/gilangmahardhika/golang-web/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()
	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", bookcontroller.Index)
	book.Get("/:id", bookcontroller.Show)
	book.Post("", bookcontroller.Create)
	book.Put("/:id", bookcontroller.Update)
	book.Delete("/:id", bookcontroller.Delete)

	app.Listen(":8080")
}
