package entity

import "time"

type Document struct {
	ID                 uint      `gorm:"primaryKey"`
	DocumentCategoryID uint      `gorm:"not null"`
	Name               string    `gorm:"size:255;not null"`
	UserID             uint      `gorm:"not null"`
	FilePath           string    `gorm:"size:255"`
	Description        string    `gorm:"size:500"`
	IsPublished        bool      `gorm:"not null"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`

	DocumentCategory DocumentCategory `gorm:"foreignKey:DocumentCategoryID;constraint:OnDelete:SET NULL;"`
	User             User             `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL;"`
}
