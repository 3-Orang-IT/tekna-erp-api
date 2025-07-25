package entity

import (
	"time"
)

type Supplier struct {
	ID            uint   `gorm:"primaryKey"`
	UserID        uint   `gorm:"not null"`
	Code          uint   `gorm:"not null;autoIncrement"`
	Name          string `gorm:"size:255;not null"`
	InvoiceName   string `gorm:"size:255;not null"`
	NPWP          string `gorm:"size:50;not null"`
	Address       string `gorm:"size:255;not null"`
	CityID        uint   `gorm:"not null"`
	Phone         string `gorm:"size:50"`
	Email         string `gorm:"size:100"`
	Greeting      string `gorm:"size:100"`
	ContactPerson string `gorm:"size:100"`
	ContactPhone  string `gorm:"size:50"`
	BankAccount   string `gorm:"size:100"`
	Type          string `gorm:"size:50"`
	LogoFilename  string `gorm:"size:255"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	City City `gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;"`
}
