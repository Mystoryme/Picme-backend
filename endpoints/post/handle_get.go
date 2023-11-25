package postEndpoint

//หน้า home (post)
import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	mod "picme-backend/modules"
	"picme-backend/types/misc"
	"picme-backend/types/model"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"
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

	db := mod.DB.Preload("Owner")
	if query.Category != nil {
		db = db.Where("category = ?", query.Category)
	}

	// * Query posts
	var posts []*model.PostWithCount
	if result := db.Model(new(table.Post)).Select("posts.*, (SELECT COUNT(*) FROM post_likes WHERE post_id = posts.id) AS like_count, (SELECT COUNT(*) FROM post_comments WHERE post_id = posts.id) as comment_count,"+"(SELECT COUNT(*) FROM post_likes WHERE post_id = posts.id AND user_id = ?) AS liked, "+
		"(SELECT COUNT(*) FROM post_book_marks WHERE post_id = posts.id AND user_id = ?) AS booked", l.Id, l.Id).Order("RAND()").Find(&posts); result.Error != nil {
		return response.Error(false, "Unable to query posts", result.Error)
	}

	for _, post := range posts {
		if *post.Liked == 1 {
			post.IsLiked = true
		} else {
			post.IsLiked = false
		}
		if *post.Booked == 1 {
			post.IsBooked = true
		} else {
			post.IsBooked = false
		}
	}

	// * Map table to payload
	// make ถ้าเป็น null จะเป็น array เปล่า
	mappedPosts := make([]*payload.PostResponse, 0)
	for _, post := range posts {

		mappedPosts = append(mappedPosts, &payload.PostResponse{
			PostId:        post.Id,
			OwnerId:       l.Id,
			OwnerUsername: post.Owner.Username,
			Caption:       post.Caption,
			ImageUrl:      post.ImageUrl,
			Application:   post.Application,
			LikeCount:     post.LikeCount,
			CommentCount:  post.CommentCount,
			IsLiked:       &post.IsLiked,
			IsBooked:      &post.IsBooked,
		})
	}

	return c.JSON(response.Info(map[string]any{
		"posts": mappedPosts,
	}))
}
