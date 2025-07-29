package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type toDoTemplateManagementRepo struct {
	db *gorm.DB
}

func NewToDoTemplateManagementRepository(db *gorm.DB) adminRepository.ToDoTemplateManagementRepository {
	return &toDoTemplateManagementRepo{db: db}
}

func (r *toDoTemplateManagementRepo) CreateToDoTemplate(toDoTemplate *entity.ToDoTemplate) error {
	return r.db.Create(toDoTemplate).Error
}

func (r *toDoTemplateManagementRepo) GetToDoTemplates(page, limit int, search string) ([]entity.ToDoTemplate, error) {
	var toDoTemplates []entity.ToDoTemplate
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(activity) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Preload("JobPosition").Limit(limit).Offset(offset).Order("id ASC").Find(&toDoTemplates).Error; err != nil {
		return nil, err
	}
	return toDoTemplates, nil
}

func (r *toDoTemplateManagementRepo) GetToDoTemplatesCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.ToDoTemplate{})
	if search != "" {
		query = query.Where("LOWER(activity) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *toDoTemplateManagementRepo) GetToDoTemplatesByJobPosition(page, limit int, search string) (map[uint][]entity.ToDoTemplate, int64, error) {
	var toDoTemplates []entity.ToDoTemplate
	var total int64
	
	// First, get the total count of to-do templates for pagination info
	countQuery := r.db.Model(&entity.ToDoTemplate{})
	if search != "" {
		countQuery = countQuery.Where("LOWER(activity) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	// Query to get all job positions with their templates
	query := r.db
	if search != "" {
		query = query.Where("LOWER(activity) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	
	// Get all templates - we don't paginate by job positions anymore
	// Instead, we'll return all job positions but paginate the total templates count
	if err := query.Preload("JobPosition").
		Order("job_position_id, priority, order_number").
		Find(&toDoTemplates).Error; err != nil {
		return nil, 0, err
	}
	
	// Group templates by job position ID
	grouped := make(map[uint][]entity.ToDoTemplate)
	for _, template := range toDoTemplates {
		grouped[template.JobPositionID] = append(grouped[template.JobPositionID], template)
	}
	
	return grouped, total, nil
}

func (r *toDoTemplateManagementRepo) GetToDoTemplateByID(id string) (*entity.ToDoTemplate, error) {
	var toDoTemplate entity.ToDoTemplate
	if err := r.db.Preload("JobPosition").First(&toDoTemplate, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &toDoTemplate, nil
}

func (r *toDoTemplateManagementRepo) GetToDoTemplatesByJobPositionID(jobPositionID uint) ([]entity.ToDoTemplate, error) {
	var templates []entity.ToDoTemplate
	if err := r.db.Where("job_position_id = ?", jobPositionID).
		Preload("JobPosition").
		Order("priority, order_number").
		Find(&templates).Error; err != nil {
		return nil, err
	}
	return templates, nil
}

func (r *toDoTemplateManagementRepo) UpdateToDoTemplate(id string, toDoTemplate *entity.ToDoTemplate) error {
	var existingToDoTemplate entity.ToDoTemplate
	if err := r.db.First(&existingToDoTemplate, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingToDoTemplate).Updates(toDoTemplate).Error
}

func (r *toDoTemplateManagementRepo) DeleteToDoTemplate(id string) error {
	var toDoTemplate entity.ToDoTemplate
	if err := r.db.First(&toDoTemplate, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&toDoTemplate).Error
}

func (r *toDoTemplateManagementRepo) GetJobPositions(page, limit int, search string) ([]entity.JobPosition, error) {
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
