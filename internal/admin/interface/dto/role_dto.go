package dto

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type CreateRoleInput struct {
	Name    string `json:"name" binding:"required"`
	Code    string `json:"code" binding:"required"`
	MenuIDs []uint `json:"menu_ids" binding:"required"`
}

type UpdateRoleInput struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	MenuIDs []uint `json:"menu_ids"`
}

type RoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Menus []entity.Menu `json:"menus"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

type RoleResponseOnlyName struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}