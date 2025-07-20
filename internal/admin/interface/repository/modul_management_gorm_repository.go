package adminRepositoryImpl

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type modulManagementRepo struct {
	db *gorm.DB
}

func NewModulManagementRepository(db *gorm.DB) adminRepository.ModulManagementRepository {
	return &modulManagementRepo{db: db}
}

func (r *modulManagementRepo) CreateModul(modul *entity.Modul) error {
	return r.db.Create(modul).Error
}

func (r *modulManagementRepo) GetModuls() ([]entity.Modul, error) {
	var moduls []entity.Modul
	if err := r.db.Find(&moduls).Error; err != nil {
		return nil, err
	}
	return moduls, nil
}

func (r *modulManagementRepo) GetModulByID(id string) (*entity.Modul, error) {
	var modul entity.Modul
	if err := r.db.First(&modul, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &modul, nil
}

func (r *modulManagementRepo) UpdateModul(id string, modul *entity.Modul) error {
	var existingModul entity.Modul
	if err := r.db.First(&existingModul, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingModul).Updates(modul).Error
}

func (r *modulManagementRepo) DeleteModul(id string) error {
	var modul entity.Modul
	if err := r.db.First(&modul, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&modul).Error
}
