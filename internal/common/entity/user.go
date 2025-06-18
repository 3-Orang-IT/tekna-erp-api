package entity

import "time"

type User struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Username         string    `gorm:"size:100;uniqueIndex" json:"username"`
	Password         string    `json:"password"`
	Name             string    `gorm:"size:100" json:"name"`
	Email            string    `gorm:"size:100;uniqueIndex" json:"email"`
	Telp             string    `gorm:"size:20" json:"telp"`
	PhotoProfileURL  string    `gorm:"size:255" json:"photo_profile_url"`
	Status           string    `gorm:"size:50" json:"status"` // Atur sesuai kebutuhan: aktif/nonaktif/dll

	Role             []Role    `gorm:"many2many:user_roles;" json:"roles"`

	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
