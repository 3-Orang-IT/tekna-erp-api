package entity

type ProductSalesOrder struct {
	ID           uint `gorm:"primaryKey"`
	ProductID    uint `gorm:"not null"`
	SalesOrderID uint `gorm:"not null"`
	Stock        int  `gorm:"not null"`

	Product    Product    `gorm:"foreignKey:ProductID;constraint:OnDelete:SET NULL;"`
	SalesOrder SalesOrder `gorm:"foreignKey:SalesOrderID;constraint:OnDelete:SET NULL;"`
}
