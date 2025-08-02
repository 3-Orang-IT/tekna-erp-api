package dto

type CreateNewsletterInput struct {
	Type        string `json:"type" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	File        string `json:"file"`
	ValidFrom   string `json:"valid_from" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type UpdateNewsletterInput struct {
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	File        string `json:"file"`
	ValidFrom   string `json:"valid_from"`
	Status      string `json:"status"`
}

type NewsletterResponse struct {
	ID          uint   `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	File        string `json:"file"`
	ValidFrom   string `json:"valid_from"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
