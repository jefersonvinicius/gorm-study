package models

import "time"

// Product : product model
type Product struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Price     float64
}
