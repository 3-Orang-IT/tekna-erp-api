package entity

type RequestOrder struct {
	ID                 uint   `gorm:"primaryKey"`
	RequestOrderNumber string `gorm:"size:50;not null"`
	SupplierID         uint   `gorm:"not null"`
	RequesterID        uint   `gorm:"not null"`
	Term               string `gorm:"size:100"`
	RequestDate        string `gorm:"size:50;not null"`
	BusinessUnitID     uint   `gorm:"not null"`
	DeliveryAddress    string `gorm:"size:255"`
	Note               string `gorm:"size:255"`

	Supplier     Supplier     `gorm:"foreignKey:SupplierID;constraint:OnDelete:SET NULL;"`
	Requester    User         `gorm:"foreignKey:RequesterID;constraint:OnDelete:SET NULL;"`
	BusinessUnit BusinessUnit `gorm:"foreignKey:BusinessUnitID;constraint:OnDelete:SET NULL;"`
}
