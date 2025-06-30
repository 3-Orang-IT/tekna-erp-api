package entity

type BankAccount struct {
	ID               uint   `gorm:"primaryKey"`
	ChartOfAccountID uint   `gorm:"not null"`
	AccountNumber    string `gorm:"size:50;not null"`
	BankName         string `gorm:"size:255;not null"`
	BranchAddress    string `gorm:"size:255"`
	CityID           uint   `gorm:"not null"`
	PhoneNumber      string `gorm:"size:50"`
	Priority         int    `gorm:"not null"`

	ChartOfAccount ChartOfAccount `gorm:"foreignKey:ChartOfAccountID;constraint:OnDelete:SET NULL;"`
	City           City           `gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;"`
}
