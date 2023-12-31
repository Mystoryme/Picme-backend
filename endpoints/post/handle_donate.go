package postEndpoint

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"picme-backend/helper"
	mod "picme-backend/modules"
	"picme-backend/types/misc"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"
)

func DonateHandler(c *fiber.Ctx) error {
	// * Parse user claims
	l := c.Locals("l").(*jwt.Token).Claims.(*misc.UserClaim)

	// * Parse body
	body := new(payload.CreateDonateBody)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	//create
	//ctr+space เติมfillแบบรวดเร็ว

	currentPost := new(table.Post)

	if result := mod.DB.Preload("Owner").First(currentPost, "id = ?", body.PostId); result.Error != nil {
		return response.Error(false, "Unable to query post", result.Error)
	}

	transactionId := text.GenerateTransactionId(10)

	donate := &table.PostDonate{
		User:          nil,
		UserId:        l.Id,
		Paid:          text.Ptr(false),
		Post:          nil,
		PostId:        body.PostId,
		DonateAmount:  body.DonateAmount,
		TransactionId: &transactionId,
		CreatedAt:     nil,
	}

	if result := mod.DB.Create(donate); result.Error != nil {
		return response.Error(false, "Unable to donate post", result.Error)
	}

	// create qr code
	qrData := helper.ScbCreateQrPayment(uint(*body.DonateAmount), transactionId)

	donateResponse := payload.CreateDonateQrResponse{
		TransactionId: transactionId,
		QrImage:       qrData.QrImage,
		QrRawData:     qrData.QrRawData,
	}

	//postDonateType := enum.NotificationPostDonate
	//notification := &table.Notification{
	//	Trigger:          nil,
	//	TriggerId:        l.Id,
	//	Triggee:          nil,
	//	TriggeeId:        currentPost.OwnerId,
	//	Post:             nil,
	//	PostId:           body.PostId,
	//	NotificationType: &postDonateType,
	//	CreatedAt:        nil,
	//}
	//
	//if result := mod.DB.Create(&notification); result.Error != nil {
	//	return response.Error(false, "Unable to create notification", result.Error)
	//}

	return c.JSON(response.Info(donateResponse))
}
