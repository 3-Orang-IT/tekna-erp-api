package dto

type CreateProductCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateProductCategoryInput struct {
	Name string `json:"name"`
}

type ProductCategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
