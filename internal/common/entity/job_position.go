package entity

import "time"

type JobPosition struct {
   ID        uint      `gorm:"primaryKey"`
   Name      string    `gorm:"size:100;not null"`
   CreatedAt time.Time `json:"created_at"`
   UpdatedAt time.Time `json:"updated_at"`
}
