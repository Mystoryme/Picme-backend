package postEndpoint

//หน้า home (post)
import (
	"github.com/gofiber/fiber/v2"
	mod "picme-backend/modules"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"
)

func GetHandler(c *fiber.Ctx) error {
	// * Parse user claims
	//l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse query
	query := new(payload.PostQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate query
	if err := text.Validator.Struct(query); err != nil {
		return err
	}

	db := mod.DB.Preload("Owner")
	if query.Category != nil {
		db = db.Where("category = ?", query.Category)
	}

	// * Query posts
	var posts []*table.Post
	if result := db.Order("RAND()").Find(&posts); result.Error != nil {
		return response.Error(false, "Unable to query posts", result.Error)
	}

	// * Map table to payload
	mappedPosts := make([]*payload.PostResponse, 0)
	for _, post := range posts {
		var likeCount int64
		if result := mod.DB.Model(&table.PostLike{}).Where("post_id = ?", post.Id).Count(&likeCount); result.Error != nil {
			return response.Error(false, "Unable to count likes", result.Error)

		}
		//&table.PostLike{} == new(table.PostComment)
		var commentCount int64
		if result := mod.DB.Model(new(table.PostComment)).Where("post_id = ?", post.Id).Count(&commentCount); result.Error != nil {
			return response.Error(false, "Unable to count comments", result.Error)
		}
		mappedPosts = append(mappedPosts, &payload.PostResponse{
			PostId:        post.Id,
			OwnerId:       post.OwnerId,
			OwnerUsername: post.Owner.Username,
			Caption:       post.Caption,
			ImageUrl:      post.ImageUrl,
			Application:   post.Application,
			LikeCount:     &likeCount,
			CommentCount:  &commentCount,
		})
	}

	return c.JSON(response.Info(map[string]any{
		"posts": mappedPosts,
	}))
}
