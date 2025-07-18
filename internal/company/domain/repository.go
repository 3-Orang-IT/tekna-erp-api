package repository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type CompanyRepository interface {
    GetCompany() (*entity.Company, error)
    UpdateCompany(company *entity.Company) error
    CreateCompany(company *entity.Company) error
    DeleteCompany(id string) error
}