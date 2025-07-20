package entity

type Menu struct {
	ID       uint   `gorm:"primaryKey"`
	ParentID *uint  `gorm:"index"`
	ModulID  uint   `gorm:"not null"`
	Name     string `gorm:"size:100;not null"`
	URL      string `gorm:"size:255;not null"`
	Icon     string `gorm:"size:255"`
	Order    int    `gorm:"not null"`
	Children []Menu `gorm:"foreignKey:ParentID"`
}
