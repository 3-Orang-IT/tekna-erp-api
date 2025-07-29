package adminRepositoryImpl

import (
	"fmt"
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type areaManagementRepo struct {
	db *gorm.DB
}

func NewAreaManagementRepository(db *gorm.DB) adminRepository.AreaManagementRepository {
	return &areaManagementRepo{db: db}
}

func (r *areaManagementRepo) CreateArea(area *entity.Area) error {
	// Ensure we're not trying to set the ID manually
	area.ID = 0
	
	// Check for existing area with same name
	var existingArea entity.Area
	if err := r.db.Where("LOWER(name) = LOWER(?)", area.Name).First(&existingArea).Error; err == nil {
		return gorm.ErrDuplicatedKey
	}
	
	return r.db.Create(area).Error
}

func (r *areaManagementRepo) GetAreas(page, limit int, search string) ([]entity.Area, error) {
	var areas []entity.Area
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(areas.name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	
	// Using joins for many-to-many relationship
	err := query.
		Preload("Employees").
		Preload("Employees.User").
		Limit(limit).
		Offset(offset).
		Order("areas.id ASC").
		Find(&areas).Error
		
	if err != nil {
		return nil, err
	}
	return areas, nil
}

func (r *areaManagementRepo) GetAreaByID(id string) (*entity.Area, error) {
	var area entity.Area
	err := r.db.
		Preload("Employees").
		Preload("Employees.User").
		First(&area, "areas.id = ?", id).Error
		
	if err != nil {
		return nil, err
	}
	return &area, nil
}

func (r *areaManagementRepo) UpdateArea(id string, area *entity.Area) error {
	// Find the area we want to update
	var existingArea entity.Area
	if err := r.db.First(&existingArea, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	
	// Check if there's already another area with the new name
	if area.Name != "" && area.Name != existingArea.Name {
		var duplicateArea entity.Area
		if err := r.db.Where("LOWER(name) = LOWER(?) AND id != ?", area.Name, id).First(&duplicateArea).Error; err == nil {
			return gorm.ErrDuplicatedKey
		}
	}
	
	return r.db.Model(&existingArea).Updates(area).Error
}

func (r *areaManagementRepo) DeleteArea(id string) error {
	var area entity.Area
	if err := r.db.First(&area, "id = ?", id).Error; err != nil {
		return err
	}
	
	// Begin a transaction to ensure data integrity
	tx := r.db.Begin()
	
	// First, check if there are customers using this area
	var customerCount int64
	if err := tx.Model(&entity.Customer{}).Where("area_id = ?", area.ID).Count(&customerCount).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// If there are customers using this area, we need to find or create a replacement
	if customerCount > 0 {
		// Find another area to use as replacement if one exists
		var replacementArea entity.Area
		if err := tx.Where("id != ?", area.ID).First(&replacementArea).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				tx.Rollback()
				return err
			}
			// No replacement area found, we can't delete this area
			tx.Rollback()
			return fmt.Errorf("cannot delete the last remaining area when customers are assigned to it")
		}
		
		// Update all customers that use this area to use the replacement area
		if err := tx.Model(&entity.Customer{}).Where("area_id = ?", area.ID).Update("area_id", replacementArea.ID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	
	// Remove the area associations from employees
	if err := tx.Exec("DELETE FROM employee_areas WHERE area_id = ?", area.ID).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// Delete the area
	if err := tx.Delete(&area).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// Commit the transaction
	return tx.Commit().Error
}

// Method to get total count of areas for pagination
func (r *areaManagementRepo) GetAreasCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Area{})
	if search != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
