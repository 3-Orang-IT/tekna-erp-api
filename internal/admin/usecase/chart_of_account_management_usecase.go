package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type ChartOfAccountManagementUsecase interface {
	CreateChartOfAccount(chartOfAccount *entity.ChartOfAccount) error
	GetChartOfAccounts(page, limit int, search string) ([]entity.ChartOfAccount, error)
	GetChartOfAccountByID(id string) (*entity.ChartOfAccount, error)
	UpdateChartOfAccount(id string, chartOfAccount *entity.ChartOfAccount) error
	DeleteChartOfAccount(id string) error
}

type chartOfAccountManagementUsecase struct {
	repo adminRepository.ChartOfAccountManagementRepository
}

func NewChartOfAccountManagementUsecase(r adminRepository.ChartOfAccountManagementRepository) ChartOfAccountManagementUsecase {
	return &chartOfAccountManagementUsecase{repo: r}
}

func (u *chartOfAccountManagementUsecase) CreateChartOfAccount(chartOfAccount *entity.ChartOfAccount) error {
	return u.repo.CreateChartOfAccount(chartOfAccount)
}

func (u *chartOfAccountManagementUsecase) GetChartOfAccounts(page, limit int, search string) ([]entity.ChartOfAccount, error) {
	return u.repo.GetChartOfAccounts(page, limit, search)
}

func (u *chartOfAccountManagementUsecase) GetChartOfAccountByID(id string) (*entity.ChartOfAccount, error) {
	return u.repo.GetChartOfAccountByID(id)
}

func (u *chartOfAccountManagementUsecase) UpdateChartOfAccount(id string, chartOfAccount *entity.ChartOfAccount) error {
	return u.repo.UpdateChartOfAccount(id, chartOfAccount)
}

func (u *chartOfAccountManagementUsecase) DeleteChartOfAccount(id string) error {
	return u.repo.DeleteChartOfAccount(id)
}
