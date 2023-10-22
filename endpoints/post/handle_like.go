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

func LikeHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.CreateLikeBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}
	//มีไ where user_id ด้วย เพราะจะนับ likecount ของคนนั้น ถ้าเป็ฯ 0 จะlike ได้ แต่ถ้าเป็น 1 แล้วจะlike ไม่ได้ละ
	var likeCount int64
	if result := mod.DB.Model(new(table.PostLike)).Where("post_id = ? and user_id = ?", body.PostId, l.Id).Count(&likeCount); result.Error != nil {
		return response.Error(false, "Unable to count likes", result.Error)
	}
	//create
	//ctr+space เติมfillแบบรวดเร็ว
	if likeCount == 0 {
		like := &table.PostLike{
			User:      nil,
			UserId:    l.Id,
			Post:      nil,
			PostId:    body.PostId,
			CreatedAt: nil,
		}

		if result := mod.DB.Create(like); result.Error != nil {
			return response.Error(false, "Unable to like post", result.Error)
		}
	}

	return c.JSON(response.Info("Successfully like!"))
}
