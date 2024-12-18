package models

type User struct {
	ID            int     `gorm:"primaryKey;column:user_id"`
	Email         string  `gorm:"column:email"`
	Password      string  `gorm:"column:password"`
	DepositAmount float64 `gorm:"column:deposit_amount"`
}
