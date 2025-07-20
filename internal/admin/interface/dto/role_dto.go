package dto

type CreateRoleInput struct {
	Name    string `json:"name" binding:"required"`
	MenuIDs []uint `json:"menu_ids" binding:"required"`
}

type UpdateRoleInput struct {
	Name    string `json:"name"`
	MenuIDs []uint `json:"menu_ids"`
}
