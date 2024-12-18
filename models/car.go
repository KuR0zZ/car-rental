package models

type Car struct {
	ID                int     `gorm:"primaryKey;column:car_id"`
	Name              string  `gorm:"column:name"`
	StockAvailability int     `gorm:"column:stock_availability"`
	RentalCosts       float64 `gorm:"column:rental_costs"`
	Category          string  `gorm:"column:category"`
}
