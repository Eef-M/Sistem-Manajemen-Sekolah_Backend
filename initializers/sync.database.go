package initializers

import "sistem-manajemen-sekolah/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
