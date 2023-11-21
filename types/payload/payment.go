package payload

type ScbGetTokenRequest struct {
	ApplicationKey    string `json:"applicationKey"`
	ApplicationSecret string `json:"applicationSecret"`
}

type ScbGetTokenResponse struct {
	Status ScbGetTokenStatusResponse `json:"status,omitempty"`
	Data   ScbGetTokenDataResponse   `json:"data,omitempty"`
}

type ScbGetTokenStatusResponse struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type ScbGetTokenDataResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
	ExpiresIn   int    `json:"expiresIn,omitempty"`
	TokenType   string `json:"tokenType,omitempty"`
	ExpiresAt   int    `json:"expiresAt,omitempty"`
}
