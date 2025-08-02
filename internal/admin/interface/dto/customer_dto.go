package dto

type CreateCustomerInput struct {
	UserID            uint   `json:"user_id" binding:"required"`
	AreaID            uint   `json:"area_id" binding:"required"`
	CityID            uint   `json:"city_id" binding:"required"`
	Name              string `json:"name" binding:"required"`
	InvoiceName       string `json:"invoice_name" binding:"required"`
	Address           string `json:"address" binding:"required"`
	Phone             string `json:"phone"`
	Email             string `json:"email"`
	Tax               string `json:"tax"`
	Greeting          string `json:"greeting"`
	ContactPersonName string `json:"contact_person_name"`
	ContactPhone      string `json:"contact_phone"`
	Segment           string `json:"segment"`
	Type              string `json:"type"`
	NPWP              string `json:"npwp"`
	Status            string `json:"status"`
	BEName            string `json:"be_name"`
	ProcurementType   string `json:"procurement_type"`
	MarketingName     string `json:"marketing_name"`
	Note              string `json:"note"`
	PaymentTerm       string `json:"payment_term"`
	Level             string `json:"level"`
}

type UpdateCustomerInput struct {
	AreaID            uint   `json:"area_id"`
	CityID            uint   `json:"city_id"`
	Name              string `json:"name"`
	Code              string `json:"code"`
	InvoiceName       string `json:"invoice_name"`
	Address           string `json:"address"`
	Phone             string `json:"phone"`
	Email             string `json:"email"`
	Tax               string `json:"tax"`
	Greeting          string `json:"greeting"`
	ContactPersonName string `json:"contact_person_name"`
	ContactPhone      string `json:"contact_phone"`
	Segment           string `json:"segment"`
	Type              string `json:"type"`
	NPWP              string `json:"npwp"`
	Status            string `json:"status"`
	BEName            string `json:"be_name"`
	ProcurementType   string `json:"procurement_type"`
	MarketingName     string `json:"marketing_name"`
	Note              string `json:"note"`
	PaymentTerm       string `json:"payment_term"`
	Level             string `json:"level"`
}

type CustomerResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	InvoiceName string `json:"invoice_name"`
	Code        string `json:"code"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Segment     string `json:"segment"`
	Area        string `json:"area"`
	Type        string `json:"type"`
	Level       string `json:"level"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
