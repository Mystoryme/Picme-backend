package profileEndpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	mod "picme-backend/modules"
	"picme-backend/types/misc"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"
)

func EditHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.EditBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	//query user by userid
	var user table.User
	if result := mod.DB.Where("id = ?", l.Id).First(&user); result.Error != nil {
		return response.Error(false, "Unable to fetch user profile", result.Error)
	}

	//check ว่าเหมือนอันเดิมไหม
	if body.Username != user.Username && *body.Username != "" {
		user.Username = body.Username
	}

	if body.Contact != user.Contact && *body.Contact != "" {
		user.Contact = body.Contact
	}

	if body.Bio != user.Bio && *body.Bio != "" {
		user.Bio = body.Bio
	}

	if result := mod.DB.Save(&user); result.Error != nil {
		return response.Error(false, "This username already exist", result.Error)
	}

	return c.JSON(response.Info("Successfully update!"))
}
