package initializers

import (
	"log"
	"sistem-manajemen-sekolah/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Siswa{},
		&models.Guru{},
		&models.OrangTua{},
		&models.Absensi{},
		&models.AbsensiDetail{},
		&models.Jadwal{},
		&models.Kelas{},
		&models.Keuangan{},
		&models.Pengumuman{},
	)

	if err != nil {
		log.Fatalf("Error migrating tables: %v", err)
	}
}
