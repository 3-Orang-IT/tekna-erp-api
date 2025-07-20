package dto

import "time"

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


type UserResponse struct {
	ID               uint      `json:"id"`
	Username         string    `json:"username"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	Telp             string    `json:"telp"`
	PhotoProfileURL  string    `json:"photo_profile_url"`
	Status           string    `json:"status"`
	Roles            []string  `json:"roles"` // << hanya nama-nama role
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type UpdateUserInput struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Name            string `json:"name"`
	Telp            string `json:"telp"`
	PhotoProfileURL string `json:"photo_profile_url"`
	Status          string `json:"status"`
	RoleIDs         []uint `json:"roles"`  // roles sebagai array of ID
}
