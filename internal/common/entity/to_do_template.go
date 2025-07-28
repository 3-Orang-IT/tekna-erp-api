package entity

type ToDoTemplate struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	JobPositionID uint   `gorm:"not null" json:"job_position_id"`
	Activity      string `gorm:"size:255;not null" json:"activity"`
	Priority      int    `gorm:"not null" json:"priority"`
	OrderNumber   int    `gorm:"not null" json:"order_number"`

	JobPosition JobPosition `gorm:"foreignKey:JobPositionID;constraint:OnDelete:CASCADE;" json:"job_position"`
}
