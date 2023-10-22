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

func DeleteHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.DeleteCommentRequest)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	//delete
	var comment *table.PostComment
	if result := mod.DB.Where("id = ? and user_id = ? ", body.Id, l.Id).Delete(&comment); result.Error != nil {
		return response.Error(false, "Unable to delete the comment", result.Error)
	}

	return c.JSON(response.Info("Successfully delete comment!"))
}
