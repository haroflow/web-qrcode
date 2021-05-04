package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/skip2/go-qrcode"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Post("/generate", func(c *fiber.Ctx) error {
		type request struct {
			Data string
			Size int
		}

		var r request
		if err := c.BodyParser(&r); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Failed to parse request body: "+err.Error())
		}

		img, err := qrcode.Encode(r.Data, qrcode.Medium, r.Size)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate qrcode: "+err.Error())
		}

		c.Set("Content-Type", "image/png")

		return c.Send(img)
	})

	app.Static("/", "./static/")

	log.Fatal(app.Listen(":" + port))
}
