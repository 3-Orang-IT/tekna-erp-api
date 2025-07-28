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

type BankAccountManagementHandler struct {
	usecase adminUsecase.BankAccountManagementUsecase
}

func NewBankAccountManagementHandler(r *gin.Engine, uc adminUsecase.BankAccountManagementUsecase, db *gorm.DB) {
	h := &BankAccountManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/bank-accounts", h.CreateBankAccount)
	admin.GET("/bank-accounts", h.GetBankAccounts)
	admin.GET("/bank-accounts/:id", h.GetBankAccountByID)
	admin.GET("/bank-accounts/:id/edit", h.GetEditBankAccountPage)
	admin.PUT("/bank-accounts/:id", h.UpdateBankAccount)
	admin.DELETE("/bank-accounts/:id", h.DeleteBankAccount)
}

func (h *BankAccountManagementHandler) CreateBankAccount(c *gin.Context) {
	var input dto.CreateBankAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bankAccount := entity.BankAccount{
		ChartOfAccountID: input.ChartOfAccountID,
		AccountNumber:    input.AccountNumber,
		BankName:         input.BankName,
		BranchAddress:    input.BranchAddress,
		CityID:           input.CityID,
		PhoneNumber:      input.PhoneNumber,
		Priority:         input.Priority,
	}

	if err := h.usecase.CreateBankAccount(&bankAccount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "bank account created successfully", "data": bankAccount})
}

func (h *BankAccountManagementHandler) GetBankAccounts(c *gin.Context) {
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

	bankAccounts, err := h.usecase.GetBankAccounts(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get total count for pagination
	totalData, err := h.usecase.GetBankAccountsCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(totalData) / limit
	if int(totalData)%limit > 0 {
		totalPages++
	}

	// Map bank accounts to response format
	var responseData []dto.BankAccountResponse
	for _, bankAccount := range bankAccounts {
		responseData = append(responseData, dto.BankAccountResponse{
			ID:               bankAccount.ID,
			ChartOfAccount:   bankAccount.ChartOfAccount.Name,
			AccountNumber:    bankAccount.AccountNumber,
			BankName:         bankAccount.BankName,
			BranchAddress:    bankAccount.BranchAddress,
			City:             bankAccount.City.Name,
			Province:         bankAccount.City.Province.Name,
			PhoneNumber:      bankAccount.PhoneNumber,
			Priority:         bankAccount.Priority,
			UpdatedAt:        "", // We will update this if needed
		})
	}

	response := gin.H{
		"data": responseData,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total_data":  totalData,
			"total_pages": totalPages,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *BankAccountManagementHandler) GetBankAccountByID(c *gin.Context) {
	id := c.Param("id")
	bankAccount, err := h.usecase.GetBankAccountByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "bank account not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.BankAccountResponse{
		ID:               bankAccount.ID,
		ChartOfAccount:   bankAccount.ChartOfAccount.Name,
		AccountNumber:    bankAccount.AccountNumber,
		BankName:         bankAccount.BankName,
		BranchAddress:    bankAccount.BranchAddress,
		City:             bankAccount.City.Name,
		Province:         bankAccount.City.Province.Name,
		PhoneNumber:      bankAccount.PhoneNumber,
		Priority:         bankAccount.Priority,
		UpdatedAt:        "", // We will update this if needed
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *BankAccountManagementHandler) GetEditBankAccountPage(c *gin.Context) {
	id := c.Param("id")

	// Fetch bank account by ID
	bankAccount, err := h.usecase.GetBankAccountByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "bank account not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch list of cities
	cities, err := h.usecase.GetCities(1, 1000, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch list of chart of accounts
	chartOfAccounts, err := h.usecase.GetChartOfAccounts(1, 1000, "")
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

	var chartOfAccountList []dto.ChartOfAccountResponse
	for _, coa := range chartOfAccounts {
		chartOfAccountList = append(chartOfAccountList, dto.ChartOfAccountResponse{
			ID:   coa.ID,
			Code: coa.Code,
			Name: coa.Name,
		})
	}

	response := gin.H{
		"data": bankAccount,
		"references": gin.H{
			"cities":           cityList,
			"chart_of_accounts": chartOfAccountList,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *BankAccountManagementHandler) UpdateBankAccount(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	var input dto.UpdateBankAccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bankAccount := entity.BankAccount{
		ID:               uint(idUint),
		ChartOfAccountID: input.ChartOfAccountID,
		AccountNumber:    input.AccountNumber,
		BankName:         input.BankName,
		BranchAddress:    input.BranchAddress,
		CityID:           input.CityID,
		PhoneNumber:      input.PhoneNumber,
		Priority:         input.Priority,
	}

	if err := h.usecase.UpdateBankAccount(id, &bankAccount); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "bank account not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "bank account updated successfully", "data": bankAccount})
}

func (h *BankAccountManagementHandler) DeleteBankAccount(c *gin.Context) {
	id := c.Param("id")

	if err := h.usecase.DeleteBankAccount(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "bank account not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "bank account deleted successfully", "data": gin.H{"id": id}})
}
