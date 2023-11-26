package accountEndpoint

import (
	"github.com/gofiber/fiber/v2"
	mod "picme-backend/modules"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/crypto"
	"picme-backend/utils/text"
	"strings"
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

	// * Check Duplicated Email
	var checkDuplicateUser *table.User
	if result := mod.DB.First(&checkDuplicateUser, "email = ?", body.Email); result.Error == nil {
		return response.Error(false, "Duplicated email")
	}

	// * Check that username has more than 4 charactor
	if !strings.Contains(*body.Email, "@") {
		return response.Error(false, "Invalid email address")
	}
	if len(*body.Username) < 4 {
		return response.Error(false, "Required 4 character or more")
	}
	if len(*body.Password) < 4 {
		return response.Error(false, "Required 8 character or more")
	}
	// * Compare Password
	if *body.Password != *body.ConfirmPassword {
		return response.Error(false, "The Password confirmation does not match")
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
		return response.Error(false, "This username already exist")
	}

	// * Response
	return c.JSON(response.Info(
		true,
		&payload.RegisterResponse{
			UserId: user.Id,
		},
	))
}
