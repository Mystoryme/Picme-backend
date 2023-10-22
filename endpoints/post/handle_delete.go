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

func DeleteHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.DeletePostRequest)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	// check ก่อนว่า เป็นเจ้าของ post ไหม
	var post *table.Post
	if result := mod.DB.Where("id = ? ", body.PostId).First(&post); result.Error != nil {
		return response.Error(false, "Unable to call the post", result.Error)
	}
	//post.OwnerId == l.Id เฉยๆไม่ได้ OwnerId: 0xc00041c7d8, User ID: 0xc00041c6b8 เพราะมันคือ address, actual value ต้อง use pointer
	if *post.OwnerId == *l.Id {
		//delete postlike
		var like *table.PostLike
		if result := mod.DB.Where("post_id = ?  ", body.PostId).Delete(&like); result.Error != nil {
			return response.Error(false, "Unable to delete the like", result.Error)
		}

		//delete postcomment
		var comment *table.PostComment
		if result := mod.DB.Where("post_id = ?  ", body.PostId).Delete(&comment); result.Error != nil {
			return response.Error(false, "Unable to delete the comment", result.Error)
		}

		//delete postbookmark
		var bookmark *table.PostBookMark
		if result := mod.DB.Where("post_id = ?  ", body.PostId).Delete(&bookmark); result.Error != nil {
			return response.Error(false, "Unable to delete the bookmark", result.Error)
		}
		//delete postboost
		var boost *table.PostBoost
		if result := mod.DB.Where("post_id = ?  ", body.PostId).Delete(&boost); result.Error != nil {
			return response.Error(false, "Unable to delete the boost post", result.Error)
		}

		//delete postview
		var view *table.PostViews
		if result := mod.DB.Where("post_id = ? ", body.PostId).Delete(&view); result.Error != nil {
			return response.Error(false, "Unable to delete the view", result.Error)
		}

		//delete postdonate
		var donate *table.PostDonate
		if result := mod.DB.Where("post_id = ? ", body.PostId).Delete(&donate); result.Error != nil {
			return response.Error(false, "Unable to delete the ", result.Error)
		}

		//delete post

		if result := mod.DB.Where("id = ?  ", body.PostId).Delete(&post); result.Error != nil {
			return response.Error(false, "Unable to delete the post", result.Error)
		}

	} else {

		return response.Error(false, "Unable to delete the post. You are not the owner of this post.")
	}

	return c.JSON(response.Info("Successfully delete post!"))
}
