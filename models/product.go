package models

import "time"

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       int64
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
