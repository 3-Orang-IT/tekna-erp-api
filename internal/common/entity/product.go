package entity

type Product struct {
	ID                uint    `gorm:"primaryKey"`
	ProductCategoryID uint    `gorm:"not null"`
	SupplierID        uint    `gorm:"not null"`
	BusinessUnitID    uint    `gorm:"not null"`
	UnitID            uint    `gorm:"not null"`
	Code              string  `gorm:"size:50;not null"`
	Barcode           string  `gorm:"size:100"`
	Name              string  `gorm:"size:255;not null"`
	NameBackup        string  `gorm:"size:255"`
	Description       string  `gorm:"size:500"`
	CatalogNumber     string  `gorm:"size:100"`
	ImageFilename     string  `gorm:"size:255"`
	MaxQuantity       int     `gorm:"not null"`
	MinQuantity       int     `gorm:"not null"`
	PurchasePrice     float64 `gorm:"not null"`
	HPPTaxed          float64 `gorm:"not null"`
	SellingPrice      float64 `gorm:"not null"`
	LastPrice         float64 `gorm:"not null"`
	Packaging         string  `gorm:"size:100"`
	BrochureLink      string  `gorm:"size:255"`
	IsRecommended     bool    `gorm:"not null"`
	ProductType       string  `gorm:"size:100"`
	ProductFocus      string  `gorm:"size:100"`
	Brand             string  `gorm:"size:100"`

	ProductCategory ProductCategory `gorm:"foreignKey:ProductCategoryID;constraint:OnDelete:SET NULL;"`
	Supplier        Supplier        `gorm:"foreignKey:SupplierID;constraint:OnDelete:SET NULL;"`
	BusinessUnit    BusinessUnit    `gorm:"foreignKey:BusinessUnitID;constraint:OnDelete:SET NULL;"`
	Unit            UnitOfMeasure   `gorm:"foreignKey:UnitID;constraint:OnDelete:SET NULL;"`
}
