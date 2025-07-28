package entity

import "time"

type Province struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Cities []City `gorm:"foreignKey:ProvinceID;constraint:OnDelete:CASCADE;" json:"cities,omitempty"`
}
