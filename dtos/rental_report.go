package dtos

type RentalReportResponse struct {
	ID          int     `json:"rental_id"`
	UserID      int     `json:"user_id"`
	CarName     string  `json:"car_name"`
	CarCategory string  `json:"car_category"`
	Duration    int     `json:"duration"`
	TotalCosts  float64 `json:"total_costs"`
	Status      string  `json:"status"`
}
