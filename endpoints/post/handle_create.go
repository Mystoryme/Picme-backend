package postEndpoint

import (
	"github.com/gofiber/fiber/v2"
	"picme-backend/types/payload"
	"picme-backend/types/response"
)

func CreateHandler(c *fiber.Ctx) error {
	body := new(payload.CreatePostBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// TODO: Database operation

	return c.JSON(response.Info("Successfully posted!"))
}
