package models

type User struct {
	ID            int     `gorm:"primaryKey;column:user_id" example:"1"`
	Name          string  `gorm:"name" example:"John Doe"`
	Email         string  `gorm:"column:email" example:"John.Doe@example.com"`
	Password      string  `gorm:"column:password" swaggerignore:"true"`
	DepositAmount float64 `gorm:"column:deposit_amount" example:"2000000"`
}
