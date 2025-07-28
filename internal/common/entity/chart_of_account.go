package entity

type ChartOfAccount struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Type string `gorm:"size:100;not null" json:"type"`
	Code string `gorm:"size:50;not null" json:"code"`
	Name string `gorm:"size:255;not null" json:"name"`
}
