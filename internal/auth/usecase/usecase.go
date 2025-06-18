package usecase

import (
	"fmt"
	"strings"

	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/domain/entity"
	"github.com/3-Orang-IT/tekna-erp-api/internal/auth/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
    Register(user *entity.User) error
    Login(email, password string) (*entity.User, error)
    GetMenus(roleID uint) ([]entity.Menu, error)
}

type authUsecase struct {
    repo repository.AuthRepository
}

func NewAuthUsecase(r repository.AuthRepository) AuthUsecase {
    return &authUsecase{repo: r}
}

func (u *authUsecase) Register(user *entity.User) error {
    hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashed)
    
    err = u.repo.Register(user)
    if err != nil {
        if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "UNIQUE constraint failed") {
            return fmt.Errorf("email already registered")
        }
        return err
    }

    return nil
}

func (u *authUsecase) Login(email, password string) (*entity.User, error) {
    user, err := u.repo.FindByEmail(email)
    if err != nil {
        return nil, fmt.Errorf("invalid credentials")
    }

    if user == nil {
        return nil, fmt.Errorf("invalid credentials")
    }

    if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
        return nil, fmt.Errorf("invalid credentials")
    }

    return user, nil
}


func (u *authUsecase) GetMenus(roleID uint) ([]entity.Menu, error) {
    return u.repo.GetMenusByRoleID(roleID)
}
