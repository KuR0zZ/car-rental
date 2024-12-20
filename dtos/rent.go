package dtos

type RentRequest struct {
	CarID    int `json:"car_id" validate:"required" example:"1"`
	Duration int `json:"duration" validate:"required" example:"2"`
}

type RentResponse struct {
	ID            int     `json:"rental_id" example:"1"`
	UserID        int     `json:"user_id" example:"1"`
	CarName       string  `json:"car_name" example:"Mercedes AMG G63"`
	CarCategory   string  `json:"car_category" example:"SUV"`
	StartDate     string  `json:"start_date" example:"2024-02-01"`
	EndDate       string  `json:"end_date" example:"2024-02-03"`
	InvoiceUrl    string  `json:"invoce_url" example:"example_url"`
	DepositAmount float64 `json:"deposit_amount" example:"10000.00"`
}
