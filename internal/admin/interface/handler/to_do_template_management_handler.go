package adminHandler

import (
	"math"
	"net/http"
	"strconv"

	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/dto"
	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/middleware"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ToDoTemplateManagementHandler struct {
	usecase adminUsecase.ToDoTemplateManagementUsecase
}

func NewToDoTemplateManagementHandler(r *gin.Engine, uc adminUsecase.ToDoTemplateManagementUsecase, db *gorm.DB) {
	h := &ToDoTemplateManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/todo-templates", h.CreateToDoTemplate)
	admin.GET("/todo-templates", h.GetToDoTemplatesByJobPosition)
	admin.GET("/todo-templates/list", h.GetToDoTemplates)
	admin.GET("/todo-templates/:id", h.GetToDoTemplateByID)
	admin.GET("/todo-templates/:id/edit", h.GetEditToDoTemplatePage)
	admin.PUT("/todo-templates/:id", h.UpdateToDoTemplate)
	admin.DELETE("/todo-templates/:id", h.DeleteToDoTemplate)
}

func (h *ToDoTemplateManagementHandler) CreateToDoTemplate(c *gin.Context) {
	var input dto.CreateToDoTemplateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toDoTemplate := entity.ToDoTemplate{
		JobPositionID: input.JobPositionID,
		Activity:      input.Activity,
		Priority:      input.Priority,
		OrderNumber:   input.OrderNumber,
	}

	if err := h.usecase.CreateToDoTemplate(&toDoTemplate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "to-do template created successfully", "data": toDoTemplate})
}

func (h *ToDoTemplateManagementHandler) GetToDoTemplates(c *gin.Context) {
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
	toDoTemplates, err := h.usecase.GetToDoTemplates(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	// Get total count for pagination
	totalData, err := h.usecase.GetToDoTemplatesCount(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Calculate total pages
	totalPages := int(totalData) / limit
	if int(totalData)%limit > 0 {
		totalPages++
	}
	
	var response []dto.ToDoTemplateResponse
	for _, template := range toDoTemplates {
		response = append(response, dto.ToDoTemplateResponse{
			ID:            template.ID,
			JobPositionID: template.JobPositionID,
			JobPosition:   template.JobPosition.Name,
			Activity:      template.Activity,
			Priority:      template.Priority,
			OrderNumber:   template.OrderNumber,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": response,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total_data":  totalData,
			"total_pages": totalPages,
		},
	})
}



func (h *ToDoTemplateManagementHandler) GetToDoTemplateByID(c *gin.Context) {
	id := c.Param("id")
	toDoTemplate, err := h.usecase.GetToDoTemplateByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "to-do template not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get all templates for the same job position
	templatesForJobPosition, err := h.usecase.GetToDoTemplatesByJobPositionID(toDoTemplate.JobPositionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map each template to response format
	var activitiesResponse []dto.ToDoTemplateResponse
	for _, template := range templatesForJobPosition {
		activitiesResponse = append(activitiesResponse, dto.ToDoTemplateResponse{
			ID:            template.ID,
			JobPositionID: template.JobPositionID,
			JobPosition:   template.JobPosition.Name,
			Activity:      template.Activity,
			Priority:      template.Priority,
			OrderNumber:   template.OrderNumber,
		})
	}

	// Create the grouped response
	groupedResponse := dto.GroupedToDoTemplateResponse{
		JobPositionID: toDoTemplate.JobPositionID,
		JobPosition:   toDoTemplate.JobPosition.Name,
		Activities:    activitiesResponse,
	}

	c.JSON(http.StatusOK, gin.H{"data": groupedResponse})
}

func (h *ToDoTemplateManagementHandler) GetEditToDoTemplatePage(c *gin.Context) {
	id := c.Param("id")

	// Fetch to-do template by ID
	toDoTemplate, err := h.usecase.GetToDoTemplateByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "to-do template not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch job positions for reference
	jobPositions, err := h.usecase.GetJobPositions(1, 1000, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var jobPositionOptions []dto.JobPositionOption
	for _, jobPosition := range jobPositions {
		jobPositionOptions = append(jobPositionOptions, dto.JobPositionOption{
			ID:   jobPosition.ID,
			Name: jobPosition.Name,
		})
	}

	response := gin.H{
		"data": toDoTemplate,
		"references": gin.H{
			"job_positions": jobPositionOptions,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *ToDoTemplateManagementHandler) UpdateToDoTemplate(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
		return
	}

	var input dto.UpdateToDoTemplateInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	toDoTemplate := entity.ToDoTemplate{
		ID:            uint(idUint),
		JobPositionID: input.JobPositionID,
		Activity:      input.Activity,
		Priority:      input.Priority,
		OrderNumber:   input.OrderNumber,
	}

	if err := h.usecase.UpdateToDoTemplate(id, &toDoTemplate); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "to-do template not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "to-do template updated successfully", "data": toDoTemplate})
}

func (h *ToDoTemplateManagementHandler) DeleteToDoTemplate(c *gin.Context) {
	id := c.Param("id")

	if err := h.usecase.DeleteToDoTemplate(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "to-do template not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "to-do template deleted successfully", "data": gin.H{"id": id}})
}

func (h *ToDoTemplateManagementHandler) GetToDoTemplatesByJobPosition(c *gin.Context) {
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

	// Get templates grouped by job position with pagination
	groupedTemplates, total, err := h.usecase.GetToDoTemplatesByJobPosition(page, limit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a map to store job position details
	jobPositionsMap := make(map[uint]string)
	for jobPositionID := range groupedTemplates {
		// We need at least one template from each job position to get the name
		if len(groupedTemplates[jobPositionID]) > 0 {
			jobPositionsMap[jobPositionID] = groupedTemplates[jobPositionID][0].JobPosition.Name
		}
	}

	// Format the response
	var response []dto.GroupedToDoTemplateResponse
	for jobPositionID, templates := range groupedTemplates {
		var activitiesResponse []dto.ToDoTemplateResponse
		
		for _, template := range templates {
			activitiesResponse = append(activitiesResponse, dto.ToDoTemplateResponse{
				ID:            template.ID,
				JobPositionID: template.JobPositionID,
				JobPosition:   template.JobPosition.Name,
				Activity:      template.Activity,
				Priority:      template.Priority,
				OrderNumber:   template.OrderNumber,
			})
		}
		
		groupedResponse := dto.GroupedToDoTemplateResponse{
			JobPositionID: jobPositionID,
			JobPosition:   jobPositionsMap[jobPositionID],
			Activities:    activitiesResponse,
		}
		
		response = append(response, groupedResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total_data": total,
			"total_pages": int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}
