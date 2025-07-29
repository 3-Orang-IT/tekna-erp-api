package dto

type CreateDocumentInput struct {
	DocumentCategoryID uint   `form:"document_category_id" binding:"required"`
	Name               string `form:"name" binding:"required"`
	Description        string `form:"description"`
	IsPublished        bool   `form:"is_published"`
}

type UpdateDocumentInput struct {
	DocumentCategoryID uint   `form:"document_category_id"`
	Name               string `form:"name"`
	Description        string `form:"description"`
	IsPublished        bool   `form:"is_published"`
}

type DocumentResponse struct {
	ID                 uint                            `json:"id"`
	DocumentCategoryID uint                            `json:"document_category_id"`
	DocumentCategory   DocumentCategoryResponseLimited `json:"document_category"`
	Name               string                          `json:"name"`
	UserID             uint                            `json:"user_id"`
	User               UserResponseLimited             `json:"user"`
	FilePath           string                          `json:"file_path"`
	Description        string                          `json:"description"`
	IsPublished        bool                            `json:"is_published"`
	CreatedAt          string                          `json:"created_at"`
	UpdatedAt          string                          `json:"updated_at"`
}

type DocumentCategoryResponseLimited struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserResponseLimited struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
