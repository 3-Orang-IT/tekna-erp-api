package entity

type BankAccount struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	ChartOfAccountID uint   `gorm:"not null" json:"chart_of_account_id"`
	AccountNumber    string `gorm:"size:50;not null" json:"account_number"`
	BankName         string `gorm:"size:255;not null" json:"bank_name"`
	BranchAddress    string `gorm:"size:255" json:"branch_address"`
	CityID           uint   `gorm:"not null" json:"city_id"`
	PhoneNumber      string `gorm:"size:50" json:"phone_number"`
	Priority         int    `gorm:"not null" json:"priority"`

	ChartOfAccount ChartOfAccount `gorm:"foreignKey:ChartOfAccountID;constraint:OnDelete:SET NULL;" json:"chart_of_account"`
	City           City           `gorm:"foreignKey:CityID;constraint:OnDelete:SET NULL;" json:"city"`
}
