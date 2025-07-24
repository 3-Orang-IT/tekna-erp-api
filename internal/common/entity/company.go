package entity

import "time"

type Company struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `gorm:"size:255;not null" json:"name"`
	Address          string    `gorm:"size:255;not null" json:"address"`
	CityID           uint      `gorm:"not null" json:"city_id"`
	Phone            string    `gorm:"size:50" json:"phone"`
	Fax              string    `gorm:"size:50" json:"fax"`
	Email            string    `gorm:"size:100;not null" json:"email"`
	StartHour        string    `gorm:"size:10;not null" json:"start_hour"`
	EndHour          string    `gorm:"size:10;not null" json:"end_hour"`
	Latitude         float64   `gorm:"not null" json:"latitude"`
	Longitude        float64   `gorm:"not null" json:"longitude"`
	TotalShares      int       `gorm:"not null" json:"total_shares"`
	AnnualLeaveQuota int       `gorm:"not null" json:"annual_leave_quota"`
	City             City      `gorm:"foreignKey:CityID" json:"city"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}