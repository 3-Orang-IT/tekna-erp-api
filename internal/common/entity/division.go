package entity

type Division struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null"`
}
