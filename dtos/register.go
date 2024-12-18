package dtos

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
