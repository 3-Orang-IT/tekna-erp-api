package entity

type Employee struct {
    ID                uint    `gorm:"primaryKey"`
    UserID            uint    `gorm:"not null"`
    JobPositionID     uint    `gorm:"not null"`
    DivisionID        uint    `gorm:"not null"`
    CityID            uint    `gorm:"not null"`
    EmployeeStatusID  uint    `gorm:"not null"`
    MarketingAreaID   uint    `gorm:"not null"`
    ContractStatusID  uint    `gorm:"not null"`
    NIP               string  `gorm:"size:50"`
    NIK               string  `gorm:"size:50"`
    BPJSEmploymentNo  string  `gorm:"size:50"`
    BPJSHealthNo      string  `gorm:"size:50"`
    Address           string  `gorm:"size:255"`
    Phone             string  `gorm:"size:50"`
    JoinDate          string  `gorm:"size:50"`
    KTPStatus         string  `gorm:"size:50"`
    ContractNo        string  `gorm:"size:50"`
    NPWPStatus        string  `gorm:"size:50"`

    User              User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
    JobPosition       JobPosition `gorm:"foreignKey:JobPositionID;constraint:OnDelete:SET NULL;"`
    Division          Division        `gorm:"foreignKey:DivisionID;constraint:OnDelete:SET NULL;"`
    City              City    `gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;"`
}
