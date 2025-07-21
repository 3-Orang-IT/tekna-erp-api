package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedMenus(db *gorm.DB) error {
	menusList := []entity.Menu{
		{Name: "Dashboard", URL: "/admin", Icon: "HomeIcon", Order: 1},
		{Name: "Master Data", Icon: "DatabaseIcon", Order: 2, Children: []entity.Menu{
			{Name: "Perusahaan & Organisasi", Icon: "BuildingOfficeIcon", Order: 1, Children: []entity.Menu{
				{Name: "Data Perusahaan", URL: "/master-data/perusahaan", Icon: "BuildingIcon", Order: 1},
				{Name: "Data Divisi", URL: "/master-data/divisi", Icon: "RectangleStackIcon", Order: 2},
				{Name: "Data Jabatan", URL: "/master-data/jabatan", Icon: "IdentificationIcon", Order: 3},
				{Name: "Data Departemen", URL: "/master-data/departemen", Icon: "UserGroupIcon", Order: 4},
				{Name: "Struktur Organisasi", URL: "/master-data/struktur-organisasi", Icon: "ChartBarIcon", Order: 5},
			}},
			{Name: "Manajemen Pengguna", Icon: "UsersIcon", Order: 2, Children: []entity.Menu{
				{Name: "Data Pengguna", URL: "/master-data/pengguna", Icon: "UserIcon", Order: 1},
				{Name: "Role & Permission", URL: "/master-data/role-permission", Icon: "LockClosedIcon", Order: 2},
				{Name: "Level Akses", URL: "/master-data/level-akses", Icon: "ShieldCheckIcon", Order: 3},
				{Name: "User Activity Log", URL: "/master-data/user-activity", Icon: "ClockIcon", Order: 4},
				{Name: "User Group", URL: "/master-data/user-group", Icon: "UsersIcon", Order: 5},
			}},
		}},
	}

	var seedMenu func(menu entity.Menu, parentID *uint) error
	seedMenu = func(menu entity.Menu, parentID *uint) error {
		menu.ParentID = parentID
		if err := db.Where("name = ? AND parent_id = ?", menu.Name, parentID).FirstOrCreate(&menu).Error; err != nil {
			log.Printf("Failed to create menu %s: %v", menu.Name, err)
			return err
		}
		log.Printf("Successfully added menu %s", menu.Name)

		for _, child := range menu.Children {
			if err := seedMenu(child, &menu.ID); err != nil {
				return err
			}
		}
		return nil
	}

	for _, menu := range menusList {
		if err := seedMenu(menu, nil); err != nil {
			return err
		}
	}

	return nil
}
