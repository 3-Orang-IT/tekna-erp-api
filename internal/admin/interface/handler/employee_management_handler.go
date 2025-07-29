package adminHandler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/dto"
	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/middleware"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ensureDir creates a directory if it does not exist
func ensureDirEmployee(dirName string) error {
    err := os.MkdirAll(dirName, os.ModePerm)
    if err != nil && !os.IsExist(err) {
        return err
    }
    return nil
}

type EmployeeManagementHandler struct {
    usecase adminUsecase.EmployeeManagementUsecase
}

func NewEmployeeManagementHandler(r *gin.Engine, uc adminUsecase.EmployeeManagementUsecase, db *gorm.DB) {
    h := &EmployeeManagementHandler{uc}
    admin := r.Group("/api/v1/admin")
    admin.Use(middleware.AdminRoleMiddleware(db))
    admin.POST("/employees", h.CreateEmployee)                  // Create employee for existing user
    admin.POST("/employees/with-user", h.CreateEmployeeWithUser) // Create employee with new user
    admin.GET("/employees", h.GetEmployees)
    admin.GET("/employees/add", h.GetAddEmployeePage)           // New route for add page
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
    
    // Handle area IDs if provided
    if len(input.AreaIDs) > 0 {
        // Convert area IDs to Area entities
        for _, areaID := range input.AreaIDs {
            employee.Area = append(employee.Area, entity.Area{ID: areaID})
        }
    }

    // Set default values if not provided
    if employee.Status == "" {
        employee.Status = "active"
    }
    if employee.ContractStatus == "" {
        employee.ContractStatus = "active"
    }

    if err := h.usecase.CreateEmployee(&employee); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Fetch the created employee with its relations
    createdEmployee, err := h.usecase.GetEmployeeByID(strconv.FormatUint(uint64(employee.ID), 10))
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "message": "employee created successfully, but could not fetch details",
            "data": employee,
        })
        return
    }

    // Get base URL for photo profile
    baseUrl := os.Getenv("BASE_URL")
    
    // Format the response with related data
    response := dto.EmployeeResponse{
        ID:               createdEmployee.ID,
        UserID:           createdEmployee.UserID,
        Name:             createdEmployee.User.Name,
        PhotoProfileURL:  baseUrl + createdEmployee.User.PhotoProfileURL,
        JobPosition:      createdEmployee.JobPosition.Name,
        Division:         createdEmployee.Division.Name,
        City:             createdEmployee.City.Name,
        NIP:              createdEmployee.NIP,
        NIK:              createdEmployee.NIK,
        BPJSEmploymentNo: createdEmployee.BPJSEmploymentNo,
        BPJSHealthNo:     createdEmployee.BPJSHealthNo,
        Address:          createdEmployee.Address,
        Phone:            createdEmployee.Phone,
        JoinDate:         createdEmployee.JoinDate,
        KTPStatus:        createdEmployee.KTPStatus,
        ContractNo:       createdEmployee.ContractNo,
        NPWPStatus:       createdEmployee.NPWPStatus,
        ContractStatus:   createdEmployee.ContractStatus,
        Status:           createdEmployee.Status,
        CreatedAt:        createdEmployee.CreatedAt.Format("2006-01-02 15:04:05"),
        UpdatedAt:        createdEmployee.UpdatedAt.Format("2006-01-02 15:04:05"),
    }
    
    // Add area names to the response
    var areaNames []string
    for _, area := range createdEmployee.Area {
        areaNames = append(areaNames, area.Name)
    }
    response.Area = areaNames

    c.JSON(http.StatusCreated, gin.H{
        "message": "employee created successfully", 
        "data": response,
    })
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

    baseUrl := os.Getenv("BASE_URL")

    employees, err := h.usecase.GetEmployees(page, limit, search)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var responseData []dto.EmployeeResponse
    for _, employee := range employees {
        // Extract area names
        var areaNames []string
        for _, area := range employee.Area {
            areaNames = append(areaNames, area.Name)
        }
        
        responseData = append(responseData, dto.EmployeeResponse{
            ID:                employee.ID,
            UserID:            employee.UserID,
            Name:              employee.User.Name,
            PhotoProfileURL:   baseUrl + employee.User.PhotoProfileURL,
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
            ContractStatus:    employee.ContractStatus,
            Status:            employee.Status,
            Area:              areaNames,
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

    // Extract area names
    var areaNames []string
    for _, area := range employee.Area {
        areaNames = append(areaNames, area.Name)
    }
    
    baseUrl := os.Getenv("BASE_URL")
    
    response := dto.EmployeeResponse{
        ID:                employee.ID,
        UserID:            employee.UserID,
        Name:              employee.User.Name,
        PhotoProfileURL:   baseUrl + employee.User.PhotoProfileURL,
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
        ContractStatus:    employee.ContractStatus,
        Status:            employee.Status,
        Area:              areaNames,
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
        ContractStatus:   input.ContractStatus,
        Status:           input.Status,
    }
    
    // Handle area IDs if provided
    if len(input.AreaIDs) > 0 {
        // Convert area IDs to Area entities
        for _, areaID := range input.AreaIDs {
            employee.Area = append(employee.Area, entity.Area{ID: areaID})
        }
    }

    if err := h.usecase.UpdateEmployee(id, &employee); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Fetch the updated employee with its relations
    updatedEmployee, err := h.usecase.GetEmployeeByID(id)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{
            "message": "employee updated successfully, but could not fetch updated details",
            "data": employee,
        })
        return
    }

    // Format the response with related data
    response := dto.EmployeeResponse{
        ID:               updatedEmployee.ID,
        UserID:           updatedEmployee.UserID,
        Name:             updatedEmployee.User.Name,
        JobPosition:      updatedEmployee.JobPosition.Name,
        Division:         updatedEmployee.Division.Name,
        City:             updatedEmployee.City.Name,
        NIP:              updatedEmployee.NIP,
        NIK:              updatedEmployee.NIK,
        BPJSEmploymentNo: updatedEmployee.BPJSEmploymentNo,
        BPJSHealthNo:     updatedEmployee.BPJSHealthNo,
        Address:          updatedEmployee.Address,
        Phone:            updatedEmployee.Phone,
        JoinDate:         updatedEmployee.JoinDate,
        KTPStatus:        updatedEmployee.KTPStatus,
        ContractNo:       updatedEmployee.ContractNo,
        NPWPStatus:       updatedEmployee.NPWPStatus,
        ContractStatus:   updatedEmployee.ContractStatus,
        Status:           updatedEmployee.Status,
        CreatedAt:        updatedEmployee.CreatedAt.Format("2006-01-02 15:04:05"),
        UpdatedAt:        updatedEmployee.UpdatedAt.Format("2006-01-02 15:04:05"),
    }
    
    // Add area names to the response
    var areaNames []string
    for _, area := range updatedEmployee.Area {
        areaNames = append(areaNames, area.Name)
    }
    response.Area = areaNames

    c.JSON(http.StatusOK, gin.H{"message": "employee updated successfully", "data": response})
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
    
    // Fetch areas for reference
    areas, err := h.usecase.GetAreas(1, 1000, "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    var areaList []dto.AreaResponse
    for _, area := range areas {
        areaList = append(areaList, dto.AreaResponse{
            ID:   area.ID,
            Name: area.Name,
            CreatedAt: area.CreatedAt.Format("02-01-2006 15:04:05"),
            UpdatedAt: area.UpdatedAt.Format("02-01-2006 15:04:05"),
        })
    }

    response := gin.H{
        "data": employee,
        "references": gin.H{
            "job_positions": jobPositionList,
            "divisions":     divisionList,
            "cities":        cityList,
            "provinces":     provinceList,
            "areas":         areaList,
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
    
    // Fetch users for reference (for employee without user creation)
    users, err := h.usecase.GetUsers(1, 1000, "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    var userList []dto.UserResponse
    for _, user := range users {
        var roleNames []string
        for _, role := range user.Role {
            roleNames = append(roleNames, role.Name)
        }
        
        userList = append(userList, dto.UserResponse{
            ID:       user.ID,
            Name:     user.Name,
            Username: user.Username,
            Email:    user.Email,
            Roles:    roleNames,
        })
    }

    // Fetch areas for reference
    areas, err := h.usecase.GetAreas(1, 1000, "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    var areaList []dto.AreaResponse
    for _, area := range areas {
        areaList = append(areaList, dto.AreaResponse{
            ID:   area.ID,
            Name: area.Name,
            CreatedAt: area.CreatedAt.Format("02-01-2006 15:04:05"),
            UpdatedAt: area.UpdatedAt.Format("02-01-2006 15:04:05"),
        })
    }

    response := gin.H{
        "references": gin.H{
            "job_positions": jobPositionList,
            "divisions":     divisionList,
            "provinces":     provinceList,
            "users":         userList,
            "areas":         areaList,
        },
    }
    c.JSON(http.StatusOK, response)
}

// CreateEmployeeWithUser creates a new user and assigns them as an employee
func (h *EmployeeManagementHandler) CreateEmployeeWithUser(c *gin.Context) {
    // Parse form-data (multipart)
    if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB max
        c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form: " + err.Error()})
        return
    }

    // Get fields from form-data
    username := c.PostForm("username")
    password := c.PostForm("password")
    name := c.PostForm("name")
    email := c.PostForm("email")
    telp := c.PostForm("telp")
    
    // Convert roleIDs from string array to uint array
    roleIDs := c.PostFormArray("roles")
    var roleIDsUint []uint
    for _, rid := range roleIDs {
        id, err := strconv.ParseUint(rid, 10, 64)
        if err == nil {
            roleIDsUint = append(roleIDsUint, uint(id))
        }
    }

    // Handle file upload
    file, header, err := c.Request.FormFile("photo_profile")
    var photoProfileURL string
    if err == nil && header != nil {
        sanitizedFilename := strings.ReplaceAll(header.Filename, " ", "_")
        filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), sanitizedFilename)
        savePath := fmt.Sprintf("uploads/profile/%s", filename)
        // Ensure directory exists
        if err := ensureDirEmployee("uploads/profile"); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create upload directory: " + err.Error()})
            return
        }
        out, err := os.Create(savePath)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file: " + err.Error()})
            return
        }
        defer out.Close()
        if _, err := io.Copy(out, file); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to write file: " + err.Error()})
            return
        }
        photoProfileURL = savePath
    }

    // Get employee fields from form
    jobPositionID, _ := strconv.ParseUint(c.PostForm("job_position_id"), 10, 64)
    divisionID, _ := strconv.ParseUint(c.PostForm("division_id"), 10, 64)
    cityID, _ := strconv.ParseUint(c.PostForm("city_id"), 10, 64)
    nip := c.PostForm("nip")
    nik := c.PostForm("nik")
    bpjsEmploymentNo := c.PostForm("bpjs_employment_no")
    bpjsHealthNo := c.PostForm("bpjs_health_no")
    address := c.PostForm("address")
    phone := c.PostForm("phone")
    joinDate := c.PostForm("join_date")
    ktpStatus := c.PostForm("ktp_status")
    contractNo := c.PostForm("contract_no")
    npwpStatus := c.PostForm("npwp_status")
    contractStatus := c.PostForm("contract_status")
    status := c.PostForm("status")
    areaIDs := c.PostFormArray("area_ids")

    // Convert areaIDs from string array to uint array
    var areaIDsUint []uint
    for _, rid := range areaIDs {
        id, err := strconv.ParseUint(rid, 10, 64)
        if err == nil {
            areaIDsUint = append(areaIDsUint, uint(id))
        }
    }

    // Validate required fields
    if username == "" || password == "" || name == "" || email == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "username, password, name, and email are required"})
        return
    }
    
    if jobPositionID == 0 || divisionID == 0 || cityID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "job position ID, division ID, and city ID are required"})
        return
    }
    
    // Create user entity
    user := entity.User{
        Username:        username,
        Password:        password,
        Name:            name,
        Email:           email,
        Telp:            telp,
        PhotoProfileURL: photoProfileURL,
        Status:          "active", // Default status for new users
    }
    
    // Create employee entity
    employee := entity.Employee{
        JobPositionID:     uint(jobPositionID),
        DivisionID:        uint(divisionID),
        CityID:            uint(cityID),
        NIP:               nip,
        NIK:               nik,
        BPJSEmploymentNo:  bpjsEmploymentNo,
        BPJSHealthNo:      bpjsHealthNo,
        Address:           address,
        Phone:             phone,
        JoinDate:          joinDate,
        KTPStatus:         ktpStatus,
        ContractNo:        contractNo,
        NPWPStatus:        npwpStatus,
        ContractStatus:    contractStatus,
        Status:            status,
    }
    
    // Add area associations if area IDs were provided
    if len(areaIDsUint) > 0 {
        for _, areaID := range areaIDsUint {
            employee.Area = append(employee.Area, entity.Area{ID: areaID})
        }
    }
    
    // Set default values if not provided
    if employee.Status == "" {
        employee.Status = "active"
    }
    if employee.ContractStatus == "" {
        employee.ContractStatus = "active"
    }
    
    if err := h.usecase.CreateEmployeeWithUser(&user, &employee, roleIDsUint); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    // Format photo URL with base URL if provided
    baseUrl := os.Getenv("BASE_URL")
    var photoURL string
    if user.PhotoProfileURL != "" {
        photoURL = baseUrl + user.PhotoProfileURL
    }
    
    // Fetch the created employee with all its relations for a complete response
    createdEmployee, err := h.usecase.GetEmployeeByID(strconv.FormatUint(uint64(employee.ID), 10))
    if err != nil {
        // If there's an error fetching the complete employee, return a basic response
        response := gin.H{
            "message": "user and employee created successfully, but could not fetch complete details",
            "data": gin.H{
                "user": gin.H{
                    "id":               user.ID,
                    "username":         user.Username,
                    "name":             user.Name,
                    "email":            user.Email,
                    "photo_profile_url": photoURL,
                },
                "employee": gin.H{
                    "id":               employee.ID,
                    "user_id":          employee.UserID,
                    "job_position_id":  employee.JobPositionID,
                    "division_id":      employee.DivisionID,
                    "city_id":          employee.CityID,
                    "nip":              employee.NIP,
                    "nik":              employee.NIK,
                },
            },
        }
        c.JSON(http.StatusCreated, response)
        return
    }
    
    // Extract area names for the response
    var areaNames []string
    for _, area := range createdEmployee.Area {
        areaNames = append(areaNames, area.Name)
    }

    // Create a detailed response with all employee information
    response := gin.H{
        "message": "user and employee created successfully",
        "data": gin.H{
            "user": gin.H{
                "id":               user.ID,
                "username":         user.Username,
                "name":             user.Name,
                "email":            user.Email,
                "photo_profile_url": photoURL,
            },
            "employee": gin.H{
                "id":               createdEmployee.ID,
                "user_id":          createdEmployee.UserID,
                "name":             createdEmployee.User.Name,
                "job_position":     createdEmployee.JobPosition.Name,
                "division":         createdEmployee.Division.Name,
                "city":             createdEmployee.City.Name,
                "nip":              createdEmployee.NIP,
                "nik":              createdEmployee.NIK,
                "bpjs_employment_no": createdEmployee.BPJSEmploymentNo,
                "bpjs_health_no":   createdEmployee.BPJSHealthNo,
                "address":          createdEmployee.Address,
                "phone":            createdEmployee.Phone,
                "join_date":        createdEmployee.JoinDate,
                "ktp_status":       createdEmployee.KTPStatus,
                "contract_no":      createdEmployee.ContractNo,
                "npwp_status":      createdEmployee.NPWPStatus,
                "contract_status":  createdEmployee.ContractStatus,
                "status":           createdEmployee.Status,
                "area":             areaNames,
                "created_at":       createdEmployee.CreatedAt.Format("2006-01-02 15:04:05"),
                "updated_at":       createdEmployee.UpdatedAt.Format("2006-01-02 15:04:05"),
            },
        },
    }
    
    c.JSON(http.StatusCreated, response)
}
