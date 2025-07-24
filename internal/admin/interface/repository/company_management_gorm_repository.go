package adminRepositoryImpl

import (
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

func (r *companyManagementRepo) GetCompanies(page, limit int) ([]entity.Company, error) {
	var companies []entity.Company
	offset := (page - 1) * limit
	if err := r.db.Preload("City.Province").Preload("City").Limit(limit).Offset(offset).Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}

func (r *companyManagementRepo) GetCompanyByID(id string) (*entity.Company, error) {
	var company entity.Company
	if err := r.db.First(&company, "id = ?", id).Error; err != nil {
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
