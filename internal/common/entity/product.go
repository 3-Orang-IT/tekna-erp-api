package entity

import "time"

type Product struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	ProductCategoryID uint      `gorm:"not null" json:"product_category_id"`
	SupplierID        uint      `gorm:"not null" json:"supplier_id"`
	BusinessUnitID    uint      `gorm:"not null" json:"business_unit_id"`
	UnitID            uint      `gorm:"not null" json:"unit_id"`
	Code              string    `gorm:"size:50;not null" json:"code" `
	Barcode           string    `gorm:"size:100" json:"barcode"`
	Name              string    `gorm:"size:255;not null" json:"name"`
	NameBackup        string    `gorm:"size:255" json:"name_backup"`
	Description       string    `gorm:"size:500" json:"description"`
	CatalogNumber     string    `gorm:"size:100" json:"catalog_number"`
	ImageFilename     string    `gorm:"size:255" json:"image_filename"`
	MaxQuantity       int       `gorm:"not null" json:"max_quantity"`
	MinQuantity       int       `gorm:"not null" json:"min_quantity"`
	PurchasePrice     float64   `gorm:"not null" json:"purchase_price"`
	HPPTaxed          float64   `gorm:"not null" json:"hpp_taxed"`
	SellingPrice      float64   `gorm:"not null" json:"selling_price"`
	LastPrice         float64   `gorm:"not null" json:"last_price"`
	Packaging         string    `gorm:"size:100" json:"packaging"`
	BrochureLink      string    `gorm:"size:255" json:"brochure_link"`
	IsRecommended     bool      `gorm:"not null" json:"is_recommended"`
	ProductType       string    `gorm:"size:100" json:"product_type"`
	ProductFocus      string    `gorm:"size:100" json:"product_focus"`
	Brand             string    `gorm:"size:100" json:"brand"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`

	ProductCategory ProductCategory `gorm:"foreignKey:ProductCategoryID;constraint:OnDelete:SET NULL;"`
	Supplier        Supplier        `gorm:"foreignKey:SupplierID;constraint:OnDelete:SET NULL;"`
	BusinessUnit    BusinessUnit    `gorm:"foreignKey:BusinessUnitID;constraint:OnDelete:SET NULL;"`
	Unit            UnitOfMeasure   `gorm:"foreignKey:UnitID;constraint:OnDelete:SET NULL;"`
}
