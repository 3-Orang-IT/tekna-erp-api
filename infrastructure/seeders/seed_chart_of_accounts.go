package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedChartOfAccounts(db *gorm.DB) error {
	chartOfAccounts := []entity.ChartOfAccount{
		{Type: "Asset", Code: "COA001", Name: "Cash"},
		{Type: "Liability", Code: "COA002", Name: "Accounts Payable"},
		{Type: "Equity", Code: "COA003", Name: "Retained Earnings"},
		{Type: "Revenue", Code: "COA004", Name: "Sales Revenue"},
		{Type: "Expense", Code: "COA005", Name: "Operating Expenses"},
	}

	for _, account := range chartOfAccounts {
		if err := db.Where("code = ?", account.Code).FirstOrCreate(&account).Error; err != nil {
			log.Printf("Error seeding chart of account %s: %v", account.Code, err)
			return err
		}
	}

	return nil
}
