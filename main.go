package main

import (
	"log"

	"github.com/caribbeanyute/freeze-frame/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Fiber instance
	app := fiber.New()
	app.Use(cors.New())
	router.SetupRoutes(app)
	

	// Start server
	log.Fatal(app.Listen(":3001"))
}

// Handler
/*
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}*/
