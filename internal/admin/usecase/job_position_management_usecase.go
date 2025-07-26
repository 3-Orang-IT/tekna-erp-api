package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type JobPositionManagementUsecase interface {
	CreateJobPosition(jobPosition *entity.JobPosition) error
GetJobPositions(page, limit int, search string) ([]entity.JobPosition, error)
	GetJobPositionByID(id string) (*entity.JobPosition, error)
	UpdateJobPosition(id string, jobPosition *entity.JobPosition) error
	DeleteJobPosition(id string) error
}

type jobPositionManagementUsecase struct {
	repo adminRepository.JobPositionManagementRepository
}

func NewJobPositionManagementUsecase(r adminRepository.JobPositionManagementRepository) JobPositionManagementUsecase {
	return &jobPositionManagementUsecase{repo: r}
}

func (u *jobPositionManagementUsecase) CreateJobPosition(jobPosition *entity.JobPosition) error {
	return u.repo.CreateJobPosition(jobPosition)
}

func (u *jobPositionManagementUsecase) GetJobPositions(page, limit int, search string) ([]entity.JobPosition, error) {
	   return u.repo.GetJobPositions(page, limit, search)
}

func (u *jobPositionManagementUsecase) GetJobPositionByID(id string) (*entity.JobPosition, error) {
	return u.repo.GetJobPositionByID(id)
}

func (u *jobPositionManagementUsecase) UpdateJobPosition(id string, jobPosition *entity.JobPosition) error {
	return u.repo.UpdateJobPosition(id, jobPosition)
}

func (u *jobPositionManagementUsecase) DeleteJobPosition(id string) error {
	return u.repo.DeleteJobPosition(id)
}
