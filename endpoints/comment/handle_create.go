package comment

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
	body := new(payload.CreateCommentBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}
	//create
	//ctr+space เติมfillแบบรวดเร็ว
	comment := &table.PostComment{
		User:      nil,
		UserId:    l.Id,
		Post:      nil,
		PostId:    body.PostId,
		Message:   body.Message,
		CreatedAt: nil,
		UpdatedAt: nil,
	}

	if result := mod.DB.Create(comment); result.Error != nil {
		return response.Error(false, "Unable to comment post", result.Error)
	}

	return c.JSON(response.Info("Successfully comment!"))
}
