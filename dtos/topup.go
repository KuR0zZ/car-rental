package dtos

type TopUpRequest struct {
	DepositAmount float64 `json:"deposit_amount" validate:"required"`
}

type TopUpResponse struct {
	Message       string  `json:"message"`
	UserID        int     `json:"user_id"`
	DepositAmount float64 `json:"deposit_amount"`
}
