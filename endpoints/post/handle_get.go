package postEndpoint

import (
	"github.com/gofiber/fiber/v2"
	mod "picme-backend/modules"
	"picme-backend/types/model"
	"picme-backend/types/payload"
	"picme-backend/types/response"
)

func GetHandler(c *fiber.Ctx) error {
	// * Query posts
	var posts []*model.Posts
	if result := mod.DB.Preload("Owner").Find(&posts); result.Error != nil {
		return response.Error(false, "Unable to query posts", result.Error)
	}

	// * Map table to payload
	var mappedPosts []*payload.PostResponse
	for _, post := range posts {
		var likeCount uint64 = 0
		var commentCount uint64 = 0
		mappedPosts = append(mappedPosts, &payload.PostResponse{
			PostId:        post.Id,
			OwnerId:       post.OwnerId,
			OwnerUsername: post.Owner.Username,
			Caption:       post.Caption,
			ImageUrl:      post.ImageUrl,
			LikeCount:     &likeCount,
			CommentCount:  &commentCount,
		})
	}

	return c.JSON(response.Info(map[string]any{
		"posts": mappedPosts,
	}))
}
