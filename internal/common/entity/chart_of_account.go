package entity

type ChartOfAccount struct {
	ID   uint   `gorm:"primaryKey"`
	Type string `gorm:"size:100;not null"`
	Code string `gorm:"size:50;not null"`
	Name string `gorm:"size:255;not null"`
}
