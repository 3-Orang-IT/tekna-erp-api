package repository

import (
    "github.com/3-Orang-IT/tekna-erp-api/internal/company/domain"
    "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
    "gorm.io/gorm"
)

type companyRepo struct {
    db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) repository.CompanyRepository {
    return &companyRepo{db: db}
}

func (r *companyRepo) GetCompany() (*entity.Company, error) {
    var company entity.Company
    err := r.db.First(&company).Error
    if err != nil {
        return nil, err
    }
    return &company, nil
}

func (r *companyRepo) CreateCompany(company *entity.Company) error {
    return r.db.Create(company).Error
}

func (r *companyRepo) UpdateCompany(company *entity.Company) error {
    return r.db.Save(company).Error
}

func (r *companyRepo) DeleteCompany(id string) error {
    return r.db.Delete(&entity.Company{}, "id = ?", id).Error
}