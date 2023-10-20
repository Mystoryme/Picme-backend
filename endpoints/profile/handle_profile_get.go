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
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Query user by id
	var user *table.User
	if result := mod.DB.Where("id = ?", l.Id).First(&user); result.Error != nil {
		return response.Error(false, "Unable to query user", result.Error)
	}

	// * Construct user
	profile := &payload.ProfileInfo{
		Id:        l.Id,
		Username:  user.Username,
		Bio:       user.Bio,
		Contact:   user.Contact,
		AvatarUrl: user.AvatarUrl,
	}

	// * Response
	return c.JSON(response.Info(profile))
}
