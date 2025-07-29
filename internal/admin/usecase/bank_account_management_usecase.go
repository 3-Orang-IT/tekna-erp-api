package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type BankAccountManagementUsecase interface {
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

type bankAccountManagementUsecase struct {
	repo adminRepository.BankAccountManagementRepository
}

func NewBankAccountManagementUsecase(r adminRepository.BankAccountManagementRepository) BankAccountManagementUsecase {
	return &bankAccountManagementUsecase{repo: r}
}

func (u *bankAccountManagementUsecase) CreateBankAccount(bankAccount *entity.BankAccount) error {
	return u.repo.CreateBankAccount(bankAccount)
}

func (u *bankAccountManagementUsecase) GetBankAccounts(page, limit int, search string) ([]entity.BankAccount, error) {
	return u.repo.GetBankAccounts(page, limit, search)
}

func (u *bankAccountManagementUsecase) GetBankAccountsCount(search string) (int64, error) {
	return u.repo.GetBankAccountsCount(search)
}

func (u *bankAccountManagementUsecase) GetBankAccountByID(id string) (*entity.BankAccount, error) {
	return u.repo.GetBankAccountByID(id)
}

func (u *bankAccountManagementUsecase) UpdateBankAccount(id string, bankAccount *entity.BankAccount) error {
	return u.repo.UpdateBankAccount(id, bankAccount)
}

func (u *bankAccountManagementUsecase) DeleteBankAccount(id string) error {
	return u.repo.DeleteBankAccount(id)
}

func (u *bankAccountManagementUsecase) GetCities(page, limit int, search string) ([]entity.City, error) {
	return u.repo.GetCities(page, limit, search)
}

func (u *bankAccountManagementUsecase) GetChartOfAccounts(page, limit int, search string) ([]entity.ChartOfAccount, error) {
	return u.repo.GetChartOfAccounts(page, limit, search)
}

// GetProvinces fetches provinces with their cities for the bank account form
func (u *bankAccountManagementUsecase) GetProvinces(page, limit int, search string) ([]entity.Province, error) {
	return u.repo.GetProvinces(page, limit, search)
}
