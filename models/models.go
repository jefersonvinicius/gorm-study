package models

import "time"

// User : user model
type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
	Card      Card
}

// Card : card model
type Card struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Number    string
	Password  string
	UserID    uint
}

// Product : product model
type Product struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Price     float64
}
