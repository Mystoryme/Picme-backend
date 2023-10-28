package comment

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

	// * Parse body
	body := new(payload.CommentRequest)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	//// * Query comments
	//var comments []*model.CommentWithProfile
	//if result := mod.DB.Model(new(table.PostComment)).Select("post_comments.* ,(select avatar_url from users where id = post_comments.user_id) as AvatarUrl").Where("post_id = ?", body.PostId).Find(&comments); result.Error != nil {
	//	return response.Error(false, "Unable to query comments", result.Error)
	//}

	comments := make([]*table.PostComment, 0)
	if result := mod.DB.Preload("User").Find(&comments, "post_id = ?", body.PostId); result.Error != nil {
		return response.Error(false, "Unable to query comments", result.Error)
	}
	//การpreload ทำให้ สามารถ ใช้ user.avatarUrlได้ สามารถทำให้ลึกเข้าไปได้ (คล้ายกับหาร join)
	// * Map table to payload
	var mappedComments []*payload.CommentRespond
	for _, comment := range comments {

		mappedComments = append(mappedComments, &payload.CommentRespond{
			UserId:    comment.UserId,
			Username:  comment.User.Username,
			PostId:    comment.PostId,
			AvatarUrl: comment.User.AvatarUrl,
			Message:   comment.Message,
		})
	}
	return c.JSON(response.Info(map[string]any{
		"posts": mappedComments,
	}))

}
