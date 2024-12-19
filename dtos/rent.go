package dtos

type RentRequest struct {
	CarID    int `json:"car_id" validate:"required"`
	Duration int `json:"duration" validate:"required"`
}

type RentResponse struct {
	ID            int         `json:"rental_id"`
	UserID        int         `json:"user_id"`
	CarRent       interface{} `json:"car_rent"`
	InvoiceUrl    string      `json:"invoce_url"`
	DepositAmount float64     `json:"deposit_amount"`
}
