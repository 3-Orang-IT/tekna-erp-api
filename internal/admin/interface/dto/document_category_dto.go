package dto

type CreateDocumentCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateDocumentCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

type DocumentCategoryResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
