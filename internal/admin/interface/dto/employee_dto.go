package dto

type CreateEmployeeInput struct {
	UserID           uint   `json:"user_id" binding:"required"`
	JobPositionID    uint   `json:"job_position_id" binding:"required"`
	DivisionID       uint   `json:"division_id" binding:"required"`
	CityID           uint   `json:"city_id" binding:"required"`
	NIP              string `json:"nip"`
	NIK              string `json:"nik"`
	BPJSEmploymentNo string `json:"bpjs_employment_no"`
	BPJSHealthNo     string `json:"bpjs_health_no"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	JoinDate         string `json:"join_date"`
	KTPStatus        string `json:"ktp_status"`
	ContractNo       string `json:"contract_no"`
	NPWPStatus       string `json:"npwp_status"`
	ContractStatus   string `json:"contract_status"`
	Status           string `json:"status"`
	AreaIDs          []uint `json:"area_ids"` // For many2many relation
}

type UpdateEmployeeInput struct {
	JobPositionID    uint   `json:"job_position_id"`
	DivisionID       uint   `json:"division_id"`
	CityID           uint   `json:"city_id"`
	NIP              string `json:"nip"`
	NIK              string `json:"nik"`
	BPJSEmploymentNo string `json:"bpjs_employment_no"`
	BPJSHealthNo     string `json:"bpjs_health_no"`
	Address          string `json:"address"`
	Phone            string `json:"phone"`
	JoinDate         string `json:"join_date"`
	KTPStatus        string `json:"ktp_status"`
	ContractNo       string `json:"contract_no"`
	NPWPStatus       string `json:"npwp_status"`
	ContractStatus   string `json:"contract_status"`
	Status           string `json:"status"`
	AreaIDs          []uint `json:"area_ids"`
}

type EmployeeResponse struct {
	ID               uint     `json:"id"`
	UserID           uint     `json:"user_id"`
	Name             string   `json:"name"` // Assuming you want to include user's name
	JobPosition      string   `json:"job_position"`
	Division         string   `json:"division"`
	City             string   `json:"city"`
	NIP              string   `json:"nip"`
	NIK              string   `json:"nik"`
	BPJSEmploymentNo string   `json:"bpjs_employment_no"`
	BPJSHealthNo     string   `json:"bpjs_health_no"`
	Address          string   `json:"address"`
	Phone            string   `json:"phone"`
	JoinDate         string   `json:"join_date"`
	KTPStatus        string   `json:"ktp_status"`
	ContractNo       string   `json:"contract_no"`
	NPWPStatus       string   `json:"npwp_status"`
	ContractStatus   string   `json:"contract_status"`
	Status           string   `json:"status"`
	Area             []string `json:"area"` // Area names
	UpdatedAt        string   `json:"updated_at"`
}
