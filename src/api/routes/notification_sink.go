package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func notification_sink() {

	// Definition and implementation of the /notify endpoint.
	Frouter.Post("/notify", func(c *fiber.Ctx) error {
		// Create the response payload object.
		responsePayload := struct {
			Message string `json:"message"`
		}{}

		// Parse fiber's input from context to responsePayload.
		if err := c.BodyParser(&responsePayload); err != nil {
			return err
		}

		log.Println("responsePayload")
		log.Println(responsePayload)

		// Dummy return of what has already been received.
		return c.JSON(responsePayload)
	})
}
