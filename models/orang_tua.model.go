package models

import "time"

type OrangTua struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	UserID    string    `json:"user_id" gorm:"type:uuid;not null"`
	SiswaID   string    `json:"siswa_id" gorm:"type:uuid;not null"`
	Nama      string    `json:"nama" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relasi
	User  User  `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Siswa Siswa `json:"siswa" gorm:"foreignKey:SiswaID;references:ID"`
}

func (OrangTua) TableName() string {
	return "orang_tua"
}
