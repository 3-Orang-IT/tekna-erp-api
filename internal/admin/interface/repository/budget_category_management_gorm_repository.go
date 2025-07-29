package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type budgetCategoryManagementRepo struct {
	db *gorm.DB
}

func NewBudgetCategoryManagementRepository(db *gorm.DB) adminRepository.BudgetCategoryManagementRepository {
	return &budgetCategoryManagementRepo{db: db}
}

func (r *budgetCategoryManagementRepo) CreateBudgetCategory(category *entity.BudgetCategory) error {
	return r.db.Create(category).Error
}

func (r *budgetCategoryManagementRepo) GetBudgetCategories(page, limit int, search string) ([]entity.BudgetCategory, error) {
	var categories []entity.BudgetCategory
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Preload("ChartOfAccount").Limit(limit).Offset(offset).Order("id ASC").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *budgetCategoryManagementRepo) GetBudgetCategoryByID(id string) (*entity.BudgetCategory, error) {
	var category entity.BudgetCategory
	if err := r.db.Preload("ChartOfAccount").First(&category, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *budgetCategoryManagementRepo) UpdateBudgetCategory(id string, category *entity.BudgetCategory) error {
	var existing entity.BudgetCategory
	if err := r.db.First(&existing, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existing).Updates(category).Error
}

func (r *budgetCategoryManagementRepo) DeleteBudgetCategory(id string) error {
	var category entity.BudgetCategory
	if err := r.db.First(&category, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&category).Error
}

func (r *budgetCategoryManagementRepo) GetBudgetCategoriesCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.BudgetCategory{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *budgetCategoryManagementRepo) GetChartOfAccounts() ([]entity.ChartOfAccount, error) {
	var accounts []entity.ChartOfAccount
	if err := r.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}
