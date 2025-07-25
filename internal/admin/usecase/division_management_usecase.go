package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type DivisionManagementUsecase interface {
	CreateDivision(division *entity.Division) error
	GetDivisions(page, limit int, search string) ([]entity.Division, error)
	GetDivisionByID(id string) (*entity.Division, error)
	UpdateDivision(id string, division *entity.Division) error
	DeleteDivision(id string) error
}

type divisionManagementUsecase struct {
	repo adminRepository.DivisionManagementRepository
}

func NewDivisionManagementUsecase(r adminRepository.DivisionManagementRepository) DivisionManagementUsecase {
	return &divisionManagementUsecase{repo: r}
}

func (u *divisionManagementUsecase) CreateDivision(division *entity.Division) error {
	return u.repo.CreateDivision(division)
}

func (u *divisionManagementUsecase) GetDivisions(page, limit int, search string) ([]entity.Division, error) {
	divisions, err := u.repo.GetDivisions(page, limit, search)
	if err != nil {
		return nil, err
	}
	return divisions, nil
}

func (u *divisionManagementUsecase) GetDivisionByID(id string) (*entity.Division, error) {
	return u.repo.GetDivisionByID(id)
}

func (u *divisionManagementUsecase) UpdateDivision(id string, division *entity.Division) error {
	return u.repo.UpdateDivision(id, division)
}

func (u *divisionManagementUsecase) DeleteDivision(id string) error {
	return u.repo.DeleteDivision(id)
}
