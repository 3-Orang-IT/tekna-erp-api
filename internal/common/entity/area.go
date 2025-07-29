package entity

import "time"

type Area struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Employees []Employee `gorm:"many2many:employee_areas;constraint:OnDelete:CASCADE;" json:"employees"`
}
