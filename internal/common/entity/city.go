package entity

type City struct {
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"size:100;not null"`
	ProvinceID uint   `gorm:"not null"`
}