package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type jobPositionManagementRepo struct {
	db *gorm.DB
}

func NewJobPositionManagementRepository(db *gorm.DB) adminRepository.JobPositionManagementRepository {
	return &jobPositionManagementRepo{db: db}
}

func (r *jobPositionManagementRepo) CreateJobPosition(jobPosition *entity.JobPosition) error {
	return r.db.Create(jobPosition).Error
}

func (r *jobPositionManagementRepo) GetJobPositions(page, limit int, search string) ([]entity.JobPosition, error) {
	   var jobPositions []entity.JobPosition
	   offset := (page - 1) * limit
	   query := r.db
	   if search != "" {
			   query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	   }
	   if err := query.Limit(limit).Offset(offset).Find(&jobPositions).Error; err != nil {
			   return nil, err
	   }
	   return jobPositions, nil
}

func (r *jobPositionManagementRepo) GetJobPositionByID(id string) (*entity.JobPosition, error) {
	var jobPosition entity.JobPosition
	if err := r.db.First(&jobPosition, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &jobPosition, nil
}

func (r *jobPositionManagementRepo) UpdateJobPosition(id string, jobPosition *entity.JobPosition) error {
	var existingJobPosition entity.JobPosition
	if err := r.db.First(&existingJobPosition, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingJobPosition).Updates(jobPosition).Error
}

func (r *jobPositionManagementRepo) DeleteJobPosition(id string) error {
	var jobPosition entity.JobPosition
	if err := r.db.First(&jobPosition, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&jobPosition).Error
}
