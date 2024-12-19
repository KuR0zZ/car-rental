package dtos

type RegisterRequest struct {
	Name     string `json:"name" validate:"required" example:"John Doe"`
	Email    string `json:"email" validate:"required,email" example:"John.Doe@example.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

type RegisterResponse struct {
	Message string      `json:"message" example:"Successfully Register New User"`
	Data    interface{} `json:"data"`
}
