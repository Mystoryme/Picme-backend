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

func BookmarkHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.CreateBookmarkBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	var bookmarkCount int64
	if result := mod.DB.Model(new(table.PostBookMark)).Where("post_id = ? and user_id = ?", body.PostId, l.Id).Count(&bookmarkCount); result.Error != nil {
		return response.Error(false, "Unable to count bookmarks", result.Error)
	}
	//create
	//ctr+space เติมfillแบบรวดเร็ว
	if bookmarkCount == 0 {
		bookmark := &table.PostBookMark{
			User:      nil,
			UserId:    l.Id,
			Post:      nil,
			PostId:    body.PostId,
			CreatedAt: nil,
		}

		if result := mod.DB.Create(bookmark); result.Error != nil {
			return response.Error(false, "Unable to bookmark post", result.Error)
		}
	}

	return c.JSON(response.Info("Successfully bookmark!"))
}
