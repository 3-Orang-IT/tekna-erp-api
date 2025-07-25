package dto

type CreateSupplierInput struct {
	UserID        uint   `json:"user_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	InvoiceName   string `json:"invoice_name" binding:"required"`
	NPWP          string `json:"npwp" binding:"required"`
	Address       string `json:"address" binding:"required"`
	CityID        uint   `json:"city_id" binding:"required"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Greeting      string `json:"greeting"`
	ContactPerson string `json:"contact_person"`
	ContactPhone  string `json:"contact_phone"`
	BankAccount   string `json:"bank_account"`
	Type          string `json:"type"`
	LogoFilename  string `json:"logo_filename"`
}

type UpdateSupplierInput struct {
	UserID        uint   `json:"user_id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	InvoiceName   string `json:"invoice_name"`
	NPWP          string `json:"npwp"`
	Address       string `json:"address"`
	CityID        uint   `json:"city_id"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Greeting      string `json:"greeting"`
	ContactPerson string `json:"contact_person"`
	ContactPhone  string `json:"contact_phone"`
	BankAccount   string `json:"bank_account"`
	Type          string `json:"type"`
	LogoFilename  string `json:"logo_filename"`
}

type SupplierResponse struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	InvoiceName   string `json:"invoice_name"`
	NPWP          string `json:"npwp"`
	Address       string `json:"address"`
	City          string `json:"city"`
	Province      string `json:"province"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Greeting      string `json:"greeting"`
	ContactPerson string `json:"contact_person"`
	ContactPhone  string `json:"contact_phone"`
	BankAccount   string `json:"bank_account"`
	Type          string `json:"type"`
	LogoFilename  string `json:"logo_filename"`
	UpdatedAt     string `json:"updated_at"`
}
