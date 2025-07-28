package entity

import "time"

type City struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"size:100;not null" json:"name"`
	ProvinceID *uint     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"province_id"`
	Province   Province  `gorm:"foreignKey:ProvinceID" json:"province"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}