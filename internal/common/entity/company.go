package entity

type Company struct {
	ID               uint    `gorm:"primaryKey"`
	Name             string  `gorm:"size:255;not null"`
	Address          string  `gorm:"size:255;not null"`
	CityID           uint    `gorm:"not null"`
	ProvinceID       uint    `gorm:"not null"`
	Phone            string  `gorm:"size:50"`
	Fax              string  `gorm:"size:50"`
	Email            string  `gorm:"size:100;not null"`
	StartHour        string  `gorm:"size:10;not null"`
	EndHour          string  `gorm:"size:10;not null"`
	Latitude         float64 `gorm:"not null"`
	Longitude        float64 `gorm:"not null"`
	TotalShares      int     `gorm:"not null"`
	AnnualLeaveQuota int     `gorm:"not null"`
}