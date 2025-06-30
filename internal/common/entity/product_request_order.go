package entity

type ProductRequestOrder struct {
	ID             uint `gorm:"primaryKey"`
	ProductID      uint `gorm:"not null"`
	RequestOrderID uint `gorm:"not null"`
	Stock          int  `gorm:"not null"`

	Product      Product      `gorm:"foreignKey:ProductID;constraint:OnDelete:SET NULL;"`
	RequestOrder RequestOrder `gorm:"foreignKey:RequestOrderID;constraint:OnDelete:SET NULL;"`
}
