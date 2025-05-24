package models

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"type:varchar(100);not null;unique"`
	Email string `json:"email" gorm:"type:varchar(100);not null;unique"`
}