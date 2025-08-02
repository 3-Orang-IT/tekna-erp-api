package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedDivisions(db *gorm.DB) error {
	// Define division list
	divisionsList := []entity.Division{
		{Name: "Human Resources"},
		{Name: "Finance"},
		{Name: "Marketing"},
		{Name: "IT"},
		{Name: "Operations"},
	}

	// Insert divisions into the database
	if err := db.Create(&divisionsList).Error; err != nil {
		log.Printf("Error seeding divisions: %v", err)
		return err
	}

	log.Println("Divisions seeded successfully")
	return nil
}
