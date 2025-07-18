package usecase

import (
    "github.com/3-Orang-IT/tekna-erp-api/internal/company/domain"
    "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type CompanyUsecase interface {
    GetCompany() (*entity.Company, error)
    UpdateCompany(company *entity.Company) error
    CreateCompany(company *entity.Company) error
    DeleteCompany(id string) error
}

type companyUsecase struct {
    repo repository.CompanyRepository
}

func NewCompanyUsecase(r repository.CompanyRepository) CompanyUsecase {
    return &companyUsecase{repo: r}
}

func (u *companyUsecase) GetCompany() (*entity.Company, error) {
    return u.repo.GetCompany()
}

func (u *companyUsecase) CreateCompany(company *entity.Company) error {
    return u.repo.CreateCompany(company)
}

func (u *companyUsecase) UpdateCompany(company *entity.Company) error {
    return u.repo.UpdateCompany(company)
}

func (u *companyUsecase) DeleteCompany(id string) error {
    return u.repo.DeleteCompany(id)
}