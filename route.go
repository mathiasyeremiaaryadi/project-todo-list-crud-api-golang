package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewRoute(app *fiber.App) {
	app.Use(logger.New())

	app.Post("/register", RegisterHandler)
	app.Post("/login", LoginHandler)

	app.Post("/todos", CreateHandler)
	app.Put("/todos/:id", UpdateHandler)
	app.Delete("/todos/:id", DeleteHandler)
	app.Get("/todos", GetHandler)
}
