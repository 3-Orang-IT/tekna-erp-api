package entity

type Newsletter struct {
	ID          uint   `gorm:"primaryKey"`
	Type        string `gorm:"size:100;not null"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"size:500"`
	File        string `gorm:"size:255"`
	ValidFrom   string `gorm:"size:50;not null"`
	Status      string `gorm:"size:50;not null"`
}
