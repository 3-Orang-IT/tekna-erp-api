package adminUsecase

import (
	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
)

type NewsletterManagementUsecase interface {
	CreateNewsletter(newsletter *entity.Newsletter) error
	GetNewsletters(page, limit int, search string) ([]entity.Newsletter, error)
	GetNewsletterByID(id string) (*entity.Newsletter, error)
	UpdateNewsletter(id string, newsletter *entity.Newsletter) error
	DeleteNewsletter(id string) error
	GetNewslettersCount(search string) (int64, error) // Method to get total count of newsletters for pagination
}

type newsletterManagementUsecase struct {
	repo adminRepository.NewsletterManagementRepository
}

func NewNewsletterManagementUsecase(r adminRepository.NewsletterManagementRepository) NewsletterManagementUsecase {
	return &newsletterManagementUsecase{repo: r}
}

func (u *newsletterManagementUsecase) CreateNewsletter(newsletter *entity.Newsletter) error {
	return u.repo.CreateNewsletter(newsletter)
}

func (u *newsletterManagementUsecase) GetNewsletters(page, limit int, search string) ([]entity.Newsletter, error) {
	return u.repo.GetNewsletters(page, limit, search)
}

func (u *newsletterManagementUsecase) GetNewsletterByID(id string) (*entity.Newsletter, error) {
	return u.repo.GetNewsletterByID(id)
}

func (u *newsletterManagementUsecase) UpdateNewsletter(id string, newsletter *entity.Newsletter) error {
	return u.repo.UpdateNewsletter(id, newsletter)
}

func (u *newsletterManagementUsecase) DeleteNewsletter(id string) error {
	return u.repo.DeleteNewsletter(id)
}

// GetNewslettersCount gets the total count of newsletters for pagination
func (u *newsletterManagementUsecase) GetNewslettersCount(search string) (int64, error) {
	return u.repo.GetNewslettersCount(search)
}
