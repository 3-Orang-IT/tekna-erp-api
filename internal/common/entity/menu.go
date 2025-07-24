package entity

type Menu struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	ParentID *uint  `gorm:"index" json:"parent_id"`
	ModulID  uint   `gorm:"not null" json:"modul_id"`
	Name     string `gorm:"size:100;not null" json:"name"`
	URL      string `gorm:"size:255;not null" json:"url"`
	Icon     string `gorm:"size:255" json:"icon"`
	Order    int    `gorm:"not null" json:"order"`
	Children []Menu `gorm:"foreignKey:ParentID" json:"children"`
}
