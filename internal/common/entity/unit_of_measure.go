package entity

type UnitOfMeasure struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"size:100;not null" json:"name"`
	Abbreviation string `gorm:"size:50;not null" json:"abbreviation"`
}
