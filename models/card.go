package models

import "time"

// Card : card model
type Card struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Number    string
	Password  string
	UserID    uint
}
