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

type JobPositionManagementHandler struct {
	usecase adminUsecase.JobPositionManagementUsecase
}

func NewJobPositionManagementHandler(r *gin.Engine, uc adminUsecase.JobPositionManagementUsecase, db *gorm.DB) {
	h := &JobPositionManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/job-positions", h.CreateJobPosition)
	admin.GET("/job-positions", h.GetJobPositions)
	admin.GET("/job-positions/:id", h.GetJobPositionByID)
	admin.PUT("/job-positions/:id", h.UpdateJobPosition)
	admin.DELETE("/job-positions/:id", h.DeleteJobPosition)
}

func (h *JobPositionManagementHandler) CreateJobPosition(c *gin.Context) {
	var input dto.CreateJobPositionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobPosition := entity.JobPosition{
		Name: input.Name,
	}

	if err := h.usecase.CreateJobPosition(&jobPosition); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "job position created successfully", "data": jobPosition})
}

func (h *JobPositionManagementHandler) GetJobPositions(c *gin.Context) {
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

	jobPositions, err := h.usecase.GetJobPositions(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

   var responseData []dto.JobPositionResponse
   for _, jp := range jobPositions {
	   responseData = append(responseData, dto.JobPositionResponse{
		   ID:   jp.ID,
		   Name: jp.Name,
		   CreatedAt: jp.CreatedAt.Format("02-01-2006 15:04"),
		   UpdatedAt: jp.UpdatedAt.Format("02-01-2006 15:04"),
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

func (h *JobPositionManagementHandler) GetJobPositionByID(c *gin.Context) {
	id := c.Param("id")
	jobPosition, err := h.usecase.GetJobPositionByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

   response := dto.JobPositionResponse{
	   ID:        jobPosition.ID,
	   Name:      jobPosition.Name,
	   CreatedAt: jobPosition.CreatedAt.Format("02-01-2006 15:04"),
	   UpdatedAt: jobPosition.UpdatedAt.Format("02-01-2006 15:04"),
   }
   c.JSON(http.StatusOK, gin.H{"data": response})
}

func (h *JobPositionManagementHandler) UpdateJobPosition(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job position ID"})
		return
	}

	var input dto.UpdateJobPositionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobPosition := entity.JobPosition{
		ID:   uint(idUint),
		Name: input.Name,
	}

	if err := h.usecase.UpdateJobPosition(id, &jobPosition); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

   response := dto.JobPositionResponse{
	   ID:        jobPosition.ID,
	   Name:      jobPosition.Name,
	   CreatedAt: jobPosition.CreatedAt.Format("02-01-2006 15:04"),
	   UpdatedAt: jobPosition.UpdatedAt.Format("02-01-2006 15:04"),
   }
   c.JSON(http.StatusOK, gin.H{"message": "job position updated successfully", "data": response})
}

func (h *JobPositionManagementHandler) DeleteJobPosition(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteJobPosition(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "job position deleted successfully", "data": gin.H{"id": id}})
}
