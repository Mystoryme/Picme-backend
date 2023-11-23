package postEndpoint

import (
	"github.com/gofiber/fiber/v2"
	"picme-backend/helper"
	mod "picme-backend/modules"
	"picme-backend/types/payload"
	"picme-backend/types/response"
	"picme-backend/types/table"
	"picme-backend/utils/text"
)

func BoostInquiry(c *fiber.Ctx) error {
	// * Parse body
	body := new(payload.PaymentInquiryRequest)
	if err := c.BodyParser(body); err != nil {
		return response.Error(false, "Unable to parse body", err)
	}

	// * Validate body
	if err := text.Validator.Struct(body); err != nil {
		return err
	}

	inquiryData, err := helper.ScbInquiryPayment(body.TransactionId)
	if err != nil {
		return response.Error(false, "Unable to inquiry payment", err)
	}

	paymentStatusResponse := new(payload.PaymentInquiryResponse)

	if inquiryData != nil && *inquiryData.PayeeName != "" && inquiryData.PayeeName != nil {
		paymentStatusResponse.PaymentSuccess = true

		transactionData := new(table.PostBoost)
		if result := mod.DB.First(transactionData, "transaction_id = ?", body.TransactionId); result.Error != nil {
			paymentStatusResponse.Message = "Unable to query transaction"
			return response.Error(false, "Unable to query transaction", result.Error)
		}

		if err := mod.DB.Model(transactionData).Where("transaction_id = ?", body.TransactionId).Update("paid", true).Error; err != nil {
			return response.Error(false, "Unable to update transaction state", err.Error)
		}

	} else {
		paymentStatusResponse.PaymentSuccess = false
		paymentStatusResponse.Message = "Payment not found"
	}

	return c.JSON(response.Info(paymentStatusResponse))
}
