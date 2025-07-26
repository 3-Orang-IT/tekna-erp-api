package adminHandler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"io"
	"os"
	"time"

	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/dto"
	"github.com/3-Orang-IT/tekna-erp-api/internal/admin/middleware"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserManagementHandler struct {
	usecase adminUsecase.UserManagementUsecase
}

func NewUserManagementHandler(r *gin.Engine, uc adminUsecase.UserManagementUsecase, db *gorm.DB) {
	h := &UserManagementHandler{uc}
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AdminRoleMiddleware(db))
	admin.POST("/users", h.CreateUser)
	admin.GET("/users", h.GetUsers)
	admin.GET("/users/:id", h.GetUserByID)
	admin.PUT("/users/:id", h.UpdateUser)
	admin.DELETE("/users/:id", h.DeleteUser)
}

func validateUser(user *entity.User) error {
	if strings.TrimSpace(user.Username) == "" {
		return fmt.Errorf("username is required")
	}
	if strings.TrimSpace(user.Email) == "" {
		return fmt.Errorf("email is required")
	}
	if strings.TrimSpace(user.Password) == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

func (h *UserManagementHandler) CreateUser(c *gin.Context) {
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
	status := c.PostForm("status")
	roleIDs := c.PostFormArray("roles")

	// Handle file upload
	file, header, err := c.Request.FormFile("photo_profile")
	var photoProfileURL string
	if err == nil && header != nil {
		sanitizedFilename := strings.ReplaceAll(header.Filename, " ", "_")
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), sanitizedFilename)
		savePath := fmt.Sprintf("uploads/profile/%s", filename)
		// Ensure directory exists
		if err := ensureDir("uploads/profile"); err != nil {
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

	// Convert roleIDs to []uint
	var roleIDsUint []uint
	for _, rid := range roleIDs {
		id, err := strconv.ParseUint(rid, 10, 64)
		if err == nil {
			roleIDsUint = append(roleIDsUint, uint(id))
		}
	}

	user := entity.User{
		Username:        username,
		Password:        password,
		Name:            name,
		Email:           email,
		Telp:            telp,
		PhotoProfileURL: photoProfileURL,
		Status:          status,
	}
	for _, roleID := range roleIDsUint {
		user.Role = append(user.Role, entity.Role{ID: roleID})
	}

	if err := validateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully", "data": user})
}

// ensureDir creates a directory if it does not exist
func ensureDir(dirName string) error {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func (h *UserManagementHandler) GetUsers(c *gin.Context) {
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

	users, err := h.usecase.GetUsers(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userResponses []dto.UserResponse

	baseUrl := os.Getenv("BASE_URL")
	for _, user := range users {
		var roleNames []string
		for _, role := range user.Role {
			roleNames = append(roleNames, role.Name)
		}

		var photoURL string
		if user.PhotoProfileURL != "" {
			photoURL = baseUrl + user.PhotoProfileURL
		} else {
			photoURL = ""
		}

		userResponses = append(userResponses, dto.UserResponse{
			ID:              user.ID,
			Username:        user.Username,
			Name:            user.Name,
			Email:           user.Email,
			Telp:            user.Telp,
			PhotoProfileURL: photoURL,
			Status:          user.Status,
			Roles:           roleNames,
			CreatedAt:       user.CreatedAt,
			UpdatedAt:       user.UpdatedAt,
		})
	}

	response := gin.H{
		"data": userResponses,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
		},
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserManagementHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := h.usecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Convert roles to []string
	var roleNames []string
	for _, role := range user.Role {
		roleNames = append(roleNames, role.Name)
	}
	
	baseUrl := os.Getenv("BASE_URL")
	var photoURL string
	if user.PhotoProfileURL != "" {
		photoURL = baseUrl + user.PhotoProfileURL
	} else {
		photoURL = ""
	}

	// Map ke response DTO
	userResponse := dto.UserResponse{
		ID:              user.ID,
		Username:        user.Username,
		Name:            user.Name,
		Email:           user.Email,
		Telp:            user.Telp,
		PhotoProfileURL: photoURL,
		Status:          user.Status,
		Roles:           roleNames,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"data": userResponse})
}

func (h *UserManagementHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	// Parse form-data (multipart)
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10MB max
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form: " + err.Error()})
		return
	}

	// Get fields from form-data
	username := c.PostForm("username")
	password := c.PostForm("password")
	name := c.PostForm("name")
	telp := c.PostForm("telp")
	status := c.PostForm("status")
	roleIDs := c.PostFormArray("roles")

	// Handle file upload
	file, header, err := c.Request.FormFile("photo_profile")
	var photoProfileURL string
	if err == nil && header != nil {
		sanitizedFilename := strings.ReplaceAll(header.Filename, " ", "_")
		filename := fmt.Sprintf("%d_%s", time.Now().UnixNano(), sanitizedFilename)
		savePath := fmt.Sprintf("uploads/profile/%s", filename)
		// Ensure directory exists
		if err := ensureDir("uploads/profile"); err != nil {
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
	} else {
		// If no new file uploaded, keep the old value if provided
		photoProfileURL = c.PostForm("photo_profile_url")
	}

	// Convert roleIDs to []uint
	var roleIDsUint []uint
	for _, rid := range roleIDs {
		idUint, err := strconv.ParseUint(rid, 10, 64)
		if err == nil {
			roleIDsUint = append(roleIDsUint, uint(idUint))
		}
	}

	user := entity.User{
		Username:        username,
		Password:        password,
		Name:            name,
		Telp:            telp,
		PhotoProfileURL: photoProfileURL,
		Status:          status,
	}
	for _, roleID := range roleIDsUint {
		user.Role = append(user.Role, entity.Role{ID: roleID})
	}

	// Panggil usecase
	if err := h.usecase.UpdateUser(id, &user); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user ID not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully", "data": user})
}

func (h *UserManagementHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.usecase.DeleteUser(id); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "user ID not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
		"data": gin.H{"id": id},
	})
}