package profileEndpoint

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
	body := new(payload.DonateBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	//create
	//ctr+space เติมfillแบบรวดเร็ว

	donate := &table.UserDonate{
		Donor:        nil,
		DonorId:      l.Id,
		User:         nil,
		UserId:       body.UserId,
		DonateAmount: body.DonateAmount,
		CreatedAt:    nil,
	}

	if result := mod.DB.Create(donate); result.Error != nil {
		return response.Error(false, "Unable to donate artist", result.Error)
	}

	userDonateType := enum.NotificationUserDonate
	notification := &table.Notification{
		Trigger:          nil,
		TriggerId:        l.Id,
		Post:             nil,
		PostId:           nil,
		Triggee:          nil,
		TriggeeId:        body.UserId,
		NotificationType: &userDonateType,
		CreatedAt:        nil,
	}

	if result := mod.DB.Create(&notification); result.Error != nil {
		return response.Error(false, "Unable to create notification", result.Error)
	}

	return c.JSON(response.Info("Successfully donate!"))
}
