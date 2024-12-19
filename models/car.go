package models

type Car struct {
	ID                int     `gorm:"primaryKey;column:car_id" swaggerignore:"true"`
	Name              string  `gorm:"column:name" example:"Mercedes AMG G63"`
	StockAvailability int     `gorm:"column:stock_availability" swaggerignore:"true"`
	RentalCosts       float64 `gorm:"column:rental_costs" swaggerignore:"true"`
	Category          string  `gorm:"column:category" example:"SUV"`
}
