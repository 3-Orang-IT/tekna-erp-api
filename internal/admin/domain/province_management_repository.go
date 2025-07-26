package adminRepository

import "github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"

type ProvinceManagementRepository interface {
	   CreateProvince(province *entity.Province) error
	   GetProvinces(page, limit int, search string) ([]entity.Province, error)
	   GetProvinceByID(id string) (*entity.Province, error)
	   UpdateProvince(id string, province *entity.Province) error
	   DeleteProvince(id string) error
}
