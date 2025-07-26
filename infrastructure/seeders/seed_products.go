package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) error {
	// First get a ProductCategory ID, Supplier ID, BusinessUnit ID, and UnitOfMeasure ID for reference
	var productCategory entity.ProductCategory
	var supplier entity.Supplier
	var businessUnit entity.BusinessUnit
	var unit entity.UnitOfMeasure

	if err := db.Where("name = ?", "Electronics").First(&productCategory).Error; err != nil {
		log.Printf("Error finding ProductCategory 'Electronics': %v", err)
		return err
	}

	if err := db.Where("name = ?", "PT Supplier Utama").First(&supplier).Error; err != nil {
		log.Printf("Error finding Supplier 'PT Supplier Utama': %v", err)
		return err
	}

	if err := db.Where("name = ?", "Operations").First(&businessUnit).Error; err != nil {
		log.Printf("Error finding BusinessUnit 'Operations': %v", err)
		return err
	}

	if err := db.Where("name = ?", "Piece").First(&unit).Error; err != nil {
		log.Printf("Error finding UnitOfMeasure 'Piece': %v", err)
		return err
	}

	productsList := []entity.Product{
		{
			ProductCategoryID: productCategory.ID,
			SupplierID:        supplier.ID,
			BusinessUnitID:    businessUnit.ID,
			UnitID:            unit.ID,
			Code:              "PRD-1",
			Barcode:           "1234567890123",
			Name:              "Smartphone X",
			Description:       "Latest model of Smartphone X",
			CatalogNumber:     "CAT12345",
			ImageFilename:     "smartphone_x.jpg",
			MaxQuantity:       100,
			MinQuantity:       10,
			PurchasePrice:     500.00,
			HPPTaxed:          550.00,
			SellingPrice:      600.00,
			LastPrice:         580.00,
			Packaging:         "Box",
			BrochureLink:      "http://example.com/brochure_smartphone_x",
			IsRecommended:     true,
			ProductType:       "Electronics",
			ProductFocus:      "Consumer",
			Brand:             "BrandX",
		},
	}

	for _, product := range productsList {
		var p entity.Product
		if err := db.Where("code = ?", product.Code).First(&p).Error; err == gorm.ErrRecordNotFound {
			if err := db.Create(&product).Error; err != nil {
				log.Printf("Failed to create product %s: %v", product.Name, err)
				return err
			}
			log.Printf("Successfully added product %s", product.Name)
		}
	}
	return nil
}
