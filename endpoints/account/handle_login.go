package accountEndpoint

import (
	"github.com/gofiber/fiber/v2"
	"picme-backend/modules"
	"picme-backend/types/misc"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/crypto"
	"picme-backend/utils/text"
)

func LoginHandler(c *fiber.Ctx) error {
	// * Parse body
	body := new(payload.LoginRequest)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	// * Query user by username
	var user *table.User
	if result := mod.DB.Where("email = ?", body.Email).First(&user); result.Error != nil {
		return response.Error(false, "Unable to query user", result.Error)
	}

	// * Check password
	if crypto.ComparePassword(*body.Password, *user.Password) {
		return response.Error(false, "Incorrect password")
	}

	// # At this point, user is authenticated

	// * Construct JWT claim
	claim := &misc.UserClaim{
		Id: user.Id,
	}

	// * Sign claim
	token, err := crypto.SignJwt(claim)
	if err != nil {
		return response.Error(false, "Unable to sign token", err)
	}

	// * Response
	return c.JSON(response.Info(map[string]any{
		"token": token,
		//"test":  222,
	}))
}
