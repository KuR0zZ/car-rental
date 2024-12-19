package dtos

type TopUpRequest struct {
	DepositAmount float64 `json:"deposit_amount" validate:"required" example:"100000.00"`
}

type TopUpResponse struct {
	Message       string  `json:"message" example:"Successfully Top Up Balance"`
	UserID        int     `json:"user_id" example:"1"`
	DepositAmount float64 `json:"deposit_amount" example:"100000.00"`
}
