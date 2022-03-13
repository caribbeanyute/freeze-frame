package handler

import (
	"github.com/caribbeanyute/freeze-frame/util"
	"github.com/gofiber/fiber/v2"
)

// Hello hanlde api status
func GetFrame(c *fiber.Ctx) error {
	type Stream struct {
		URL string `json:"url"`
	}
	var stream Stream
	if err := c.BodyParser(&stream); err != nil {
		return c.Status(400).SendString("Invalid JSON")
	}

	buf, err := util.CaptureStream(stream.URL)
	if err != nil {
		return c.Status(500).SendString("Failed to get frame")
	}

	//return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
	c.Set("content-type", "image/jpeg")
	return c.Status(200).SendStream(buf)
}
