package router

import (
	"github.com/caribbeanyute/freeze-frame/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api")
	api.Get("/thumbnail", handler.GetFrame)

}
