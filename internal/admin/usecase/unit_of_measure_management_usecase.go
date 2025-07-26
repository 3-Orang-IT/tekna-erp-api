package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type UnitOfMeasureManagementUsecase interface {
	CreateUnitOfMeasure(unitOfMeasure *entity.UnitOfMeasure) error
	GetUnitOfMeasures(page, limit int, search string) ([]entity.UnitOfMeasure, error)
	GetUnitOfMeasureByID(id string) (*entity.UnitOfMeasure, error)
	UpdateUnitOfMeasure(id string, unitOfMeasure *entity.UnitOfMeasure) error
	DeleteUnitOfMeasure(id string) error
}

type unitOfMeasureManagementUsecase struct {
	repo adminRepository.UnitOfMeasureManagementRepository
}

func NewUnitOfMeasureManagementUsecase(repo adminRepository.UnitOfMeasureManagementRepository) UnitOfMeasureManagementUsecase {
	return &unitOfMeasureManagementUsecase{repo: repo}
}

func (u *unitOfMeasureManagementUsecase) CreateUnitOfMeasure(unitOfMeasure *entity.UnitOfMeasure) error {
	return u.repo.CreateUnitOfMeasure(unitOfMeasure)
}

func (u *unitOfMeasureManagementUsecase) GetUnitOfMeasures(page, limit int, search string) ([]entity.UnitOfMeasure, error) {
	return u.repo.GetUnitOfMeasures(page, limit, search)
}

func (u *unitOfMeasureManagementUsecase) GetUnitOfMeasureByID(id string) (*entity.UnitOfMeasure, error) {
	return u.repo.GetUnitOfMeasureByID(id)
}

func (u *unitOfMeasureManagementUsecase) UpdateUnitOfMeasure(id string, unitOfMeasure *entity.UnitOfMeasure) error {
	return u.repo.UpdateUnitOfMeasure(id, unitOfMeasure)
}

func (u *unitOfMeasureManagementUsecase) DeleteUnitOfMeasure(id string) error {
	return u.repo.DeleteUnitOfMeasure(id)
}
