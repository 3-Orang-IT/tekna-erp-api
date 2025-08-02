package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type JobPositionManagementRepository interface {
	CreateJobPosition(jobPosition *entity.JobPosition) error
	GetJobPositions(page, limit int, search string) ([]entity.JobPosition, error)
	GetJobPositionByID(id string) (*entity.JobPosition, error)
	UpdateJobPosition(id string, jobPosition *entity.JobPosition) error
	DeleteJobPosition(id string) error
	// Method to get total count of job positions for pagination
	GetJobPositionsCount(search string) (int64, error)
}
