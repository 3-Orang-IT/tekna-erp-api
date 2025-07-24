package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type CompanyManagementRepository interface {
	CreateCompany(company *entity.Company) error
	GetCompanies(page, limit int) ([]entity.Company, error)
	GetCompanyByID(id string) (*entity.Company, error)
	UpdateCompany(id string, company *entity.Company) error
	DeleteCompany(id string) error
}
