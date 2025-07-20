package dto

type CreateUserInput struct {
	Username        string   `json:"username" binding:"required"`
	Password        string   `json:"password" binding:"required"`
	Name            string   `json:"name" binding:"required"`
	Email           string   `json:"email" binding:"required,email"`
	Telp            string   `json:"telp"`
	PhotoProfileURL string   `json:"photo_profile_url"`
	Status          string   `json:"status"`
	RoleIDs         []uint   `json:"roles" binding:"required"`
}