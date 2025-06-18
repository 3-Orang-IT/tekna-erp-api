package repository
import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type AuthRepository interface {
    Register(user *entity.User) error
    FindByUsername(username string) (*entity.User, error)
    GetMenusByUserID(userId uint) ([]entity.Menu, error)
}
