package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type CompanyManagementRepository interface {
	CreateCompany(company *entity.Company) error
	GetCompanies(page, limit int, search string) ([]entity.Company, error)
	GetCompanyByID(id string) (*entity.Company, error)
	UpdateCompany(id string, company *entity.Company) error
	DeleteCompany(id string) error
	// Method to fetch cities for the edit page
	GetCities(page, limit int, search string) ([]entity.City, error)
	// Method to fetch provinces with their cities for the add page
	GetProvinces(page, limit int, search string) ([]entity.Province, error)
	// Method to get total count of companies for pagination
	GetCompaniesCount(search string) (int64, error)
}
