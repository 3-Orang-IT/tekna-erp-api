package dto

type CreateCompanyInput struct {
	Name             string  `json:"name" binding:"required"`
	Address          string  `json:"address" binding:"required"`
	CityID           uint    `json:"city_id" binding:"required"`
	ProvinceID       uint    `json:"province_id" binding:"required"`
	Phone            string  `json:"phone"`
	Fax              string  `json:"fax"`
	Email            string  `json:"email" binding:"required"`
	StartHour        string  `json:"start_hour" binding:"required"`
	EndHour          string  `json:"end_hour" binding:"required"`
	Latitude         float64 `json:"latitude" binding:"required"`
	Longitude        float64 `json:"longitude" binding:"required"`
	TotalShares      int     `json:"total_shares" binding:"required"`
	AnnualLeaveQuota int     `json:"annual_leave_quota" binding:"required"`
}

type UpdateCompanyInput struct {
	Name             string  `json:"name"`
	Address          string  `json:"address"`
	CityID           uint    `json:"city_id"`
	ProvinceID       uint    `json:"province_id"`
	Phone            string  `json:"phone"`
	Fax              string  `json:"fax"`
	Email            string  `json:"email"`
	StartHour        string  `json:"start_hour"`
	EndHour          string  `json:"end_hour"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	TotalShares      int     `json:"total_shares"`
	AnnualLeaveQuota int     `json:"annual_leave_quota"`
}
