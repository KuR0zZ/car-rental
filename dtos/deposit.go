package dtos

type DepositRequest struct {
	DepositAmount float64 `json:"deposit_amount" validate:"required"`
}

type DepositResponse struct {
	Message       string  `json:"message"`
	DepositAmount float64 `json:"deposit_amount"`
}
