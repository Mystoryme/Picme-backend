package profileEndpoint

//profile page post sortby
import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	mod "picme-backend/modules"
	"picme-backend/types/enum"
	"picme-backend/types/misc"
	"picme-backend/types/model"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"
)

func GridPostGetHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse query
	query := new(payload.ProfileQuery)
	if err := c.QueryParser(query); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate query
	if err := text.Validator.Struct(query); err != nil {
		return err
	}

	// * Query p with sort bysts
	db := mod.DB.Preload("Owner")
	if query.SortBy == nil || *query.SortBy == enum.SortByDate {
		db = db.Where("owner_id =?", l.Id).Order("created_at DESC")
	} else if *query.SortBy == enum.SortByLike {
		db = db.Where("owner_id = ?", l.Id).Order("like_count DESC")
	}

	// * Query posts

	var posts []*model.PostWithCount
	if result := db.Model(new(table.Post)).Select("posts.*, (SELECT COUNT(*) FROM post_likes WHERE post_id = posts.id) AS like_count, (SELECT COUNT(*) FROM post_comments WHERE post_id = posts.id) as comment_count").Find(&posts); result.Error != nil {
		return response.Error(false, "Unable to query posts", result.Error)
	}

	// * Map table to payload
	//make ถ้าเป็น null จะเป็น array เปล่า
	mappedPosts := make([]*payload.GridPostResponse, 0)
	for _, post := range posts {

		mappedPosts = append(mappedPosts, &payload.GridPostResponse{
			ImageUrl:  post.ImageUrl,
			LikeCount: post.LikeCount, // ใส่เพื่อจะได้ดูว่าเรียงตามlikeหรือเปล่า
		})
	}

	return c.JSON(response.Info(map[string]any{
		"posts": mappedPosts,
	}))
}
