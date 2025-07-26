package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedBusinessUnits(db *gorm.DB) error {
	businessUnits := []entity.BusinessUnit{
		{Name: "Finance"},
		{Name: "Human Resources"},
		{Name: "Operations"},
		{Name: "Marketing"},
		{Name: "IT"},
	}

	for _, businessUnit := range businessUnits {
		if err := db.FirstOrCreate(&entity.BusinessUnit{}, businessUnit).Error; err != nil {
			log.Printf("Error seeding business unit %s: %v", businessUnit.Name, err)
			return err
		}
	}

	return nil
}
