package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type ToDoTemplateManagementUsecase interface {
	CreateToDoTemplate(toDoTemplate *entity.ToDoTemplate) error
	GetToDoTemplates(page, limit int, search string) ([]entity.ToDoTemplate, error)
	GetToDoTemplatesByJobPosition(page, limit int, search string) (map[uint][]entity.ToDoTemplate, int64, error)
	GetToDoTemplateByID(id string) (*entity.ToDoTemplate, error)
	GetToDoTemplatesByJobPositionID(jobPositionID uint) ([]entity.ToDoTemplate, error)
	UpdateToDoTemplate(id string, toDoTemplate *entity.ToDoTemplate) error
	DeleteToDoTemplate(id string) error
	GetJobPositions(page, limit int, search string) ([]entity.JobPosition, error)
}

type toDoTemplateManagementUsecase struct {
	repo adminRepository.ToDoTemplateManagementRepository
}

func NewToDoTemplateManagementUsecase(r adminRepository.ToDoTemplateManagementRepository) ToDoTemplateManagementUsecase {
	return &toDoTemplateManagementUsecase{repo: r}
}

func (u *toDoTemplateManagementUsecase) CreateToDoTemplate(toDoTemplate *entity.ToDoTemplate) error {
	return u.repo.CreateToDoTemplate(toDoTemplate)
}

func (u *toDoTemplateManagementUsecase) GetToDoTemplates(page, limit int, search string) ([]entity.ToDoTemplate, error) {
	return u.repo.GetToDoTemplates(page, limit, search)
}

func (u *toDoTemplateManagementUsecase) GetToDoTemplatesByJobPosition(page, limit int, search string) (map[uint][]entity.ToDoTemplate, int64, error) {
	return u.repo.GetToDoTemplatesByJobPosition(page, limit, search)
}

func (u *toDoTemplateManagementUsecase) GetToDoTemplateByID(id string) (*entity.ToDoTemplate, error) {
	return u.repo.GetToDoTemplateByID(id)
}

func (u *toDoTemplateManagementUsecase) GetToDoTemplatesByJobPositionID(jobPositionID uint) ([]entity.ToDoTemplate, error) {
	return u.repo.GetToDoTemplatesByJobPositionID(jobPositionID)
}

func (u *toDoTemplateManagementUsecase) UpdateToDoTemplate(id string, toDoTemplate *entity.ToDoTemplate) error {
	return u.repo.UpdateToDoTemplate(id, toDoTemplate)
}

func (u *toDoTemplateManagementUsecase) DeleteToDoTemplate(id string) error {
	return u.repo.DeleteToDoTemplate(id)
}

func (u *toDoTemplateManagementUsecase) GetJobPositions(page, limit int, search string) ([]entity.JobPosition, error) {
	return u.repo.GetJobPositions(page, limit, search)
}
