package dtos

type RentalReportResponse struct {
	ID          int     `json:"rental_id" example:"1"`
	UserID      int     `json:"user_id" example:"1"`
	CarName     string  `json:"car_name" example:"Mercedes AMG G63"`
	CarCategory string  `json:"car_category" example:"SUV"`
	Duration    int     `json:"duration" example:"2"`
	StartDate   string  `json:"start_date" example:"2024-02-01"`
	EndDate     string  `json:"end_date" example:"2024-02-03"`
	TotalCosts  float64 `json:"total_costs" example:"10000.00"`
	Status      string  `json:"status" example:"Active"`
}
