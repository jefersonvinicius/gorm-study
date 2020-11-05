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
	Products  []*Product `gorm:"many2many:sales;"`
}
