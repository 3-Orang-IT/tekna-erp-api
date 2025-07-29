package dto

type CreateBudgetCategoryInput struct {
	ChartOfAccountID *uint  `json:"chart_of_account_id"`
	Name             string `json:"name" binding:"required"`
	Description      string `json:"description"`
	Order            int    `json:"order" binding:"required"`
}

type UpdateBudgetCategoryInput struct {
	ChartOfAccountID *uint  `json:"chart_of_account_id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Order            int    `json:"order"`
}

type BudgetCategoryResponse struct {
	ID             uint   `json:"id"`
	ChartOfAccount string `json:"chart_of_account"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Order          int    `json:"order"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
