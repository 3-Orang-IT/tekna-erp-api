package dto

type CreateProductInput struct {
	ProductCategoryID uint    `json:"product_category_id" binding:"required"`
	SupplierID        uint    `json:"supplier_id" binding:"required"`
	BusinessUnitID    uint    `json:"business_unit_id" binding:"required"`
	UnitID            uint    `json:"unit_id" binding:"required"`
	Barcode           string  `json:"barcode"`
	Name              string  `json:"name" binding:"required"`
	Description       string  `json:"description"`
	MaxQuantity       int     `json:"max_quantity" binding:"required"`
	MinQuantity       int     `json:"min_quantity" binding:"required"`
	PurchasePrice     float64 `json:"purchase_price" binding:"required"`
	SellingPrice      float64 `json:"selling_price" binding:"required"`
	IsRecommended     bool    `json:"is_recommended"`
	ProductFocus      string  `json:"product_focus"`
	Brand             string  `json:"brand"`
}

type UpdateProductInput struct {
	ProductCategoryID uint    `json:"product_category_id"`
	SupplierID        uint    `json:"supplier_id"`
	BusinessUnitID    uint    `json:"business_unit_id"`
	UnitID            uint    `json:"unit_id"`
	Barcode           string  `json:"barcode"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	MaxQuantity       int     `json:"max_quantity"`
	MinQuantity       int     `json:"min_quantity"`
	PurchasePrice     float64 `json:"purchase_price"`
	SellingPrice      float64 `json:"selling_price"`
	IsRecommended     bool    `json:"is_recommended"`
	ProductType       string  `json:"product_type"`
	ProductFocus      string  `json:"product_focus"`
	Brand             string  `json:"brand"`
}

type ProductResponse struct {
	ID                  uint    `json:"id"`
	Code                string  `json:"code"`
	Barcode             string  `json:"barcode"`
	Name                string  `json:"name"`
	ProductCategoryName string  `json:"product_category_name"`
	SupplierName        string  `json:"supplier_name"`
	BusinessUnitName    string  `json:"business_unit_name"`
	UnitName            string  `json:"unit_name"`
	Description         string  `json:"description"`
	MaxQuantity         int     `json:"max_quantity"`
	MinQuantity         int     `json:"min_quantity"`
	PurchasePrice       float64 `json:"purchase_price"`
	SellingPrice        float64 `json:"selling_price"`
	IsRecommended       bool    `json:"is_recommended"`
	ProductType         string  `json:"product_type"`
	ProductFocus        string  `json:"product_focus"`
	Brand               string  `json:"brand"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}
