package profileEndpoint

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

func BookmarkPostGetHandler(c *fiber.Ctx) error { // * Parse user claims
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

	// sortby
	db := mod.DB.Preload("Owner").Joins("JOIN post_book_marks ON post_book_marks.post_id = posts.id").
		Where("post_book_marks.user_id = ?", l.Id)

	if query.SortBy == nil || *query.SortBy == enum.SortByDate {
		db = db.Order("created_at DESC")
	} else if *query.SortBy == enum.SortByLike {
		db = db.Order("like_count DESC")
	}

	// * Query posts
	var posts []*model.PostWithCount
	if result := db.Model(new(table.Post)).Select("posts.*, (SELECT COUNT(*) FROM post_likes WHERE post_id = posts.id) AS like_count, (SELECT COUNT(*) FROM post_comments WHERE post_id = posts.id) as comment_count,"+"(SELECT COUNT(*) FROM post_likes WHERE post_id = posts.id AND user_id = ?) AS liked, "+
		"(SELECT COUNT(*) FROM post_book_marks WHERE post_id = posts.id AND user_id = ?) AS booked", l.Id, l.Id).Find(&posts); result.Error != nil {
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
	//make ถ้าเป็น null จะเป็น array เปล่า
	mappedPosts := make([]*payload.BookmarkPostResponse, 0)
	for _, post := range posts {

		mappedPosts = append(mappedPosts, &payload.BookmarkPostResponse{
			PostId:    post.Id,
			ImageUrl:  post.ImageUrl,
			LikeCount: post.LikeCount,
			IsLiked:   &post.IsLiked,
			IsBooked:  &post.IsBooked,
		})
	}

	return c.JSON(response.Info(map[string]any{
		"posts": mappedPosts,
	}))
}
