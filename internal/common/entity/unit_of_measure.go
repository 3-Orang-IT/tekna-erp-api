package entity

type UnitOfMeasure struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:100;not null"`
	Abbreviation string `gorm:"size:50;not null"`
}
