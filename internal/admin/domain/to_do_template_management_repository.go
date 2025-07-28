package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type ToDoTemplateManagementRepository interface {
	CreateToDoTemplate(toDoTemplate *entity.ToDoTemplate) error
	GetToDoTemplates(page, limit int, search string) ([]entity.ToDoTemplate, error)
	GetToDoTemplatesByJobPosition(page, limit int, search string) (map[uint][]entity.ToDoTemplate, int64, error)
	GetToDoTemplateByID(id string) (*entity.ToDoTemplate, error)
	GetToDoTemplatesByJobPositionID(jobPositionID uint) ([]entity.ToDoTemplate, error)
	UpdateToDoTemplate(id string, toDoTemplate *entity.ToDoTemplate) error
	DeleteToDoTemplate(id string) error
	GetJobPositions(page, limit int, search string) ([]entity.JobPosition, error)
}
