package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type BankAccountManagementRepository interface {
	CreateBankAccount(bankAccount *entity.BankAccount) error
	GetBankAccounts(page, limit int, search string) ([]entity.BankAccount, error)
	GetBankAccountsCount(search string) (int64, error)
	GetBankAccountByID(id string) (*entity.BankAccount, error)
	UpdateBankAccount(id string, bankAccount *entity.BankAccount) error
	DeleteBankAccount(id string) error
	GetCities(page, limit int, search string) ([]entity.City, error)
	GetChartOfAccounts(page, limit int, search string) ([]entity.ChartOfAccount, error)
	GetProvinces(page, limit int, search string) ([]entity.Province, error) // Method for fetching provinces with cities
}
