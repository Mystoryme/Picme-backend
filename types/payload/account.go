package payload

type RegisterRequest struct {
	Username        *string `json:"username" validate:"required"`
	Password        *string `json:"password" validate:"required"`
	Email           *string `json:"email" validate:"required"`
	ConfirmPassword *string `json:"confirmPassword" validate:"required"`
}

type RegisterResponse struct {
	UserId *uint64 `json:"userId"`
}

type LoginRequest struct {
	Email    *string `json:"email" validate:"required"`
	Password *string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token *string `json:"token"`
}
