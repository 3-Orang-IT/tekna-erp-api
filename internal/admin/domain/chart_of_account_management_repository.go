package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type ChartOfAccountManagementRepository interface {
	CreateChartOfAccount(chartOfAccount *entity.ChartOfAccount) error
	GetChartOfAccounts(page, limit int, search string) ([]entity.ChartOfAccount, error)
	GetChartOfAccountByID(id string) (*entity.ChartOfAccount, error)
	UpdateChartOfAccount(id string, chartOfAccount *entity.ChartOfAccount) error
	DeleteChartOfAccount(id string) error
}
