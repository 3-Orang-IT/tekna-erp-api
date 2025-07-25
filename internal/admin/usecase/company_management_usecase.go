package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type CompanyManagementUsecase interface {
	CreateCompany(company *entity.Company) error
	GetCompanies(page, limit int, search string) ([]entity.Company, error)
	GetCompanyByID(id string) (*entity.Company, error)
	UpdateCompany(id string, company *entity.Company) error
	DeleteCompany(id string) error
	GetCities(page, limit int, search string) ([]entity.City, error) // New method for fetching cities
}

type companyManagementUsecase struct {
	repo adminRepository.CompanyManagementRepository
}

func NewCompanyManagementUsecase(r adminRepository.CompanyManagementRepository) CompanyManagementUsecase {
	return &companyManagementUsecase{repo: r}
}

func (u *companyManagementUsecase) CreateCompany(company *entity.Company) error {
	return u.repo.CreateCompany(company)
}

func (u *companyManagementUsecase) GetCompanies(page, limit int, search string) ([]entity.Company, error) {
	return u.repo.GetCompanies(page, limit, search)
}

func (u *companyManagementUsecase) GetCompanyByID(id string) (*entity.Company, error) {
	return u.repo.GetCompanyByID(id)
}

func (u *companyManagementUsecase) UpdateCompany(id string, company *entity.Company) error {
	return u.repo.UpdateCompany(id, company)
}

func (u *companyManagementUsecase) DeleteCompany(id string) error {
	return u.repo.DeleteCompany(id)
}

// GetCities fetches cities for the edit page
func (u *companyManagementUsecase) GetCities(page, limit int, search string) ([]entity.City, error) {
	return u.repo.GetCities(page, limit, search)
}
