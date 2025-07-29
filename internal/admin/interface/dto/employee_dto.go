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

// CreateEmployeeWithUserInput combines user and employee data for creating both at once
type CreateEmployeeWithUserInput struct {
	// User data
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Telp     string `form:"telp"`
	RoleIDs  []uint `form:"roles" binding:"required"`

	// Employee data
	JobPositionID    uint   `form:"job_position_id" binding:"required"`
	DivisionID       uint   `form:"division_id" binding:"required"`
	CityID           uint   `form:"city_id" binding:"required"`
	NIP              string `form:"nip"`
	NIK              string `form:"nik"`
	BPJSEmploymentNo string `form:"bpjs_employment_no"`
	BPJSHealthNo     string `form:"bpjs_health_no"`
	Address          string `form:"address"`
	Phone            string `form:"phone"`
	JoinDate         string `form:"join_date"`
	KTPStatus        string `form:"ktp_status"`
	ContractNo       string `form:"contract_no"`
	NPWPStatus       string `form:"npwp_status"`
	ContractStatus   string `form:"contract_status"`
	Status           string `form:"status"`
	AreaIDs          []uint `form:"area_ids"` // For many2many relation
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
	Name             string   `json:"name"`              // Assuming you want to include user's name
	PhotoProfileURL  string   `json:"photo_profile_url"` // URL to the employee's photo profile
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
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
}
