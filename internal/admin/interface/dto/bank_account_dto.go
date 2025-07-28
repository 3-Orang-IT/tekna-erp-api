package dto

type CreateBankAccountInput struct {
	ChartOfAccountID uint   `json:"chart_of_account_id" binding:"required"`
	AccountNumber    string `json:"account_number" binding:"required"`
	BankName         string `json:"bank_name" binding:"required"`
	BranchAddress    string `json:"branch_address"`
	CityID           uint   `json:"city_id" binding:"required"`
	PhoneNumber      string `json:"phone_number"`
	Priority         int    `json:"priority" binding:"required"`
}

type UpdateBankAccountInput struct {
	ChartOfAccountID uint   `json:"chart_of_account_id"`
	AccountNumber    string `json:"account_number"`
	BankName         string `json:"bank_name"`
	BranchAddress    string `json:"branch_address"`
	CityID           uint   `json:"city_id"`
	PhoneNumber      string `json:"phone_number"`
	Priority         int    `json:"priority"`
}

type BankAccountResponse struct {
	ID               uint   `json:"id"`
	ChartOfAccount   string `json:"chart_of_account"`
	AccountNumber    string `json:"account_number"`
	BankName         string `json:"bank_name"`
	BranchAddress    string `json:"branch_address"`
	City             string `json:"city"`
	Province         string `json:"province"`
	PhoneNumber      string `json:"phone_number"`
	Priority         int    `json:"priority"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}
