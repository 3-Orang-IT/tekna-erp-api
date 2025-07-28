package entity

import "time"

type Customer struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	UserID            uint      `gorm:"not null" json:"user_id"`
	AreaID            uint      `gorm:"not null" json:"area_id"`
	Name              string    `gorm:"size:100;not null" json:"name"`
	CityID            uint      `gorm:"not null" json:"city_id"`
	Code              string    `gorm:"size:50;not null" json:"code"`
	InvoiceName       string    `gorm:"size:255;not null" json:"invoice_name"`
	Address           string    `gorm:"size:255;not null" json:"address"`
	Phone             string    `gorm:"size:50" json:"phone"`
	Email             string    `gorm:"size:100" json:"email"`
	Tax               string    `gorm:"size:50" json:"tax"`
	Greeting          string    `gorm:"size:100" json:"greeting"`
	ContactPersonName string    `gorm:"size:100" json:"contact_person_name"`
	ContactPhone      string    `gorm:"size:50" json:"contact_phone"`
	Segment           string    `gorm:"size:100" json:"segment"`
	Type              string    `gorm:"size:50" json:"type"`
	NPWP              string    `gorm:"size:50" json:"npwp"`
	Status            string    `gorm:"size:50;default:'active'" json:"status"`
	BEName            string    `gorm:"size:100" json:"be_name"`
	ProcurementType   string    `gorm:"size:100" json:"procurement_type"`
	MarketingName     string    `gorm:"size:100" json:"marketing_name"`
	Note              string    `gorm:"size:255" json:"note"`
	PaymentTerm       string    `gorm:"size:50" json:"payment_term"`
	Level             string    `gorm:"size:50" json:"level"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
	City City `gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;" json:"city"`
	Area Area `gorm:"foreignKey:AreaID;constraint:OnDelete:SET NULL;" json:"area"`
}