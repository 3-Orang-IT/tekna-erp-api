package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type MenuManagementRepository interface {
	CreateMenu(menu *entity.Menu) error
	GetMenus(page, limit int) ([]entity.Menu, error)
	GetMenuByID(id string) (*entity.Menu, error)
	UpdateMenu(id string, menu *entity.Menu) error
	DeleteMenu(id string) error
}