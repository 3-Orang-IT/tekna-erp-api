package entity

type SalesOrder struct {
	ID            uint   `gorm:"primaryKey"`
	SONumber2     string `gorm:"size:50;not null"`
	CustomerPONo  string `gorm:"size:50"`
	CustomerID    uint   `gorm:"not null"`
	SalesmanID    uint   `gorm:"not null"`
	Term          string `gorm:"size:100"`
	OrderDate     string `gorm:"size:50;not null"`
	DeliveryDate  string `gorm:"size:50"`
	DueDate       string `gorm:"size:50"`
	AreaCode      string `gorm:"size:50"`
	ProjectStatus string `gorm:"size:100"`
	Note          string `gorm:"size:255"`

	Customer Customer `gorm:"foreignKey:CustomerID;constraint:OnDelete:SET NULL;"`
	Salesman User     `gorm:"foreignKey:SalesmanID;constraint:OnDelete:SET NULL;"`
}
