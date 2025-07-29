package seeders

import (
	"time"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedDocumentCategories(db *gorm.DB) error {
	var count int64
	db.Model(&entity.DocumentCategory{}).Count(&count)
	if count > 0 {
		return nil // Document categories already seeded
	}

	documentCategories := []entity.DocumentCategory{
		{
			Name:      "Contract",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Policy",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Report",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Guideline",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Standard Operating Procedure",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return db.Create(&documentCategories).Error
}
