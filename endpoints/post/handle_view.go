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

func ViewHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.ViewBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	//create
	//ctr+space เติมfillแบบรวดเร็ว

	view := &table.PostViews{
		User:      nil,
		UserId:    l.Id,
		Post:      nil,
		PostId:    body.PostId,
		CreatedAt: nil,
	}

	if result := mod.DB.Create(view); result.Error != nil {
		return response.Error(false, "Unable to view post", result.Error)
	}

	go func() {
		insightLikeType := enum.InsightLike
		currentPost := new(table.Post)

		if result := mod.DB.Preload("Owner").First(currentPost, "id = ?", body.PostId); result.Error != nil {
			fmt.Printf("Unable to query post %v", result.Error)
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

	return c.JSON(response.Info("Successfully view!"))
}
