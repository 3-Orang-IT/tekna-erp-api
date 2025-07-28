package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedAreas(db *gorm.DB) error {
	areas := []entity.Area{
		{ID: 1, Name: "North Area"},
		{ID: 2, Name: "South Area"},
		{ID: 3, Name: "East Area"},
		{ID: 4, Name: "West Area"},
	}

	for _, area := range areas {
		if err := db.FirstOrCreate(&area, entity.Area{ID: area.ID}).Error; err != nil {
			log.Printf("Error seeding area with ID %d: %v", area.ID, err)
			return err
		}
	}

	return nil
}
