package model

import "time"

type User struct {
	ID uint `gorm:"primarykey"`
	Email string `gorm:"UniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Role      string    `gorm:"default:customer"`
	CreatedAt time.Time
	UpdatedAt time.Time
}