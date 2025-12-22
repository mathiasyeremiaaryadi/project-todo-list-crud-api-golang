package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewRoute(app *fiber.App) {
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 30 * time.Second,
	}))

	app.Post("/register", RegisterHandler)
	app.Post("/login", LoginHandler)
	app.Post("/refresh-token", RefreshTokenHandler)

	app.Post("/todos", JWTMiddleware, CreateHandler)
	app.Put("/todos/:id", JWTMiddleware, UpdateHandler)
	app.Delete("/todos/:id", JWTMiddleware, DeleteHandler)
	app.Get("/todos", JWTMiddleware, GetHandler)
}
