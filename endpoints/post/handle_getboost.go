package postEndpoint

import (
	"github.com/gofiber/fiber/v2"
	mod "picme-backend/modules"
	"picme-backend/types/model"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"time"
)

func GetBoostHandler(c *fiber.Ctx) error {
	// * Parse user claims
	//l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Query boostpost
	var boostpost []*model.BoostPostwithImage
	if result := mod.DB.Model(new(table.PostBoost)).
		Select("post_boosts.*, posts.image_url").
		Joins("LEFT JOIN posts ON posts.id = post_boosts.post_id").
		Where("boost_end > ?", time.Now()).Order("RAND()").Find(&boostpost); result.Error != nil {
		return response.Error(false, "Unable to query boostpost", result.Error)
	}

	// * Map table to payload
	//make ถ้าเป็น null จะเป็น array เปล่า
	mappedPosts := make([]*payload.BoostPostResponse, 0)
	for _, post := range boostpost {
		if post.PostId == nil {
			continue
		}

		mappedPosts = append(mappedPosts, &payload.BoostPostResponse{
			PostId:   post.PostId,
			ImageUrl: post.ImageUrl,
		})
	}

	return c.JSON(response.Info(map[string]any{
		"posts": mappedPosts,
	}))
}
