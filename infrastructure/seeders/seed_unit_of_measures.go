package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedUnitOfMeasures(db *gorm.DB) error {
	unitOfMeasures := []entity.UnitOfMeasure{
		{Name: "Kilogram", Abbreviation: "kg"},
		{Name: "Meter", Abbreviation: "m"},
		{Name: "Liter", Abbreviation: "l"},
		{Name: "Piece", Abbreviation: "pcs"},
		{Name: "Box", Abbreviation: "box"},
	}

	for _, unitOfMeasure := range unitOfMeasures {
		if err := db.FirstOrCreate(&entity.UnitOfMeasure{}, unitOfMeasure).Error; err != nil {
			log.Printf("Error seeding unit of measure %s: %v", unitOfMeasure.Name, err)
			return err
		}
	}

	return nil
}
