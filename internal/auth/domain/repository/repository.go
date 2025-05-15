package repository
import "github.com/3-Orang-IT/tekna-erp-api/internal/auth/domain/entity"

type AuthRepository interface {
    Register(user *entity.User) error
    FindByEmail(email string) (*entity.User, error)
    GetMenusByRoleID(roleID uint) ([]entity.Menu, error)
}
