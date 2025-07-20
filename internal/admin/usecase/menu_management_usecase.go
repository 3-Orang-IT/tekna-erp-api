package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type MenuManagementUsecase interface {
	CreateMenu(menu *entity.Menu) error
	GetMenus(page, limit int) ([]entity.Menu, error)
	GetMenuByID(id string) (*entity.Menu, error)
	UpdateMenu(id string, menu *entity.Menu) error
	DeleteMenu(id string) error
}

type menuManagementUsecase struct {
	repo adminRepository.MenuManagementRepository
}

func NewMenuManagementUsecase(r adminRepository.MenuManagementRepository) MenuManagementUsecase {
	return &menuManagementUsecase{repo: r}
}

func (u *menuManagementUsecase) CreateMenu(menu *entity.Menu) error {
	return u.repo.CreateMenu(menu)
}

func (u *menuManagementUsecase) GetMenus(page, limit int) ([]entity.Menu, error) {
	return u.repo.GetMenus(page, limit)
}

func (u *menuManagementUsecase) GetMenuByID(id string) (*entity.Menu, error) {
	return u.repo.GetMenuByID(id)
}

func (u *menuManagementUsecase) UpdateMenu(id string, menu *entity.Menu) error {
	return u.repo.UpdateMenu(id, menu)
}

func (u *menuManagementUsecase) DeleteMenu(id string) error {
	return u.repo.DeleteMenu(id)
}
