package adminRepositoryImpl

import (
	"fmt"
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type chartOfAccountManagementRepo struct {
	db *gorm.DB
}

func NewChartOfAccountManagementRepository(db *gorm.DB) adminRepository.ChartOfAccountManagementRepository {
	return &chartOfAccountManagementRepo{db: db}
}

func (r *chartOfAccountManagementRepo) CreateChartOfAccount(chartOfAccount *entity.ChartOfAccount) error {
	// Get the last COA to generate the next code
	var lastCOA entity.ChartOfAccount
	if err := r.db.Order("id desc").First(&lastCOA).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		// If no records yet, start with COA001
		chartOfAccount.Code = "COA001"
	} else {
		// Parse the last code and increment
		lastID := lastCOA.ID
		newID := lastID + 1
		// Format: COA followed by 3-digit number with leading zeros
		chartOfAccount.Code = "COA" + fmt.Sprintf("%03d", newID)
	}
	
	return r.db.Create(chartOfAccount).Error
}

func (r *chartOfAccountManagementRepo) GetChartOfAccounts(page, limit int, search string) ([]entity.ChartOfAccount, error) {
	var chartOfAccounts []entity.ChartOfAccount
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Order("id ASC").Find(&chartOfAccounts).Error; err != nil {
		return nil, err
	}
	return chartOfAccounts, nil
}

func (r *chartOfAccountManagementRepo) GetChartOfAccountsCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.ChartOfAccount{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *chartOfAccountManagementRepo) GetChartOfAccountByID(id string) (*entity.ChartOfAccount, error) {
	var chartOfAccount entity.ChartOfAccount
	if err := r.db.First(&chartOfAccount, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &chartOfAccount, nil
}

func (r *chartOfAccountManagementRepo) UpdateChartOfAccount(id string, chartOfAccount *entity.ChartOfAccount) error {
	var existingChartOfAccount entity.ChartOfAccount
	if err := r.db.First(&existingChartOfAccount, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingChartOfAccount).Updates(chartOfAccount).Error
}

func (r *chartOfAccountManagementRepo) DeleteChartOfAccount(id string) error {
	var chartOfAccount entity.ChartOfAccount
	if err := r.db.First(&chartOfAccount, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&chartOfAccount).Error
}
