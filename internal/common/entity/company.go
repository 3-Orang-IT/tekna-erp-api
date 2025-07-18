package entity

import "time"

type Company struct {
	ID               string    `gorm:"primaryKey;size:36"`
	Name             string    `gorm:"size:255;not null"`
	Address          string    `gorm:"size:500;not null"`
	CityID           string    `gorm:"size:50;not null"`
	ProvinceID       string    `gorm:"size:50;not null"`
	Phone            string    `gorm:"size:50"`
	Fax              string    `gorm:"size:50"`
	Email            string    `gorm:"size:100;not null"`
	StartHour        string    `gorm:"size:10;not null"`
	EndHour          string    `gorm:"size:10;not null"`
	Latitude         float64   `gorm:"not null"`
	Longitude        float64   `gorm:"not null"`
	TotalShares      int       `gorm:"not null"`
	AnnualLeaveQuota int       `gorm:"not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
	DeletedAt        *time.Time `gorm:"index"`
}