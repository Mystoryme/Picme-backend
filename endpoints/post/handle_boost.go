package postEndpoint

import (
	"github.com/gofiber/fiber/v2"
	mod "picme-backend/modules"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"
	"time"
	_ "time"
)

func BoostHandler(c *fiber.Ctx) error {
	// * Parse user claims
	//l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.BoostBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	var boostCount int64
	if result := mod.DB.Model(new(table.PostBoost)).Where("post_id = ? AND boost_end > ?", body.PostId, time.Now()).Count(&boostCount); result.Error != nil {
		return response.Error(false, "Unable to count boost", result.Error)
	}
	boostEnd := time.Now().AddDate(0, 0, *body.BoostDay)
	//create
	//ctr+space เติมfillแบบรวดเร็ว
	if boostCount == 0 {
		boost := &table.PostBoost{
			Id:       nil,
			Post:     nil,
			PostId:   body.PostId,
			BoostEnd: &boostEnd,
		}

		if result := mod.DB.Create(boost); result.Error != nil {
			return response.Error(false, "Unable to boost post", result.Error)
		}
	}

	return c.JSON(response.Info("Successfully boost!"))
}
