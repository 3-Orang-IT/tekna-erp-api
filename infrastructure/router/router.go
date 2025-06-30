package router

import (
	userManagementHandler "github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/handler"
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/interface/repository"
	adminUsecase "github.com/3-Orang-IT/tekna-erp-api/internal/admin/usecase"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/interface/handler"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/interface/repository"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Init repository dan usecase
	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	
	userManagementRepo := adminRepository.NewUserManagementRepository(db)
	userManagementUsecase := adminUsecase.NewUserManagementUsecase(userManagementRepo)
	// Register handler ke router
	handler.NewAuthHandler(r, authUsecase, db)
	userManagementHandler.NewUserManagementHandler(r, userManagementUsecase)

	return r
}
