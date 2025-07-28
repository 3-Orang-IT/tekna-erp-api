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

type CustomerManagementHandler struct {
	usecase adminUsecase.CustomerManagementUsecase
}

func NewCustomerManagementHandler(r *gin.Engine, uc adminUsecase.CustomerManagementUsecase, db *gorm.DB) {
	h := &CustomerManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/customers", h.CreateCustomer)
	admin.GET("/customers", h.GetCustomers)
	admin.GET("/customers/:id", h.GetCustomerByID)
	admin.GET("/customers/:id/edit", h.GetEditCustomerPage)
	admin.PUT("/customers/:id", h.UpdateCustomer)
	admin.DELETE("/customers/:id", h.DeleteCustomer)
}

func (h *CustomerManagementHandler) CreateCustomer(c *gin.Context) {
	var input dto.CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := entity.Customer{
		UserID:            input.UserID,
		AreaID:            input.AreaID,
		CityID:            input.CityID,
		Name:              input.Name,
		Code:              input.Code,
		InvoiceName:       input.InvoiceName,
		Address:           input.Address,
		Phone:             input.Phone,
		Email:             input.Email,
		Tax:               input.Tax,
		Greeting:          input.Greeting,
		ContactPersonName: input.ContactPersonName,
		ContactPhone:      input.ContactPhone,
		Segment:           input.Segment,
		Type:              input.Type,
		NPWP:              input.NPWP,
		Status:            input.Status,
		BEName:            input.BEName,
		ProcurementType:   input.ProcurementType,
		MarketingName:     input.MarketingName,
		Note:              input.Note,
		PaymentTerm:       input.PaymentTerm,
		Level:             input.Level,
	}

	if err := h.usecase.CreateCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "customer created successfully", "data": customer})
}

func (h *CustomerManagementHandler) GetCustomers(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
		return
	}

	search := c.DefaultQuery("search", "")

	customers, err := h.usecase.GetCustomers(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responseData []dto.CustomerResponse

	for _, customer := range customers {
		responseData = append(responseData, dto.CustomerResponse{
			ID:          customer.ID,
			Name:        customer.Name,
			InvoiceName: customer.InvoiceName,
			Code:        customer.Code,
			City:        customer.City.Name,
			Province:    customer.City.Province.Name,
			Area:        customer.Area.Name,
			Type:        customer.Type,
			Level:       customer.Level,
			Status:      customer.Status,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": responseData, "pagination": gin.H{"page": page, "limit": limit}})
}

func (h *CustomerManagementHandler) GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	customer, err := h.usecase.GetCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (h *CustomerManagementHandler) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := entity.Customer{
		AreaID:            input.AreaID,
		CityID:            input.CityID,
		Name:              input.Name,
		Code:              input.Code,
		InvoiceName:       input.InvoiceName,
		Address:           input.Address,
		Phone:             input.Phone,
		Email:             input.Email,
		Tax:               input.Tax,
		Greeting:          input.Greeting,
		ContactPersonName: input.ContactPersonName,
		ContactPhone:      input.ContactPhone,
		Segment:           input.Segment,
		Type:              input.Type,
		NPWP:              input.NPWP,
		Status:            input.Status,
		BEName:            input.BEName,
		ProcurementType:   input.ProcurementType,
		MarketingName:     input.MarketingName,
		Note:              input.Note,
		PaymentTerm:       input.PaymentTerm,
		Level:             input.Level,
	}

	if err := h.usecase.UpdateCustomer(id, &customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer updated successfully"})
}

func (h *CustomerManagementHandler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteCustomer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer deleted successfully"})
}

func (h *CustomerManagementHandler) GetEditCustomerPage(c *gin.Context) {
	id := c.Param("id")
	customer, err := h.usecase.GetCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cities, err := h.usecase.GetCities(1, 100, "") // Example: Fetch all cities for reference
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer, "references": gin.H{"cities": cities}})
}
