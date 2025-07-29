package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedTravelCosts(db *gorm.DB) error {
	var count int64
	db.Model(&entity.TravelCost{}).Count(&count)
	if count > 0 {
		log.Println("Travel costs data already exist")
		return nil
	}

	log.Println("Seeding travel costs data...")

	travelCosts := []entity.TravelCost{
		{
			Name:  "Local Transportation",
			Code:  "TR-001",
			Unit:  "km",
			Price: 5000,
		},
		{
			Name:  "Hotel Accommodation",
			Code:  "TR-002",
			Unit:  "night",
			Price: 500000,
		},
		{
			Name:  "Daily Allowance",
			Code:  "TR-003",
			Unit:  "day",
			Price: 150000,
		},
		{
			Name:  "Airport Transfer",
			Code:  "TR-004",
			Unit:  "trip",
			Price: 200000,
		},
		{
			Name:  "Meal Allowance",
			Code:  "TR-005",
			Unit:  "day",
			Price: 100000,
		},
	}

	for _, travelCost := range travelCosts {
		if err := db.Create(&travelCost).Error; err != nil {
			log.Printf("Error seeding travel cost data: %v", err)
			return err
		}
	}

	log.Println("Travel costs data seeded successfully!")
	return nil
}
