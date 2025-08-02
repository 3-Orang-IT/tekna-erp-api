package seeders

import (
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedBankAccounts(tx *gorm.DB) error {
	var count int64
	tx.Model(&entity.BankAccount{}).Count(&count)
	if count > 0 {
		return nil // Bank accounts already seeded
	}

	// Get a few chart of accounts to assign to bank accounts
	var chartOfAccounts []entity.ChartOfAccount
	if err := tx.Limit(3).Find(&chartOfAccounts).Error; err != nil {
		return err
	}

	// Get a few cities to assign to bank accounts
	var cities []entity.City
	if err := tx.Limit(3).Find(&cities).Error; err != nil {
		return err
	}

	// If we have chart of accounts and cities, create some bank accounts
	if len(chartOfAccounts) > 0 && len(cities) > 0 {
		bankAccounts := []entity.BankAccount{
			{
				ChartOfAccountID: chartOfAccounts[0].ID,
				AccountNumber:    "1234567890",
				BankName:         "Bank Central Asia",
				BranchAddress:    "Jl. Sudirman No. 10, Jakarta",
				CityID:           cities[0].ID,
				PhoneNumber:      "021-5555123",
				Priority:         1,
			},
			{
				ChartOfAccountID: chartOfAccounts[1 % len(chartOfAccounts)].ID,
				AccountNumber:    "0987654321",
				BankName:         "Bank Mandiri",
				BranchAddress:    "Jl. Thamrin No. 5, Jakarta",
				CityID:           cities[1 % len(cities)].ID,
				PhoneNumber:      "021-5555456",
				Priority:         2,
			},
			{
				ChartOfAccountID: chartOfAccounts[2 % len(chartOfAccounts)].ID,
				AccountNumber:    "1122334455",
				BankName:         "Bank Negara Indonesia",
				BranchAddress:    "Jl. Gatot Subroto No. 15, Jakarta",
				CityID:           cities[2 % len(cities)].ID,
				PhoneNumber:      "021-5555789",
				Priority:         3,
			},
		}

		for _, bankAccount := range bankAccounts {
			if err := tx.Create(&bankAccount).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
