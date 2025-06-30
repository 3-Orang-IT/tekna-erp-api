package entity

type ToDoTemplate struct {
	ID            uint   `gorm:"primaryKey"`
	JobPositionID uint   `gorm:"not null"`
	Activity      string `gorm:"size:255;not null"`
	Priority      int    `gorm:"not null"`
	OrderNumber   int    `gorm:"not null"`

	JobPosition JobPosition `gorm:"foreignKey:JobPositionID;constraint:OnDelete:CASCADE;"`
}
