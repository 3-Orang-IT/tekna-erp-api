package entity

type Province struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"size:100;not null" json:"name"`
	Cities []City `gorm:"foreignKey:ProvinceID;constraint:OnDelete:CASCADE;" json:"cities,omitempty"`
}
