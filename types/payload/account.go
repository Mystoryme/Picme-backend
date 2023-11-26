package payload

type RegisterRequest struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
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
