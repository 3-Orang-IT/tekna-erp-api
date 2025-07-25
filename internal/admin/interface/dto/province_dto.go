package dto

type CreateProvinceInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateProvinceInput struct {
	Name string `json:"name"`
}

type ProvinceResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
