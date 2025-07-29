package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedBudgetCategories(db *gorm.DB) error {
	var count int64
	db.Model(&entity.BudgetCategory{}).Count(&count)
	if count > 0 {
		log.Println("Budget categories already seeded")
		return nil
	}

	// Get chart of account IDs for reference
	var chartOfAccounts []entity.ChartOfAccount
	if err := db.Find(&chartOfAccounts).Error; err != nil {
		return err
	}

	var coaID *uint
	if len(chartOfAccounts) > 0 {
		coaID = &chartOfAccounts[0].ID
	}

	categories := []entity.BudgetCategory{
		{
			ChartOfAccountID: coaID,
			Name: "Operational Expense",
			Description: "Expenses for daily operations",
			Order: 1,
		},
		{
			ChartOfAccountID: coaID,
			Name: "Marketing Expense",
			Description: "Expenses for marketing activities",
			Order: 2,
		},
		{
			ChartOfAccountID: coaID,
			Name: "Capital Expenditure",
			Description: "Long-term asset purchases",
			Order: 3,
		},
	}

	for _, category := range categories {
		if err := db.Create(&category).Error; err != nil {
			return err
		}
	}

	log.Println("Budget categories seeded successfully")
	return nil
}
