package models

import "time"

// Sale struct
type Sale struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint
	ProductID uint
	Amount    int64
	CreatedAt time.Time
}
