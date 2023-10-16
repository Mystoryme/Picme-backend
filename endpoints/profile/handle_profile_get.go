package profileEndpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	mod "picme-backend/modules"
	"picme-backend/types/misc"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
)

func ProfileGetHandler(c *fiber.Ctx) error {
	// * Parse user JWT cookie
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Query user by id
	var user *table.Users
	if result := mod.DB.Where("id = ?", l.Id).First(&user); result.Error != nil {
		return response.Error(false, "Unable to query user", result.Error)
	}

	// * Construct user
	profile := &payload.ProfileInfo{
		Username: user.Username,
		//Name:     user.Name,
	}

	// * Response
	return c.JSON(response.Info(true, profile))
}
