package seeders

import (
	"log"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	db.Transaction(func(tx *gorm.DB) error {
		if err := SeedProvinces(tx); err != nil {
			log.Printf("Error seeding provinces: %v", err)
			return err
		}
		
		if err := SeedCities(tx); err != nil {
			log.Printf("Error seeding cities: %v", err)
			return err
		}
		
	   if err := SeedCompanies(tx); err != nil {
			   log.Printf("Error seeding companies: %v", err)
			   return err
	   }

	   if err := SeedJobPositions(tx); err != nil {
			   log.Printf("Error seeding job positions: %v", err)
			   return err
	   }
		
		if err := SeedModuls(tx); err != nil {
			log.Printf("Error seeding moduls: %v", err)
			return err
		}
		
		if err := SeedMenus(tx); err != nil {
			log.Printf("Error seeding menus: %v", err)
			return err
		}

		if err := SeedRoles(tx); err != nil {
			log.Printf("Error seeding roles: %v", err)
			return err
		}

		if err := SeedDivisions(tx); err != nil {
			log.Printf("Error seeding divisions: %v", err)
			return err
		}

		if err := SeedUsers(tx); err != nil {
			log.Printf("Error seeding users: %v", err)
			return err
		}

		if err := SeedProductCategories(tx); err != nil {
			log.Printf("Error seeding product categories: %v", err)
			return err
		}

		return nil
	})
}