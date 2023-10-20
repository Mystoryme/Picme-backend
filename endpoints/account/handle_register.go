package accountEndpoint

import (
	"github.com/gofiber/fiber/v2"
	mod "picme-backend/modules"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/crypto"
	"picme-backend/utils/text"
)

func RegisterHandler(c *fiber.Ctx) error {
	// * Parse body
	body := new(payload.RegisterRequest)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	// * Hash password
	hashedPassword, err := crypto.HashPassword(*body.Password)
	if err != nil {
		return response.Error(false, "Unable to hash password", err)
	}

	// * Construct user row
	user := &table.User{
		Id:        nil,
		Username:  body.Username,
		Email:     body.Email,
		Password:  &hashedPassword,
		Bio:       nil,
		Contact:   nil,
		AvatarUrl: nil,
		CreatedAt: nil,
		UpdatedAt: nil,
	}

	// * Create user row
	if result := mod.DB.Create(user); result.Error != nil {
		return response.Error(false, "Unable to create user", result.Error)
	}

	// * Response
	return c.JSON(response.Info(
		true,
		&payload.RegisterResponse{
			UserId: user.Id,
		},
	))
}
