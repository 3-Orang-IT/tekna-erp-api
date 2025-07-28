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
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}

type ProvinceResponseWithCity struct {
	ID      uint            `json:"id"`
	Name    string          `json:"name"`
	Cities  []CityWithoutProvinceResponse `json:"cities"`
}
