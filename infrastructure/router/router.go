package router

import (
	"time"

	adminHandler "github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/handler"
	adminRepositoryImpl "github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/repository"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/interface/handler"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/interface/repository"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Add CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	// Init repository dan usecase
	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	
	userManagementRepo := adminRepositoryImpl.NewUserManagementRepository(db)
	userManagementUsecase := adminUsecase.NewUserManagementUsecase(userManagementRepo)
	
	// Register handler ke router
	handler.NewAuthHandler(r, authUsecase, db)
	adminHandler.NewUserManagementHandler(r, userManagementUsecase, db)

	return r
}
