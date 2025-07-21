package adminRepositoryImpl

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type menuManagementRepo struct {
	db *gorm.DB
}

func NewMenuManagementRepository(db *gorm.DB) adminRepository.MenuManagementRepository {
	return &menuManagementRepo{db: db}
}

func (r *menuManagementRepo) CreateMenu(menu *entity.Menu) error {
	return r.db.Create(menu).Error
}

func (r *menuManagementRepo) GetMenus(page, limit int) ([]entity.Menu, error) {
	var menus []entity.Menu
	offset := (page - 1) * limit
	if err := r.db.Limit(limit).Offset(offset).Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func (r *menuManagementRepo) GetMenuByID(id string) (*entity.Menu, error) {
	var menu entity.Menu
	if err := r.db.First(&menu, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func (r *menuManagementRepo) UpdateMenu(id string, menu *entity.Menu) error {
	var existingMenu entity.Menu
	if err := r.db.First(&existingMenu, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingMenu).Updates(menu).Error
}

func (r *menuManagementRepo) DeleteMenu(id string) error {
	var menu entity.Menu
	if err := r.db.First(&menu, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&menu).Error
}
