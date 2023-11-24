package postEndpoint

import (
	// ... (import statements)

	mod "picme-backend/modules"
	"picme-backend/types/misc"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetViewHistory(c *fiber.Ctx) error {
	// Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	insights := make([]table.Insight, 0)
	insightsResponse := make([]payload.InsightResponse, 0)
	mod.DB.Table("insights").Preload("Triggee").Preload("Trigger").Where("trigger_id = ? AND insight_type = 'view'", l.Id).Limit(4).Find(&insights)

	for _, insight := range insights {
		insightsResponse = append(insightsResponse, payload.InsightResponse{
			Trigger: &payload.ProfileInfo{
				Id:        insight.Trigger.Id,
				Username:  insight.Trigger.Username,
				AvatarUrl: insight.Trigger.AvatarUrl,
			},
			TriggerId: insight.TriggerId,
			Triggee: &payload.ProfileInfo{
				Id:        insight.Triggee.Id,
				Username:  insight.Triggee.Username,
				AvatarUrl: insight.Triggee.AvatarUrl,
			},
			TriggeeId:   insight.TriggeeId,
			InsightType: insight.InsightType,
			CreatedAt:   insight.CreatedAt,
		})
	}

	return c.JSON(response.Info(map[string]interface{}{
		"history": insightsResponse,
	}))
}
