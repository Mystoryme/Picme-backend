package profileEndpoint

import (
	"github.com/gofiber/fiber/v2"
	mod "picme-backend/modules"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"
)

func ProfileSearchGetHandler(c *fiber.Ctx) error {
	// * Parse user claims
	//l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.ProfileSearchBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	// * Query user by id
	var user *table.User
	if result := mod.DB.Where("id = ?", body.UserId).First(&user); result.Error != nil {
		return response.Error(false, "Unable to query user", result.Error)
	}

	// * Construct user
	profile := &payload.ProfileInfo{
		Id:        user.Id,
		Username:  user.Username,
		Bio:       user.Bio,
		Contact:   user.Contact,
		AvatarUrl: user.AvatarUrl,
	}

	// * Response
	return c.JSON(response.Info(profile))
}
