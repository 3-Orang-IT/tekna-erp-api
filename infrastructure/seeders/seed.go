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
		
		if err := SeedSuppliers(tx); err != nil {
			log.Printf("Error seeding suppliers: %v", err)
			return err
		}

		if err := SeedBusinessUnits(tx); err != nil {
			log.Printf("Error seeding business units: %v", err)
			return err
		}

		if err := SeedUnitOfMeasures(tx); err != nil {
			log.Printf("Error seeding unit of measures: %v", err)
			return err
		}

		if err := SeedProducts(tx); err != nil {
			log.Printf("Error seeding products: %v", err)
			return err
		}

		if err := SeedAreas(tx); err != nil {
			log.Printf("Error seeding areas: %v", err)
			return err
		}

		if err := SeedCustomers(tx); err != nil {
			log.Printf("Error seeding customers: %v", err)
			return err
		}

		if err := SeedChartOfAccounts(tx); err != nil {
			log.Printf("Error seeding chart of accounts: %v", err)
			return err
		}

		if err := SeedNewsletters(tx); err != nil {
			log.Printf("Error seeding newsletters: %v", err)
			return err
		}

		if err := SeedDocumentCategories(tx); err != nil {
			log.Printf("Error seeding document categories: %v", err)
			return err
		}

		if err := SeedDocuments(tx); err != nil {
			log.Printf("Error seeding documents: %v", err)
			return err
		}

		if err := SeedBankAccounts(tx); err != nil {
			log.Printf("Error seeding bank accounts: %v", err)
			return err
		}

		if err := SeedToDoTemplates(tx); err != nil {
			log.Printf("Error seeding to-do templates: %v", err)
			return err
		}

		if err := SeedTravelCosts(tx); err != nil {
			log.Printf("Error seeding travel costs: %v", err)
			return err
		}

		return nil
	})
}