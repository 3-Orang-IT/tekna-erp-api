package dto

type CreateProductCategoryAlternativeInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateProductCategoryAlternativeInput struct {
	Name string `json:"name"`
}

type ProductCategoryAlternativeResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
