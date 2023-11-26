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

	db := mod.DB.Preload("Trigger").Preload("Triggee").Preload("Post").Order("created_at desc")

	// * Query posts
	var notifications []*table.Notification
	if result := db.Unscoped().Where("triggee_id = ?", l.Id).Find(&notifications); result.Error != nil {
		return response.Error(false, "Unable to query notifications", result.Error)
	}

	// * Map table to payload
	mappedNotifications := make([]*payload.NotificationResponse, 0)
	for _, notification := range notifications {
		if *notification.TriggerId == *l.Id {
			continue
		}

		pl := &payload.NotificationResponse{
			Id: notification.Id,
			Trigger: &payload.ProfileInfo{
				Id:        notification.Trigger.Id,
				Username:  notification.Trigger.Username,
				AvatarUrl: notification.Trigger.AvatarUrl,
			},
			TriggerId:        notification.TriggerId,
			Post:             nil,
			PostId:           nil,
			NotificationType: notification.NotificationType,
			CreatedAt:        notification.CreatedAt,
		}

		if notification.PostId == nil {
			mappedNotifications = append(mappedNotifications, pl)
			continue
		}

		pl.Post = &payload.PostResponse{
			PostId: notification.Post.Id,
		}
		pl.PostId = notification.PostId
		mappedNotifications = append(mappedNotifications, pl)
	}

	return c.JSON(response.Info(map[string]any{
		"notifications": mappedNotifications,
	}))
}
