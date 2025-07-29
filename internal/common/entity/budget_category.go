package entity

import "time"

type BudgetCategory struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	ChartOfAccountID *uint          `gorm:"column:chart_of_account_id" json:"chart_of_account_id,omitempty"`
	Name             string         `gorm:"size:100;not null" json:"name"`
	Description      string         `gorm:"size:255" json:"description,omitempty"`
	Order            int            `gorm:"not null" json:"order"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`

	ChartOfAccount   *ChartOfAccount `gorm:"foreignKey:ChartOfAccountID;constraint:OnDelete:SET NULL;" json:"chart_of_account,omitempty"`
}
