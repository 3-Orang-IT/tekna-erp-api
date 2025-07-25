package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedJobPositions(db *gorm.DB) error {
	jobPositions := []entity.JobPosition{
		{Name: "Manager"},
		{Name: "Supervisor"},
		{Name: "Staff"},
		{Name: "Intern"},
	}

	for _, jp := range jobPositions {
		var existing entity.JobPosition
		if err := db.Where("name = ?", jp.Name).First(&existing).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&jp).Error; err != nil {
				log.Printf("Failed to create job position %s: %v", jp.Name, err)
				return err
			}
			log.Printf("Successfully added job position %s", jp.Name)
		}
	}
	return nil
}
