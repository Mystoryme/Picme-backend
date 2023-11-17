package postEndpoint

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

func DonateHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.CreateDonateBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	//create
	//ctr+space เติมfillแบบรวดเร็ว

	donate := &table.PostDonate{
		User:         nil,
		UserId:       l.Id,
		Post:         nil,
		PostId:       body.PostId,
		DonateAmount: body.DonateAmount,
		CreatedAt:    nil,
	}

	if result := mod.DB.Create(donate); result.Error != nil {
		return response.Error(false, "Unable to donate post", result.Error)
	}

	postOwner := new(table.User)
	if result := mod.DB.First(postOwner, "post_id = ?", body.PostId); result.Error != nil {
		return response.Error(false, "Unable to query post owner", result.Error)
	}

	postDonateType := enum.NotificationPostDonate
	notification := &table.Notification{
		Trigger:          nil,
		TriggerId:        l.Id,
		Triggee:          nil,
		TriggeeId:        postOwner.Id,
		Post:             nil,
		PostId:           body.PostId,
		NotificationType: &postDonateType,
		CreatedAt:        nil,
	}

	if result := mod.DB.Create(&notification); result.Error != nil {
		return response.Error(false, "Unable to create notification", result.Error)
	}

	return c.JSON(response.Info("Successfully donate!"))
}
