package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type companyManagementRepo struct {
	db *gorm.DB
}

func NewCompanyManagementRepository(db *gorm.DB) adminRepository.CompanyManagementRepository {
	return &companyManagementRepo{db: db}
}

func (r *companyManagementRepo) CreateCompany(company *entity.Company) error {
	return r.db.Create(company).Error
}

func (r *companyManagementRepo) GetCompanies(page, limit int, search string) ([]entity.Company, error) {
	var companies []entity.Company
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Preload("City.Province").Preload("City").Limit(limit).Offset(offset).Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (r *companyManagementRepo) GetCompanyByID(id string) (*entity.Company, error) {
	var company entity.Company
	if err := r.db.Preload("City.Province").Preload("City").First(&company, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &company, nil
}

func (r *companyManagementRepo) UpdateCompany(id string, company *entity.Company) error {
	var existingCompany entity.Company
	if err := r.db.First(&existingCompany, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingCompany).Updates(company).Error
}

func (r *companyManagementRepo) DeleteCompany(id string) error {
	var company entity.Company
	if err := r.db.First(&company, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&company).Error
}

// Method to fetch cities for the edit page
func (r *companyManagementRepo) GetCities(page, limit int, search string) ([]entity.City, error) {
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

// Method to fetch provinces with their cities for the add page
func (r *companyManagementRepo) GetProvinces(page, limit int, search string) ([]entity.Province, error) {
	var provinces []entity.Province
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(provinces.name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Preload("Cities").Limit(limit).Offset(offset).Find(&provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}
