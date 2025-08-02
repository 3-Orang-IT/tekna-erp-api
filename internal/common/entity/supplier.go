package entity

import (
	"time"
)

type Supplier struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	UserID        uint   `gorm:"default:null" json:"user_id"`
	Code          uint   `gorm:"not null;autoIncrement" json:"code"`
	Name          string `gorm:"size:255;not null" json:"name"`
	InvoiceName   string `gorm:"size:255;not null" json:"invoice_name"`
	NPWP          string `gorm:"size:50;not null" json:"npwp"`
	Address       string `gorm:"size:255;not null" json:"address"`
	CityID        uint   `gorm:"not null" json:"city_id"`
	Phone         string `gorm:"size:50" json:"phone"`
	Email         string `gorm:"size:100" json:"email"`
	Greeting      string `gorm:"size:100" json:"greeting"`
	ContactPerson string `gorm:"size:100" json:"contact_person"`
	ContactPhone  string `gorm:"size:50" json:"contact_phone"`
	BankAccount   string `gorm:"size:100" json:"bank_account"`
	Type          string `gorm:"size:50" json:"type"`
	LogoFilename  string `gorm:"size:255" json:"logo_filename"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	City City `gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;"`
}
