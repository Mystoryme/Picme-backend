package payload

// ScbGetTokenRequest is the request body for getting access token

type ScbGetTokenRequest struct {
	ApplicationKey    string `json:"applicationKey"`
	ApplicationSecret string `json:"applicationSecret"`
}

// ScbStatusResponse is the response body for status

type ScbStatusResponse struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

// ScbGetTokenResponse is the response body for getting access token

type ScbGetTokenResponse struct {
	Status ScbStatusResponse       `json:"status,omitempty"`
	Data   ScbGetTokenDataResponse `json:"data,omitempty"`
}

type ScbGetTokenDataResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
	ExpiresIn   int    `json:"expiresIn,omitempty"`
	TokenType   string `json:"tokenType,omitempty"`
	ExpiresAt   int    `json:"expiresAt,omitempty"`
}

// ScbCreateQrPaymentRequest is the request body for creating QR payment

type ScbCreateQrPaymentRequest struct {
	QrType string `json:"qrType"`
	Amount string `json:"amount"`
	PpType string `json:"ppType"`
	PpId   string `json:"ppId"`
	Ref1   string `json:"ref1"`
	Ref2   string `json:"ref2"`
	Ref3   string `json:"ref3"`
}

type ScbCreateQrResponse struct {
	Status ScbStatusResponse       `json:"status,omitempty"`
	Data   ScbCreateQrDataResponse `json:"data,omitempty"`
}

type ScbCreateQrDataResponse struct {
	QrRawData string `json:"qrRawData,omitempty"`
	QrImage   string `json:"qrImage,omitempty"`
}
