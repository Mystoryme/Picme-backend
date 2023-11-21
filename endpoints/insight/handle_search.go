package postEndpoint

import (
	// ... (import statements)

	mod "picme-backend/modules"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/utils/text"

	"github.com/gofiber/fiber/v2"
)

func GetSearch(c *fiber.Ctx) error {

	// Parse body
	body := new(payload.SearchQuery)
	if err := c.BodyParser(body); err != nil || body.Username == nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	users := make([]payload.ProfileInfo, 0)
	if err := mod.DB.Table("users").Where("username LIKE ?", "%"+*body.Username+"%").Find(&users).Error; err != nil {
		return response.Error(false, "Unable to find users", err)
	}

	return c.JSON(response.Info(map[string]interface{}{
		"users": users,
	}))
}
