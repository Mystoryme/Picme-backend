package postEndpoint

import (
	// ... (import statements)

	"fmt"
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

func ViewProfile(c *fiber.Ctx) error {
	// Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// Parse body
	body := new(payload.ViewProfileBody)
	if err := c.BodyParser(body); err != nil || body.UserId == nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// Validate query
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	go func() {
		insightViewType := enum.InsightView

		if l.Id == body.UserId {
			return
		}

		// check if there's existing insight
		if result := mod.DB.Table("insights").Where("trigger_id = ? AND triggee_id = ? AND insight_type = ?", l.Id, body.UserId, insightViewType).First(&table.Insight{}); result.Error == nil {
			return
		}

		if result := mod.DB.Create(&table.Insight{
			Trigger:     nil,
			TriggerId:   l.Id,
			Triggee:     nil,
			TriggeeId:   body.UserId,
			InsightType: &insightViewType,
			CreatedAt:   nil,
		}); result.Error != nil {
			fmt.Printf("Unable to create insight %v", result.Error)
		}
	}()

	return c.JSON(response.Info("Successfully profile view!"))
}
