package entity

type Purchase struct {
	ID             uint   `gorm:"primaryKey"`
	SupplierID     uint   `gorm:"not null"`
	CustomerID     uint   `gorm:"not null"`
	BusinessUnitID uint   `gorm:"not null"`
	POCode         string `gorm:"size:50;not null"`
	SONumber       string `gorm:"size:50"`
	RONumber       string `gorm:"size:50"`
	Term           string `gorm:"size:100"`
	PODate         string `gorm:"size:50;not null"`
	AreaCode       string `gorm:"size:50"`
	Note           string `gorm:"size:255"`

	Supplier     Supplier     `gorm:"foreignKey:SupplierID;constraint:OnDelete:SET NULL;"`
	Customer     Customer     `gorm:"foreignKey:CustomerID;constraint:OnDelete:SET NULL;"`
	BusinessUnit BusinessUnit `gorm:"foreignKey:BusinessUnitID;constraint:OnDelete:SET NULL;"`
}
