package entity

type Customer struct {
	ID                uint   `gorm:"primaryKey"`
	UserID            uint   `gorm:"not null"`
	CityID            uint   `gorm:"not null"`
	Code              string `gorm:"size:50;not null"`
	InvoiceName       string `gorm:"size:255;not null"`
	Address           string `gorm:"size:255;not null"`
	Phone             string `gorm:"size:50"`
	Email             string `gorm:"size:100"`
	Tax               string `gorm:"size:50"`
	Greeting          string `gorm:"size:100"`
	ContactPersonName string `gorm:"size:100"`
	ContactPhone      string `gorm:"size:50"`
	Segment           string `gorm:"size:100"`
	Area              string `gorm:"size:100"`
	Type              string `gorm:"size:50"`
	NPWP              string `gorm:"size:50"`
	Status            string `gorm:"size:50"`
	BEName            string `gorm:"size:100"`
	ProcurementType   string `gorm:"size:100"`
	MarketingName     string `gorm:"size:100"`
	Note              string `gorm:"size:255"`
	PaymentTerm       string `gorm:"size:50"`
	Level             string `gorm:"size:50"`
	User              User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	City              City   `gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;"`
}