package entity

import "time"

type ProductSalesOrder struct {
	ID           uint      `gorm:"primaryKey"`
	ProductID    uint      `gorm:"not null"`
	SalesOrderID uint      `gorm:"not null"`
	Stock        int       `gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Product    Product    `gorm:"foreignKey:ProductID;constraint:OnDelete:SET NULL;"`
	SalesOrder SalesOrder `gorm:"foreignKey:SalesOrderID;constraint:OnDelete:SET NULL;"`
}
