package entity

import "time"

type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"size:100" json:"name"`
    Email     string    `gorm:"uniqueIndex;size:100" json:"email"`
    Password  string    `json:"password"`
    RoleID    uint      `json:"role_id"`
    Role      Role      `json:"role"` 
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}