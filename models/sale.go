package models

import "time"

// Sale struct
type Sale struct {
	ID        int `gorm:"primarykey"`
	UserID    int
	ProductID int
	Amount    int
	CreatedAt time.Time
}
