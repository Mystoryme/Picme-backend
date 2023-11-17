package comment

import (
	mod "picme-backend/modules"
	"picme-backend/types/enum"
	"picme-backend/types/misc"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
		Id:        nil,
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

	postOwner := new(table.User)
	if result := mod.DB.First(postOwner, "id = ?", body.PostId); result.Error != nil {
		return response.Error(false, "Unable to query post owner", result.Error)
	}

	commentType := enum.NotificationComment
	notification := &table.Notification{
		Trigger:          nil,
		TriggerId:        l.Id,
		Post:             nil,
		PostId:           body.PostId,
		Triggee:          nil,
		TriggeeId:        postOwner.Id,
		NotificationType: &commentType,
		CreatedAt:        nil,
	}

	if result := mod.DB.Create(&notification); result.Error != nil {
		return response.Error(false, "Unable to create notification", result.Error)
	}

	return c.JSON(response.Info("Successfully comment!"))
}
