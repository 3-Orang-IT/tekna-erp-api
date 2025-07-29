package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type BudgetCategoryManagementUsecase interface {
	CreateBudgetCategory(category *entity.BudgetCategory) error
	GetBudgetCategories(page, limit int, search string) ([]entity.BudgetCategory, error)
	GetBudgetCategoryByID(id string) (*entity.BudgetCategory, error)
	UpdateBudgetCategory(id string, category *entity.BudgetCategory) error
	DeleteBudgetCategory(id string) error
	GetBudgetCategoriesCount(search string) (int64, error)
	GetChartOfAccounts() ([]entity.ChartOfAccount, error)
}

type budgetCategoryManagementUsecase struct {
	repo adminRepository.BudgetCategoryManagementRepository
}

func NewBudgetCategoryManagementUsecase(r adminRepository.BudgetCategoryManagementRepository) BudgetCategoryManagementUsecase {
	return &budgetCategoryManagementUsecase{repo: r}
}

func (u *budgetCategoryManagementUsecase) CreateBudgetCategory(category *entity.BudgetCategory) error {
	return u.repo.CreateBudgetCategory(category)
}

func (u *budgetCategoryManagementUsecase) GetBudgetCategories(page, limit int, search string) ([]entity.BudgetCategory, error) {
	return u.repo.GetBudgetCategories(page, limit, search)
}

func (u *budgetCategoryManagementUsecase) GetBudgetCategoryByID(id string) (*entity.BudgetCategory, error) {
	return u.repo.GetBudgetCategoryByID(id)
}

func (u *budgetCategoryManagementUsecase) UpdateBudgetCategory(id string, category *entity.BudgetCategory) error {
	return u.repo.UpdateBudgetCategory(id, category)
}

func (u *budgetCategoryManagementUsecase) DeleteBudgetCategory(id string) error {
	return u.repo.DeleteBudgetCategory(id)
}

func (u *budgetCategoryManagementUsecase) GetBudgetCategoriesCount(search string) (int64, error) {
	return u.repo.GetBudgetCategoriesCount(search)
}

func (u *budgetCategoryManagementUsecase) GetChartOfAccounts() ([]entity.ChartOfAccount, error) {
	return u.repo.GetChartOfAccounts()
}
