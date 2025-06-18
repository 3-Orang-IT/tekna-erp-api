package entity

type Role struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"uniqueIndex;size:50"`
    Menus []Menu `gorm:"many2many:role_menus;"`
}
