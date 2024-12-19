package dtos

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"John.Doe@example.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

type LoginResponse struct {
	Message string `json:"message" example:"Successfully logged in"`
	Token   string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"`
}
