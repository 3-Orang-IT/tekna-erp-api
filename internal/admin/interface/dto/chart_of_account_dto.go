package dto

type CreateChartOfAccountInput struct {
	Type string `json:"type" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type UpdateChartOfAccountInput struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type ChartOfAccountResponse struct {
	ID        uint   `json:"id"`
	Type      string `json:"type"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
