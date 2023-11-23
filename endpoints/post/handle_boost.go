package postEndpoint

import (
	"github.com/gofiber/fiber/v2"
	"picme-backend/helper"
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
	transactionId := text.GenerateTransactionId(10)

	if boostCount != 0 {
		return response.Error(false, "The post already boost!")
	}

	boost := &table.PostBoost{
		Id:            nil,
		Post:          nil,
		Paid:          text.Ptr(false),
		Amount:        body.Amount,
		PostId:        body.PostId,
		BoostEnd:      &boostEnd,
		TransactionId: &transactionId,
	}

	if result := mod.DB.Create(boost); result.Error != nil {
		return response.Error(false, "Unable to boost post", result.Error)
	}

	// create qr code
	qrData := helper.ScbCreateQrPayment(uint(*body.Amount), transactionId)

	donateResponse := payload.CreateDonateQrResponse{
		TransactionId: transactionId,
		QrImage:       qrData.QrImage,
		QrRawData:     qrData.QrRawData,
	}

	return c.JSON(response.Info(donateResponse))

}
