package dto

type CreateMenuInput struct {
	Name     string `json:"name" binding:"required"`
	URL      string `json:"url" binding:"required"`
	Icon     string `json:"icon"`
	Order    int    `json:"order" binding:"required"`
	ParentID *uint  `json:"parent_id"`
}

type UpdateMenuInput struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Icon     string `json:"icon"`
	Order    int    `json:"order"`
	ParentID *uint  `json:"parent_id"`
}
