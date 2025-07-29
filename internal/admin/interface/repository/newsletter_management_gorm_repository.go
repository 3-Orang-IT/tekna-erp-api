package adminRepositoryImpl

import (
	"strings"

	adminRepository "github.com/3-Orang-IT/tekna-erp-api/internal/admin/domain"
	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

type newsletterManagementRepo struct {
	db *gorm.DB
}

func NewNewsletterManagementRepository(db *gorm.DB) adminRepository.NewsletterManagementRepository {
	return &newsletterManagementRepo{db: db}
}

func (r *newsletterManagementRepo) CreateNewsletter(newsletter *entity.Newsletter) error {
	return r.db.Create(newsletter).Error
}

func (r *newsletterManagementRepo) GetNewsletters(page, limit int, search string) ([]entity.Newsletter, error) {
	var newsletters []entity.Newsletter
	offset := (page - 1) * limit
	query := r.db
	if search != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Limit(limit).Offset(offset).Order("id ASC").Find(&newsletters).Error; err != nil {
		return nil, err
	}
	return newsletters, nil
}

func (r *newsletterManagementRepo) GetNewsletterByID(id string) (*entity.Newsletter, error) {
	var newsletter entity.Newsletter
	if err := r.db.First(&newsletter, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &newsletter, nil
}

func (r *newsletterManagementRepo) UpdateNewsletter(id string, newsletter *entity.Newsletter) error {
	var existingNewsletter entity.Newsletter
	if err := r.db.First(&existingNewsletter, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	return r.db.Model(&existingNewsletter).Updates(newsletter).Error
}

func (r *newsletterManagementRepo) DeleteNewsletter(id string) error {
	var newsletter entity.Newsletter
	if err := r.db.First(&newsletter, "id = ?", id).Error; err != nil {
		return err
	}
	return r.db.Delete(&newsletter).Error
}

// Method to get total count of newsletters for pagination
func (r *newsletterManagementRepo) GetNewslettersCount(search string) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Newsletter{})
	if search != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(search)+"%")
	}
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
