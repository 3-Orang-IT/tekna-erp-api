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

type CompanyManagementHandler struct {
	usecase adminUsecase.CompanyManagementUsecase
}

func NewCompanyManagementHandler(r *gin.Engine, uc adminUsecase.CompanyManagementUsecase, db *gorm.DB) {
	h := &CompanyManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/companies", h.CreateCompany)
	admin.GET("/companies", h.GetCompanies)
	admin.GET("/companies/:id", h.GetCompanyByID)
	admin.PUT("/companies/:id", h.UpdateCompany)
	admin.DELETE("/companies/:id", h.DeleteCompany)
}

func (h *CompanyManagementHandler) CreateCompany(c *gin.Context) {
	var input dto.CreateCompanyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company := entity.Company{
		Name:             input.Name,
		Address:          input.Address,
		CityID:           input.CityID,
		Phone:            input.Phone,
		Fax:              input.Fax,
		Email:            input.Email,
		StartHour:        input.StartHour,
		EndHour:          input.EndHour,
		Latitude:         input.Latitude,
		Longitude:        input.Longitude,
		TotalShares:      input.TotalShares,
		AnnualLeaveQuota: input.AnnualLeaveQuota,
	}

	if err := h.usecase.CreateCompany(&company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "company created successfully", "data": company})
}

func (h *CompanyManagementHandler) GetCompanies(c *gin.Context) {
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

	companies, err := h.usecase.GetCompanies(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map companies to the desired response format using CompanyResponse DTO
	var responseData []dto.CompanyResponse
	for _, company := range companies {
		responseData = append(responseData, dto.CompanyResponse{
			Name:             company.Name,
			Address:          company.Address,
			City:             company.City.Name,
			Province:         company.City.Province.Name,
			Telp:             company.Phone,
			Fax:              company.Fax,
			Email:            company.Email,
			StartHour:        company.StartHour,
			EndHour:          company.EndHour,
			Latitude:         company.Latitude,
			Longitude:        company.Longitude,
			TotalShares:      company.TotalShares,
			AnnualLeaveQuota: company.AnnualLeaveQuota,
			UpdatedAt:        "", // Placeholder for UpdatedAt
		})
	}

	response := gin.H{
		"data": responseData,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *CompanyManagementHandler) GetCompanyByID(c *gin.Context) {
	id := c.Param("id")
	company, err := h.usecase.GetCompanyByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": company})
}

func (h *CompanyManagementHandler) UpdateCompany(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid company ID"})
		return
	}

	var input dto.UpdateCompanyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company := entity.Company{
		ID:               uint(idUint),
		Name:             input.Name,
		Address:          input.Address,
		CityID:           input.CityID,
		Phone:            input.Phone,
		Fax:              input.Fax,
		Email:            input.Email,
		StartHour:        input.StartHour,
		EndHour:          input.EndHour,
		Latitude:         input.Latitude,
		Longitude:        input.Longitude,
		TotalShares:      input.TotalShares,
		AnnualLeaveQuota: input.AnnualLeaveQuota,
	}

	if err := h.usecase.UpdateCompany(id, &company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "company updated successfully", "data": company})
}

func (h *CompanyManagementHandler) DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteCompany(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "company deleted successfully", "data": gin.H{"id": id}})
}
