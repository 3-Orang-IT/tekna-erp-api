package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type BudgetCategoryManagementRepository interface {
	CreateBudgetCategory(category *entity.BudgetCategory) error
	GetBudgetCategories(page, limit int, search string) ([]entity.BudgetCategory, error)
	GetBudgetCategoryByID(id string) (*entity.BudgetCategory, error)
	UpdateBudgetCategory(id string, category *entity.BudgetCategory) error
	DeleteBudgetCategory(id string) error
	GetBudgetCategoriesCount(search string) (int64, error)
	GetChartOfAccounts() ([]entity.ChartOfAccount, error)
}
