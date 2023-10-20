package postEndpoint

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

func CreateHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.CreatePostBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}
	//create
	//ctr+space เติมfillแบบรวดเร็ว
	posts := &table.Post{
		Id:          nil,
		ImageUrl:    body.ImageUrl,
		Owner:       nil,
		OwnerId:     l.Id,
		Caption:     body.Caption,
		Category:    body.Category,
		Application: body.Application,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}

	if result := mod.DB.Create(posts); result.Error != nil {
		return response.Error(false, "Unable to create user", result.Error)
	}

	return c.JSON(response.Info("Successfully posted!"))
}
