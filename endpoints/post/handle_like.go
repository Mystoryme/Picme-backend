package postEndpoint

import (
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

func LikeHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.CreateLikeBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}
	//มีไ where user_id ด้วย เพราะจะนับ likecount ของคนนั้น ถ้าเป็ฯ 0 จะlike ได้ แต่ถ้าเป็น 1 แล้วจะlike ไม่ได้ละ
	var likeCount int64
	if result := mod.DB.Model(new(table.PostLike)).Where("post_id = ? and user_id = ?", body.PostId, l.Id).Count(&likeCount); result.Error != nil {
		return response.Error(false, "Unable to count likes", result.Error)
	}
	//create
	//ctr+space เติมfillแบบรวดเร็ว
	if likeCount == 0 {
		like := &table.PostLike{
			User:      nil,
			UserId:    l.Id,
			Post:      nil,
			PostId:    body.PostId,
			CreatedAt: nil,
		}

		if result := mod.DB.Create(like); result.Error != nil {
			return response.Error(false, "Unable to like post", result.Error)
		}
	}
	currentPost := new(table.Post)

	if result := mod.DB.Preload("Owner").First(currentPost, "id = ?", body.PostId); result.Error != nil {
		return response.Error(false, "Unable to query post", result.Error)
	}

	go func() {
		likeType := enum.NotificationLike
		insightLikeType := enum.InsightLike
		notification := &table.Notification{
			Trigger:          nil,
			TriggerId:        l.Id,
			Triggee:          nil,
			TriggeeId:        currentPost.OwnerId,
			Post:             nil,
			PostId:           body.PostId,
			NotificationType: &likeType,
			CreatedAt:        nil,
		}

		if result := mod.DB.Create(&notification); result.Error != nil {
			fmt.Printf("Unable to create notification %v", result.Error)
		}

		if result := mod.DB.Create(&table.Insight{
			Trigger:     nil,
			TriggerId:   l.Id,
			Triggee:     nil,
			TriggeeId:   currentPost.OwnerId,
			InsightType: &insightLikeType,
			CreatedAt:   nil,
		}); result.Error != nil {
			fmt.Printf("Unable to create insight %v", result.Error)
		}
	}()

	return c.JSON(response.Info("Successfully like!"))
}
