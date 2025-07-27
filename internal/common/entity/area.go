package entity

type Area struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null" json:"name"`
}
