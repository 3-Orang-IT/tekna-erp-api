package entity

import "time"

type Role struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"uniqueIndex;size:50" json:"name" binding:"required"`
	Code      string    `gorm:"uniqueIndex;size:20" json:"code" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Menus []Menu `gorm:"many2many:role_menus;constraint:OnDelete:CASCADE" json:"menus"`
}
