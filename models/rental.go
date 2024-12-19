package models

import "time"

type Invoice struct {
	ID         string `json:"id"`
	InvoiceUrl string `json:"invoice_url"`
}

type Rental struct {
	ID         int       `gorm:"primaryKey;column:rental_id"`
	UserID     int       `gorm:"column:user_id"`
	CarID      int       `gorm:"column:car_id"`
	Duration   int       `gorm:"column:duration"`
	StartDate  time.Time `gorm:"column:start_date"`
	EndDate    time.Time `gorm:"column:end_date"`
	TotalCosts float64   `gorm:"column:total_costs"`
	Status     string    `gorm:"column:status"`
	Car        Car       `gorm:"foreignKey:CarID;references:ID"`
}
