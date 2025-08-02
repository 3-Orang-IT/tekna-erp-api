package dto

type CreateAreaInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAreaInput struct {
	Name string `json:"name" binding:"required"`
}

type AreaResponse struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Employees []string `json:"employees"` // Assuming employees are represented by their names or IDs
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}
