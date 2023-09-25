package accountEndpoint

import (
	"github.com/gofiber/fiber/v2"
	"picme-backend/types/payload"
	"picme-backend/types/response"
)

func LoginHandler(c *fiber.Ctx) error {
	return c.JSON(response.Success(&payload.LoginResponse{
		Token: "111",
		Email: "aim@aim.com",
	}))
}
