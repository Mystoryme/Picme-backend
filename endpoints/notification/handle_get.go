package postEndpoint

//หน้า home (post)
import (
	mod "picme-backend/modules"
	"picme-backend/types/misc"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse query
	query := new(payload.PostQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate query
	if err := text.Validator.Struct(query); err != nil {
		return err
	}

	db := mod.DB.Preload("Trigger").Preload("Triggee").Preload("Post")

	// * Query posts
	var notifications []*table.Notification
	if result := db.Where("triggee_id = ?", l.Id).Find(&notifications); result.Error != nil {
		return response.Error(false, "Unable to query notifications", result.Error)
	}

	// * Map table to payload
	mappedNotifications := make([]*payload.NotificationResponse, 0)
	for _, notification := range notifications {
		if notification.PostId == nil {
			continue
		}
		mappedNotifications = append(mappedNotifications, &payload.NotificationResponse{
			Id: notification.Id,
			Trigger: &payload.ProfileInfo{
				Id:        notification.Trigger.Id,
				Username:  notification.Trigger.Username,
				AvatarUrl: notification.Trigger.AvatarUrl,
			},
			TriggerId: notification.TriggerId,
			Post: &payload.PostResponse{
				PostId: notification.Post.Id},
			PostId:           notification.PostId,
			NotificationType: notification.NotificationType,
			CreatedAt:        notification.CreatedAt,
		})
	}

	return c.JSON(response.Info(map[string]any{
		"notifications": mappedNotifications,
	}))
}
