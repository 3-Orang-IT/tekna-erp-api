package entity

import "time"

type TravelCost struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Code      string    `gorm:"size:50;not null" json:"code"`
	Unit      string    `gorm:"size:100;not null" json:"unit"`
	Price     float64   `gorm:"not null" json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
