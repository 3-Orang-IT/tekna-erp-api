package entity

type TravelCost struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `gorm:"size:255;not null"`
	Code  string  `gorm:"size:50;not null"`
	Unit  string  `gorm:"size:100;not null"`
	Price float64 `gorm:"not null"`
}
