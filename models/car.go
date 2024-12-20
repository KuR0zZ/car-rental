package models

type Car struct {
	ID                int     `gorm:"primaryKey;column:car_id" json:"car_id" example:"1"`
	Name              string  `gorm:"column:name" json:"name" example:"Mercedes AMG G63"`
	StockAvailability int     `gorm:"column:stock_availability" json:"stock_availability" example:"3"`
	RentalCosts       float64 `gorm:"column:rental_costs" json:"rental_costs" example:"10000000.00"`
	Category          string  `gorm:"column:category" json:"category" example:"SUV"`
}
