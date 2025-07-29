package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedAreas(db *gorm.DB) error {
	areas := []entity.Area{
		{Name: "North Area"},
		{Name: "South Area"},
		{Name: "East Area"},
		{Name: "West Area"},
	}

	for _, area := range areas {
		var existingArea entity.Area
		if err := db.Where("name = ?", area.Name).First(&existingArea).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&area).Error; err != nil {
				log.Printf("Error seeding area %s: %v", area.Name, err)
				return err
			}
			log.Printf("Successfully added area %s", area.Name)
		}
	}

	return nil
}
