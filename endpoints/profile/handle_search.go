package profileEndpoint

import (
	"github.com/gofiber/fiber/v2"
	mod "picme-backend/modules"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"
)

func SearchHandler(c *fiber.Ctx) error {
	// * Parse user claims
	//l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.SearchBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	// * Query posts
	var users []*table.User
	if result := mod.DB.Where("username LIKE ?", "%"+*body.Username+"%").Limit(5).Find(&users); result.Error != nil {
		return response.Error(false, "Unable to query posts", result.Error)
	}

	// * Map table to payload
	mappedUsers := make([]*payload.SearchResponse, 0)
	for _, user := range users {

		mappedUsers = append(mappedUsers, &payload.SearchResponse{
			UserId:    user.Id,
			Username:  user.Username,
			AvatarUrl: user.AvatarUrl,
		})
	}

	return c.JSON(response.Info(map[string]any{
		"users": mappedUsers,
	}))
}
