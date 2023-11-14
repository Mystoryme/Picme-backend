package postEndpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	mod "picme-backend/modules"
	"picme-backend/types/misc"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"
)

// โชว์post ที่หลักจากกดที่ boost post // grid // bookmark

func GetPostHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.GetPostRequest)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	// * Query posts
	var posts []*table.Post
	if result := mod.DB.Preload("Owner").Where("id = ?", body.PostId).Find(&posts); result.Error != nil {
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

		var liked int64
		if result := mod.DB.Model(&table.PostLike{}).Where("post_id = ? AND user_id =?", post.Id, l.Id).Count(&liked); result.Error != nil {
			return response.Error(false, "Unable to count likes", result.Error)

		}
		var booked int64
		if result := mod.DB.Model(new(table.PostBookMark)).Where("post_id = ? AND user_id =?", post.Id, l.Id).Count(&booked); result.Error != nil {
			return response.Error(false, "Unable to count books", result.Error)
		}

		like := false
		if liked == 1 {
			like = true
		}
		book := false
		if booked == 1 {
			book = true
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
			IsLiked:       &like,
			IsBooked:      &book,
		})
	}

	return c.JSON(response.Info(map[string]any{
		"posts": mappedPosts,
	}))

}
