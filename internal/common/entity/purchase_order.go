package entity

import (
	"time"
)

// PurchaseOrder represents a purchase order in the system.
type PurchaseOrder struct {
	ID        uint      `gorm:"primaryKey"`
	SupplierID uint      `gorm:"not null" json:"supplier_id"`
	BusinessUnitID uint      `gorm:"not null" json:"business_unit_id"`
	OrderNumber string    `gorm:"size:50;not null" json:"order_number"`
	Description string    `gorm:"size:500" json:"description"`
	TotalAmount float64 `gorm:"not null" json:"total_amount"`
	Status    string    `gorm:"size:50;not null" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	Supplier  Supplier   `gorm:"foreignKey:SupplierID;constraint:OnDelete:SET NULL;"`
	Product    []Product   `gorm:"many2many:purchase_order_products;"`
}
