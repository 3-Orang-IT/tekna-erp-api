package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type bankAccountManagementRepo struct {
	db *gorm.DB
}

func NewBankAccountManagementRepository(db *gorm.DB) adminRepository.BankAccountManagementRepository {
	return &bankAccountManagementRepo{db: db}
}

func (r *bankAccountManagementRepo) CreateBankAccount(bankAccount *entity.BankAccount) error {
	return r.db.Create(bankAccount).Error
}

func (r *bankAccountManagementRepo) GetBankAccounts(page, limit int, search string) ([]entity.BankAccount, error) {
	var bankAccounts []entity.BankAccount
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(bank_name) LIKE ? OR account_number LIKE ?", 
			"%"+strings.ToLower(search)+"%", "%"+search+"%")
	}
	if err := query.Preload("City.Province").Preload("City").Preload("ChartOfAccount").
		Limit(limit).Offset(offset).Order("id ASC").Find(&bankAccounts).Error; err != nil {
		return nil, err
	}
	return bankAccounts, nil
}

func (r *bankAccountManagementRepo) GetBankAccountsCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.BankAccount{})
	if search != "" {
		query = query.Where("LOWER(bank_name) LIKE ? OR account_number LIKE ?", 
			"%"+strings.ToLower(search)+"%", "%"+search+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *bankAccountManagementRepo) GetBankAccountByID(id string) (*entity.BankAccount, error) {
	var bankAccount entity.BankAccount
	if err := r.db.Preload("City.Province").Preload("City").Preload("ChartOfAccount").
		First(&bankAccount, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &bankAccount, nil
}

func (r *bankAccountManagementRepo) UpdateBankAccount(id string, bankAccount *entity.BankAccount) error {
	var existingBankAccount entity.BankAccount
	if err := r.db.First(&existingBankAccount, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingBankAccount).Updates(bankAccount).Error
}

func (r *bankAccountManagementRepo) DeleteBankAccount(id string) error {
	var bankAccount entity.BankAccount
	if err := r.db.First(&bankAccount, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&bankAccount).Error
}

func (r *bankAccountManagementRepo) GetCities(page, limit int, search string) ([]entity.City, error) {
	var cities []entity.City
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Preload("Province").Limit(limit).Offset(offset).Find(&cities).Error; err != nil {
		return nil, err
	}
	return cities, nil
}

func (r *bankAccountManagementRepo) GetChartOfAccounts(page, limit int, search string) ([]entity.ChartOfAccount, error) {
	var chartOfAccounts []entity.ChartOfAccount
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ? OR account_code LIKE ? OR LOWER(bank_name) LIKE ?",
			"%"+strings.ToLower(search)+"%", "%"+search+"%", "%"+search+"%")
	}
	if err := query.Limit(limit).Offset(offset).Find(&chartOfAccounts).Error; err != nil {
		return nil, err
	}
	return chartOfAccounts, nil
}

// GetProvinces fetches provinces with their cities for the bank account form
func (r *bankAccountManagementRepo) GetProvinces(page, limit int, search string) ([]entity.Province, error) {
	var provinces []entity.Province
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Preload("Cities").Limit(limit).Offset(offset).Find(&provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}
