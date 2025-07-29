package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type NewsletterManagementRepository interface {
	CreateNewsletter(newsletter *entity.Newsletter) error
	GetNewsletters(page, limit int, search string) ([]entity.Newsletter, error)
	GetNewsletterByID(id string) (*entity.Newsletter, error)
	UpdateNewsletter(id string, newsletter *entity.Newsletter) error
	DeleteNewsletter(id string) error
	// Method to get total count of newsletters for pagination
	GetNewslettersCount(search string) (int64, error)
}
