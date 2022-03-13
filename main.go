// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"

	"github.com/caribbeanyute/freeze-frame/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiber instance
	app := fiber.New()

	router.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3001"))
}

// Handler
/*
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}*/
