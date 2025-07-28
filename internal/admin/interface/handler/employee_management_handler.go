package adminHandler

import (
	"net/http"
	"strconv"

	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/dto"
	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/middleware"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeManagementHandler struct {
    usecase adminUsecase.EmployeeManagementUsecase
}

func NewEmployeeManagementHandler(r *gin.Engine, uc adminUsecase.EmployeeManagementUsecase, db *gorm.DB) {
    h := &EmployeeManagementHandler{uc}
    admin := r.Group("/api/v1/admin")
    admin.Use(middleware.AdminRoleMiddleware(db))
    admin.POST("/employees", h.CreateEmployee)
    admin.GET("/employees", h.GetEmployees)
    admin.GET("/employees/add", h.GetAddEmployeePage) // New route for add page
    admin.GET("/employees/:id", h.GetEmployeeByID)
    admin.GET("/employees/:id/edit", h.GetEditEmployeePage)
    admin.PUT("/employees/:id", h.UpdateEmployee)
    admin.DELETE("/employees/:id", h.DeleteEmployee)
}

func (h *EmployeeManagementHandler) CreateEmployee(c *gin.Context) {
    var input dto.CreateEmployeeInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    employee := entity.Employee{
        UserID:            input.UserID,
        JobPositionID:     input.JobPositionID,
        DivisionID:        input.DivisionID,
        CityID:            input.CityID,
        NIP:               input.NIP,
        NIK:               input.NIK,
        BPJSEmploymentNo:  input.BPJSEmploymentNo,
        BPJSHealthNo:      input.BPJSHealthNo,
        Address:           input.Address,
        Phone:             input.Phone,
        JoinDate:          input.JoinDate,
        KTPStatus:         input.KTPStatus,
        ContractNo:        input.ContractNo,
        NPWPStatus:        input.NPWPStatus,
        ContractStatus:    input.ContractStatus,
        Status:            input.Status,
    }

    if err := h.usecase.CreateEmployee(&employee); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "employee created successfully", "data": employee})
}

func (h *EmployeeManagementHandler) GetEmployees(c *gin.Context) {
    page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
    if err != nil || page < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
        return
    }

    limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
    if err != nil || limit < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit parameter"})
        return
    }

    search := c.DefaultQuery("search", "")

    // Get total count of employees for pagination
    total, err := h.usecase.GetEmployeesCount(search)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Calculate total pages
    totalPages := int(total) / limit
    if int(total)%limit > 0 {
        totalPages++
    }

    employees, err := h.usecase.GetEmployees(page, limit, search)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var responseData []dto.EmployeeResponse
    for _, employee := range employees {
        responseData = append(responseData, dto.EmployeeResponse{
            ID:                employee.ID,
            UserID:            employee.UserID,
			Name:              employee.User.Name, // Assuming User entity has a Name field
            JobPosition:       employee.JobPosition.Name,
            Division:          employee.Division.Name,
            City:              employee.City.Name,
            NIP:               employee.NIP,
            NIK:               employee.NIK,
            BPJSEmploymentNo:  employee.BPJSEmploymentNo,
            BPJSHealthNo:      employee.BPJSHealthNo,
            Address:           employee.Address,
            Phone:             employee.Phone,
            JoinDate:          employee.JoinDate,
            KTPStatus:         employee.KTPStatus,
            ContractNo:        employee.ContractNo,
            NPWPStatus:        employee.NPWPStatus,
            CreatedAt:         employee.CreatedAt.Format("2006-01-02 15:04:05"),
            UpdatedAt:         employee.UpdatedAt.Format("2006-01-02 15:04:05"),
        })
    }

    response := gin.H{
        "data": responseData,
        "pagination": gin.H{
            "page":       page,
            "limit":      limit,
            "total_data":      total,
            "total_pages": totalPages,
        },
    }

    c.JSON(http.StatusOK, response)
}

func (h *EmployeeManagementHandler) GetEmployeeByID(c *gin.Context) {
    id := c.Param("id")
    employee, err := h.usecase.GetEmployeeByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    response := dto.EmployeeResponse{
        ID:                employee.ID,
        UserID:            employee.UserID,
		Name:              employee.User.Name, // Assuming User entity has a Name field
        JobPosition:       employee.JobPosition.Name,
        Division:          employee.Division.Name,
        City:              employee.City.Name,
        NIP:               employee.NIP,
        NIK:               employee.NIK,
        BPJSEmploymentNo:  employee.BPJSEmploymentNo,
        BPJSHealthNo:      employee.BPJSHealthNo,
        Address:           employee.Address,
        Phone:             employee.Phone,
        JoinDate:          employee.JoinDate,
        KTPStatus:         employee.KTPStatus,
        ContractNo:        employee.ContractNo,
        NPWPStatus:        employee.NPWPStatus,
        CreatedAt:         employee.CreatedAt.Format("2006-01-02 15:04:05"),
        UpdatedAt:         employee.UpdatedAt.Format("2006-01-02 15:04:05"),
    }

    c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *EmployeeManagementHandler) UpdateEmployee(c *gin.Context) {
    id := c.Param("id")
    idUint, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid employee ID"})
        return
    }

    var input dto.UpdateEmployeeInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    employee := entity.Employee{
        ID:               uint(idUint),
        JobPositionID:    input.JobPositionID,
        DivisionID:       input.DivisionID,
        CityID:           input.CityID,
        NIP:              input.NIP,
        NIK:              input.NIK,
        BPJSEmploymentNo: input.BPJSEmploymentNo,
        BPJSHealthNo:     input.BPJSHealthNo,
        Address:          input.Address,
        Phone:            input.Phone,
        JoinDate:         input.JoinDate,
        KTPStatus:        input.KTPStatus,
        ContractNo:       input.ContractNo,
        NPWPStatus:       input.NPWPStatus,
    }

    if err := h.usecase.UpdateEmployee(id, &employee); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "employee updated successfully", "data": employee})
}

func (h *EmployeeManagementHandler) DeleteEmployee(c *gin.Context) {
    id := c.Param("id")
    if err := h.usecase.DeleteEmployee(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "employee deleted successfully", "data": gin.H{"id": id}})
}

func (h *EmployeeManagementHandler) GetEditEmployeePage(c *gin.Context) {
    id := c.Param("id")
    employee, err := h.usecase.GetEmployeeByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    // Fetch job positions for reference
    jobPositions, err := h.usecase.GetJobPositions(1, 1000, "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    var jobPositionList []dto.JobPositionResponse
    for _, jp := range jobPositions {
        jobPositionList = append(jobPositionList, dto.JobPositionResponse{
            ID:   jp.ID,
            Name: jp.Name,
            CreatedAt: jp.CreatedAt.Format("02-01-2006 15:04:05"),
            UpdatedAt: jp.UpdatedAt.Format("02-01-2006 15:04:05"),
        })
    }
    
    // Fetch divisions for reference
    divisions, err := h.usecase.GetDivisions(1, 1000, "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    var divisionList []dto.DivisionResponse
    for _, division := range divisions {
        divisionList = append(divisionList, dto.DivisionResponse{
            ID:   division.ID,
            Name: division.Name,
            CreatedAt: division.CreatedAt.Format("02-01-2006 15:04:05"),
            UpdatedAt: division.UpdatedAt.Format("02-01-2006 15:04:05"),
        })
    }
    
    // Fetch cities for reference
    cities, err := h.usecase.GetCities(1, 1000, "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    var cityList []dto.CityResponse
    for _, city := range cities {
        cityList = append(cityList, dto.CityResponse{
            ID:       city.ID,
            Name:     city.Name,
            Province: city.Province.Name,
        })
    }
    
    // Fetch provinces with cities for reference
    provinces, err := h.usecase.GetProvinces(1, 1000, "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    var provinceList []dto.ProvinceResponseWithCity
    for _, province := range provinces {
        var citiesInProvince []dto.CityWithoutProvinceResponse
        for _, city := range province.Cities {
            citiesInProvince = append(citiesInProvince, dto.CityWithoutProvinceResponse{
                ID:   city.ID,
                Name: city.Name,
            })
        }
        
        provinceList = append(provinceList, dto.ProvinceResponseWithCity{
            ID:     province.ID,
            Name:   province.Name,
            Cities: citiesInProvince,
        })
    }
    
    response := gin.H{
        "data": employee,
        "references": gin.H{
            "job_positions": jobPositionList,
            "divisions":     divisionList,
            "cities":        cityList,
            "provinces":     provinceList,
        },
    }
    c.JSON(http.StatusOK, response)
}

// GetAddEmployeePage returns job positions, divisions, and provinces with cities for the add employee page
func (h *EmployeeManagementHandler) GetAddEmployeePage(c *gin.Context) {
    // Fetch job positions for reference
    jobPositions, err := h.usecase.GetJobPositions(1, 1000, "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    var jobPositionList []dto.JobPositionResponse
    for _, jp := range jobPositions {
        jobPositionList = append(jobPositionList, dto.JobPositionResponse{
            ID:   jp.ID,
            Name: jp.Name,
            CreatedAt: jp.CreatedAt.Format("02-01-2006 15:04:05"),
            UpdatedAt: jp.UpdatedAt.Format("02-01-2006 15:04:05"),
        })
    }
    
    // Fetch divisions for reference
    divisions, err := h.usecase.GetDivisions(1, 1000, "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    var divisionList []dto.DivisionResponse
    for _, division := range divisions {
        divisionList = append(divisionList, dto.DivisionResponse{
            ID:   division.ID,
            Name: division.Name,
            CreatedAt: division.CreatedAt.Format("02-01-2006 15:04:05"),
            UpdatedAt: division.UpdatedAt.Format("02-01-2006 15:04:05"),
        })
    }
    
    // Fetch provinces with cities for reference
    provinces, err := h.usecase.GetProvinces(1, 1000, "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    var provinceList []dto.ProvinceResponseWithCity
    for _, province := range provinces {
        var citiesInProvince []dto.CityWithoutProvinceResponse
        for _, city := range province.Cities {
            citiesInProvince = append(citiesInProvince, dto.CityWithoutProvinceResponse{
                ID:   city.ID,
                Name: city.Name,
            })
        }
        
        provinceList = append(provinceList, dto.ProvinceResponseWithCity{
            ID:     province.ID,
            Name:   province.Name,
            Cities: citiesInProvince,
        })
    }
    
    response := gin.H{
        "references": gin.H{
            "job_positions": jobPositionList,
            "divisions":     divisionList,
            "provinces":     provinceList,
        },
    }
    c.JSON(http.StatusOK, response)
}
