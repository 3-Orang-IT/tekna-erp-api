package seeders

import (
	"log"

	"github.com/3-Orang-IT/tekna-erp-api/internal/common/entity"
	"gorm.io/gorm"
)

func SeedMenus(db *gorm.DB) error {
	menusList := []entity.Menu{
		{Name: "Dashboard", URL: "/dashboard", Icon: "HomeIcon", Order: 1},
		{Name: "Master Data", Icon: "DatabaseIcon", Order: 2, Children: []entity.Menu{
			{Name: "Perusahaan & Organisasi", Icon: "BuildingOfficeIcon", Order: 1, Children: []entity.Menu{
				{Name: "Data Perusahaan", URL: "/master-data/perusahaan", Icon: "BuildingIcon", Order: 1},
				{Name: "Data Divisi", URL: "/master-data/divisi", Icon: "RectangleStackIcon", Order: 2},
				{Name: "Data Jabatan", URL: "/master-data/jabatan", Icon: "IdentificationIcon", Order: 3},
				{Name: "Data Pegawai", URL: "/master-data/pegawai", Icon: "IdentificationIcon", Order: 4},
				// {Name: "Data Departemen", URL: "/master-data/departemen", Icon: "UserGroupIcon", Order: 4},
				// {Name: "Struktur Organisasi", URL: "/master-data/struktur-organisasi", Icon: "ChartBarIcon", Order: 5},
			}},
			{Name: "Manajemen Pengguna", Icon: "UsersIcon", Order: 2, Children: []entity.Menu{
				{Name: "Data Pengguna", URL: "/master-data/pengguna", Icon: "UserIcon", Order: 1},
				{Name: "Role & Permission", URL: "/master-data/role-permission", Icon: "LockClosedIcon", Order: 2},
				{Name: "Data Menu", URL: "/master-data/menu", Icon: "CubeIcon", Order: 3},
				// {Name: "Level Akses", URL: "/master-data/level-akses", Icon: "ShieldCheckIcon", Order: 3},
				// {Name: "Data Modul", URL: "/master-data/modul", Icon: "CubeIcon", Order: 4},
				// {Name: "User Activity Log", URL: "/master-data/user-activity", Icon: "ClockIcon", Order: 4},
			}},
			{Name: "Produk & Inventory", Icon: "CubeIcon", Order: 3, Children: []entity.Menu{
				{Name: "Data Produk", URL: "/master-data/produk", Icon: "CubeIcon", Order: 1},
				{Name: "Kategori Produk", URL: "/master-data/kategori-produk", Icon: "TagIcon", Order: 2},
				{Name: "Unit Satuan", URL: "/master-data/unit-satuan", Icon: "ScaleIcon", Order: 3},
				{Name: "Brand/Merk", URL: "/master-data/brand", Icon: "SparklesIcon", Order: 4},
				{Name: "Business Unit", URL: "/master-data/business", Icon: "BriefcaseIcon", Order: 5},
				{Name: "Kategori Produk Alternatif", URL: "/master-data/kategori-produk-alternatif", Icon: "TagIcon", Order: 6},
				// {Name: "Gudang/Warehouse", URL: "/master-data/gudang", Icon: "HomeModernIcon", Order: 5},
			}},
			{Name: "Partner & Relasi", Icon: "HandshakeIcon", Order: 4, Children: []entity.Menu{
				{Name: "Data Supplier", URL: "/master-data/supplier", Icon: "TruckIcon", Order: 1},
				{Name: "Data Customer", URL: "/master-data/customer", Icon: "UserCircleIcon", Order: 2},
				// {Name: "Data Vendor", URL: "/master-data/vendor", Icon: "BuildingStorefrontIcon", Order: 3},
				// {Name: "Kategori Partner", URL: "/master-data/kategori-partner", Icon: "FolderIcon", Order: 4},
			}},
			{Name: "Keuangan & Akuntansi", Icon: "CurrencyDollarIcon", Order: 5, Children: []entity.Menu{
				{Name: "Data Rekening Bank", URL: "/master-data/rekening-bank", Icon: "BanknotesIcon", Order: 1},
				{Name: "Chart of Account (COA)", URL: "/master-data/coa", Icon: "DocumentChartBarIcon", Order: 2},
				// {Name: "Mata Uang", URL: "/master-data/mata-uang", Icon: "CurrencyDollarIcon", Order: 3},
				// {Name: "Pajak & Tax", URL: "/master-data/pajak", Icon: "CalculatorIcon", Order: 4},
				// {Name: "Metode Pembayaran", URL: "/master-data/metode-pembayaran", Icon: "CreditCardIcon", Order: 5},
			}},
			{Name: "Dokumen", Icon: "CogIcon", Order: 6, Children: []entity.Menu{
				{Name: "To do Template", URL: "/master-data/to-do-template", Icon: "DocumentIcon", Order: 1},
				{Name: "News Letters", URL: "/master-data/news-letters", Icon: "DocumentIcon", Order: 2},
				{Name: "Kategori Dokumen", URL: "/master-data/kategori-dokumen", Icon: "DocumentIcon", Order: 3},
				{Name: "Dokumen", URL: "/master-data/dokumen", Icon: "DocumentIcon", Order: 4},
				{Name: "Biaya Perjalanan", URL: "/master-data/biaya-perjalanan", Icon: "EnvelopeIcon", Order: 5},
			}},
			{Name: "Sistem & Konfigurasi", Icon: "CogIcon", Order: 7, Children: []entity.Menu{
				{Name: "Provinsi", URL: "/master-data/provinsi", Icon: "DocumentIcon", Order: 1},
				{Name: "Kota & Kabupaten", URL: "/master-data/kota", Icon: "DocumentIcon", Order: 2},
				{Name: "Area", URL: "/master-data/area", Icon: "DocumentIcon", Order: 3},
				{Name: "Budget Kategori", URL: "/master-data/budget-kategori", Icon: "DocumentIcon", Order: 4},
				// {Name: "Kategori Dokumen", URL: "/master-data/kategori-dokumen", Icon: "DocumentIcon", Order: 4},
				// {Name: "Template Email", URL: "/master-data/template-email", Icon: "EnvelopeIcon", Order: 5},
				// {Name: "Workflow Setting", URL: "/master-data/workflow", Icon: "ArrowPathIcon", Order: 6},
				// {Name: "System Parameter", URL: "/master-data/system-parameter", Icon: "AdjustmentsHorizontalIcon", Order: 7},
			}},
			{Name: "Annual Targets", URL: "/master-data/annual-target", Icon: "TargetIcon", Order: 7},
			// {Name: "System Info", URL: "/master-data/system-info", Icon: "InformationCircleIcon", Order: 8},
		}},
		{Name: "Pembelian", Icon: "ShoppingCartIcon", Order: 3, Children: []entity.Menu{
			{Name: "Purchase Order (PO)", URL: "/pembelian/purchase-order", Icon: "DocumentTextIcon", Order: 1},
			{Name: "Received Order PO", URL: "/pembelian/received-order", Icon: "TruckIcon", Order: 2},
			{Name: "Invoice PO", URL: "/pembelian/invoice-po", Icon: "ReceiptPercentIcon", Order: 3},
			{Name: "Payment PO", URL: "/pembelian/payment-po", Icon: "BanknotesIcon", Order: 4},
			{Name: "Retur PO", URL: "/pembelian/retur-po", Icon: "ArrowUturnLeftIcon", Order: 5},
			{Name: "Rekap Hutang", URL: "/pembelian/rekap-hutang", Icon: "DocumentChartBarIcon", Order: 6},
			{Name: "Rekap PO", URL: "/pembelian/rekap-po", Icon: "ClipboardDocumentCheckIcon", Order: 7},
			{Name: "Request Order", URL: "/pembelian/request-order", Icon: "ClipboardIcon", Order: 8},
		}},
		{Name: "Penjualan", Icon: "TrendingUpIcon", Order: 4, Children: []entity.Menu{
			{Name: "Core Sales Transaction Process", Icon: "DocumentTextIcon", Order: 1, Children: []entity.Menu{
				{Name: "Quotation", URL: "/penjualan/quotation", Icon: "DocumentDuplicateIcon", Order: 1},
				{Name: "Sales Order (SO)", URL: "/penjualan/sales-order", Icon: "DocumentTextIcon", Order: 2},
				{Name: "Invoice DP SO", URL: "/penjualan/invoice-dp-so", Icon: "DocumentIcon", Order: 3},
				{Name: "Delivery Order", URL: "/penjualan/delivery-order", Icon: "TruckIcon", Order: 4},
				{Name: "Invoice SO", URL: "/penjualan/invoice-so", Icon: "ReceiptPercentIcon", Order: 5},
				{Name: "Payment SO", URL: "/penjualan/payment-so", Icon: "BanknotesIcon", Order: 6},
			}},
			{Name: "Special Handling and Adjustments", Icon: "AdjustmentsHorizontalIcon", Order: 2, Children: []entity.Menu{
				{Name: "Retur SO", URL: "/penjualan/retur-so", Icon: "ArrowUturnLeftIcon", Order: 1},
				{Name: "CN Project", URL: "/penjualan/cn-project", Icon: "RectangleGroupIcon", Order: 2},
				{Name: "Invoice DP PO", URL: "/penjualan/invoice-dp-supplier", Icon: "DocumentIcon", Order: 3},
			}},
			{Name: "Reporting and Analytics", Icon: "ChartPieIcon", Order: 3, Children: []entity.Menu{
				{Name: "SO Transaction Progress Recap", URL: "/penjualan/rekap-progres", Icon: "ChartBarIcon", Order: 1},
				{Name: "SO Recap", URL: "/penjualan/rekap-so", Icon: "ClipboardDocumentCheckIcon", Order: 2},
				{Name: "Receivables Recap", URL: "/penjualan/rekap-piutang", Icon: "DocumentChartBarIcon", Order: 3},
				{Name: "Sales Journal", URL: "/penjualan/sales-jurnal", Icon: "BookOpenIcon", Order: 4},
				{Name: "Target", URL: "/penjualan/target", Icon: "FlagIcon", Order: 5},
				{Name: "Sales Report", URL: "/penjualan/report", Icon: "ChartPieIcon", Order: 6},
				{Name: "Logistics Delivery Schedule", URL: "/penjualan/jadwal-pengiriman", Icon: "CalendarIcon", Order: 7},
			}},
		}},
		{Name: "Stok Barang", URL: "/stok-barang", Icon: "PackageIcon", Order: 5, Children: []entity.Menu{
			{Name: "Tutup Buku Stok", URL: "/stok-barang/tutup-buku", Icon: "BookOpenIcon", Order: 1},
			{Name: "Kartu Stok", URL: "/stok-barang/kartu-stok", Icon: "ClipboardListIcon", Order: 2},
			{Name: "Stok Opname", URL: "/stok-barang/stok-opname", Icon: "ClipboardCheckIcon", Order: 3},
			{Name: "Stok Valuation", URL: "/stok-barang/stok-valuation", Icon: "ChartBarIcon", Order: 4},
		}},
		{Name: "Keuangan", URL: "/keuangan", Icon: "CreditCardIcon", Order: 6, Children: []entity.Menu{
			{Name: "Laporan Keuangan", Icon: "ChartBarIcon", Order: 1, Children: []entity.Menu{
				{Name: "Laporan Buku Besar", URL: "/keuangan/laporan-buku-besar", Icon: "BookOpenIcon", Order: 1},
				{Name: "Laporan Jurnal Umum", URL: "/keuangan/laporan-jurnal-umum", Icon: "ClipboardListIcon", Order: 2},
				{Name: "Laporan Perubahan Modal", URL: "/keuangan/laporan-perubahan-modal", Icon: "AdjustmentsHorizontalIcon", Order: 3},
				{Name: "Laporan Laba Rugi", URL: "/keuangan/laporan-laba-rugi", Icon: "CurrencyDollarIcon", Order: 4},
				{Name: "Laporan Neraca", URL: "/keuangan/laporan-neraca", Icon: "ScaleIcon", Order: 5},
				{Name: "Rasio Keuangan", URL: "/keuangan/rasio-keuangan", Icon: "ChartPieIcon", Order: 6},
			}},
			{Name: "Transaksi dan Jurnal", Icon: "DocumentTextIcon", Order: 2, Children: []entity.Menu{
				{Name: "Transaksi Kas & Bank", URL: "/keuangan/transaksi-kas-bank", Icon: "BanknotesIcon", Order: 1},
				{Name: "Jurnal Penyesuaian", URL: "/keuangan/jurnal-penyesuaian", Icon: "PencilIcon", Order: 2},
				{Name: "Jurnal Pembalik", URL: "/keuangan/jurnal-pembalik", Icon: "ArrowUturnLeftIcon", Order: 3},
				{Name: "Jurnal Pembelian", URL: "/keuangan/jurnal-pembelian", Icon: "ShoppingCartIcon", Order: 4},
				{Name: "Jurnal Penjualan (SO)", URL: "/keuangan/jurnal-penjualan", Icon: "TrendingUpIcon", Order: 5},
				{Name: "Jurnal Valuasi", URL: "/keuangan/jurnal-valuasi", Icon: "ChartBarIcon", Order: 6},
			}},
			{Name: "Proses Penutupan", Icon: "LockClosedIcon", Order: 3, Children: []entity.Menu{
				{Name: "Tutup Buku Bulanan", URL: "/keuangan/tutup-buku-bulanan", Icon: "CalendarIcon", Order: 1},
				{Name: "Tutup Buku Keuangan", URL: "/keuangan/tutup-buku-keuangan", Icon: "DocumentCheckIcon", Order: 2},
			}},
			{Name: "Buku Besar", Icon: "BookOpenIcon", Order: 4, Children: []entity.Menu{
				{Name: "Laporan Buku Besar Periodik", URL: "/keuangan/laporan-buku-besar-periodik", Icon: "ClipboardDocumentCheckIcon", Order: 1},
			}},
			{Name: "Gaji & Aset Tetap", Icon: "CurrencyDollarIcon", Order: 5, Children: []entity.Menu{
				{Name: "Data Gaji Pokok", URL: "/keuangan/data-gaji-pokok", Icon: "UserIcon", Order: 1},
				{Name: "Daftar Harta", URL: "/keuangan/daftar-harta", Icon: "BuildingIcon", Order: 2},
				{Name: "Kelompok Harta", URL: "/keuangan/kelompok-harta", Icon: "FolderIcon", Order: 3},
				{Name: "Laporan Daftar Aktiva Tetap", URL: "/keuangan/laporan-aktiva-tetap", Icon: "ClipboardListIcon", Order: 4},
			}},
			{Name: "Piutang & Hutang", Icon: "CreditCardIcon", Order: 6, Children: []entity.Menu{
				{Name: "Piutang Diluar Penjualan", URL: "/keuangan/piutang-diluar-penjualan", Icon: "ArrowRightIcon", Order: 1},
				{Name: "Hutang Diluar Pembelian", URL: "/keuangan/hutang-diluar-pembelian", Icon: "ArrowLeftIcon", Order: 2},
			}},
			{Name: "Lainnya", Icon: "DotsHorizontalIcon", Order: 7, Children: []entity.Menu{
				{Name: "Up Country (Cost)", URL: "/keuangan/up-country-cost", Icon: "GlobeAltIcon", Order: 1},
				{Name: "Budgeting", URL: "/keuangan/budgeting", Icon: "ChartPieIcon", Order: 2},
				{Name: "Budget Bulanan", URL: "/keuangan/budget-bulanan", Icon: "CalendarIcon", Order: 3},
				{Name: "Budget Tahunan", URL: "/keuangan/budget-tahunan", Icon: "ChartBarIcon", Order: 4},
			}},
		}},
		{Name: "HRD", URL: "/hrd", Icon: "UsersIcon", Order: 7, Children: []entity.Menu{
			{Name: "Manajemen Pegawai", Icon: "UserGroupIcon", Order: 1, Children: []entity.Menu{
				{Name: "Data Pegawai", URL: "/hrd/data-pegawai", Icon: "UserIcon", Order: 1},
				{Name: "PTKP Pegawai", URL: "/hrd/ptkp-pegawai", Icon: "IdentificationIcon", Order: 2},
				{Name: "Kontrak Kerja", URL: "/hrd/kontrak-kerja", Icon: "DocumentTextIcon", Order: 3},
				{Name: "Surat Peringatan", URL: "/hrd/surat-peringatan", Icon: "ExclamationCircleIcon", Order: 4},
			}},
			{Name: "Absensi & Kehadiran", Icon: "CalendarIcon", Order: 2, Children: []entity.Menu{
				{Name: "Rekap Presensi", URL: "/hrd/rekap-presensi", Icon: "ClipboardListIcon", Order: 1},
				{Name: "Laporan Kehadiran", URL: "/hrd/laporan-kehadiran", Icon: "ChartBarIcon", Order: 2},
				{Name: "Data Total Kehadiran", URL: "/hrd/data-total-kehadiran", Icon: "ClipboardCheckIcon", Order: 3},
				{Name: "Izin & Cuti", URL: "/hrd/izin-cuti", Icon: "ClockIcon", Order: 4},
			}},
			{Name: "Rekrutmen", Icon: "BriefcaseIcon", Order: 3, Children: []entity.Menu{
				{Name: "Lowongan Kerja", URL: "/hrd/lowongan-kerja", Icon: "DocumentDuplicateIcon", Order: 1},
				{Name: "Data Pelamar", URL: "/hrd/data-pelamar", Icon: "UserCircleIcon", Order: 2},
			}},
			{Name: "Pengembangan SDM", Icon: "LightBulbIcon", Order: 4, Children: []entity.Menu{
				{Name: "Skill & Training", URL: "/hrd/skill-training", Icon: "AcademicCapIcon", Order: 1},
			}},
		}},
		{Name: "Kepegawaian", URL: "/kepegawaian", Icon: "UserCheckIcon", Order: 8, Children: []entity.Menu{
			{Name: "Izin & Cuti Pegawai", URL: "/kepegawaian/izin-cuti", Icon: "ClockIcon", Order: 1},
			{Name: "Data Presensi", URL: "/kepegawaian/data-presensi", Icon: "ClipboardListIcon", Order: 2},
			{Name: "Penggajian Pegawai", URL: "/kepegawaian/penggajian", Icon: "CurrencyDollarIcon", Order: 3},
			{Name: "Slip Gaji", URL: "/kepegawaian/slip-gaji", Icon: "DocumentTextIcon", Order: 4},
			{Name: "Laporan Kegiatan Perusahaan (Daily)", URL: "/kepegawaian/laporan-kegiatan", Icon: "ChartBarIcon", Order: 5},
			{Name: "Dokumen", URL: "/kepegawaian/dokumen", Icon: "DocumentIcon", Order: 6},
			{Name: "Todo List", URL: "/kepegawaian/todo-list", Icon: "CheckCircleIcon", Order: 7},
		}},
		{Name: "Penilaian Pegawai", URL: "/penilaian-pegawai", Icon: "ClipboardListIcon", Order: 9, Children: []entity.Menu{
			{Name: "Aspek Penilaian", URL: "/penilaian-pegawai/aspek-penilaian", Icon: "AdjustmentsHorizontalIcon", Order: 1},
			{Name: "Report Penilaian Kerja", URL: "/penilaian-pegawai/report-penilaian", Icon: "ChartBarIcon", Order: 2},
			{Name: "Daftar Penilaian Pegawai", URL: "/penilaian-pegawai/daftar-penilaian", Icon: "ClipboardCheckIcon", Order: 3},
		}},
		{Name: "Sales Funnel", URL: "/sales-funnel", Icon: "FunnelIcon", Order: 10},
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
