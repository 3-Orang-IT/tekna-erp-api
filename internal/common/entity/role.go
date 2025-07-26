package entity

type Role struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `gorm:"uniqueIndex;size:50" json:"name" binding:"required"`
	Menus []Menu `gorm:"many2many:role_menus;" json:"menus"`
}
