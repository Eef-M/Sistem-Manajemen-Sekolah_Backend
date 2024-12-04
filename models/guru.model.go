package models

import "time"

type Guru struct {
	ID            string    `json:"id" gorm:"type:uuid;primaryKey"`
	UserID        string    `json:"user_id" gorm:"type:uuid;not null"`
	Nama          string    `json:"nama" gorm:"type:varchar(255);not null"`
	MataPelajaran string    `json:"mata_pelajaran" gorm:"type:varchar(255);not null"`
	TanggalLahir  time.Time `json:"tanggal_lahir" gorm:"type:date;not null"`
	Alamat        string    `json:"alamat" gorm:"type:text"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relasi
	User User `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

func (Guru) TableName() string {
	return "guru"
}
