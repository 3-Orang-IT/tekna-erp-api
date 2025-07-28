package entity

import "time"

type Employee struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	UserID           uint      `gorm:"not null" json:"user_id"`
	JobPositionID    uint      `gorm:"not null" json:"job_position_id"`
	DivisionID       uint      `gorm:"not null" json:"division_id"`
	CityID           uint      `gorm:"not null" json:"city_id"`
	NIP              string    `gorm:"column:nip;size:50" json:"nip"`
	NIK              string    `gorm:"size:50" json:"nik"`
	BPJSEmploymentNo string    `gorm:"size:50" json:"bpjs_employment_no"`
	BPJSHealthNo     string    `gorm:"size:50" json:"bpjs_health_no"`
	Address          string    `gorm:"size:255" json:"address"`
	Phone            string    `gorm:"size:50" json:"phone"`
	JoinDate         string    `gorm:"size:50" json:"join_date"`
	KTPStatus        string    `gorm:"size:50" json:"ktp_status"`
	ContractNo       string    `gorm:"size:50" json:"contract_no"`
	NPWPStatus       string    `gorm:"size:50" json:"npwp_status"`
	ContractStatus   string    `gorm:"size:50;default:'active'" json:"contract_status"`
	Status           string    `gorm:"size:50;default:'active'" json:"status"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	User        User        `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"user"`
	JobPosition JobPosition `gorm:"foreignKey:JobPositionID;constraint:OnDelete:SET NULL;" json:"job_position"`
	Division    Division    `gorm:"foreignKey:DivisionID;constraint:OnDelete:SET NULL;" json:"division"`
	City        City        `gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;" json:"city"`
	Area        []Area      `gorm:"many2many:employee_areas;constraint:OnDelete:CASCADE;" json:"area"`
}
