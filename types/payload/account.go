package payload

type LoginResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
}
