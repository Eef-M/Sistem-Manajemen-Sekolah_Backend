package models

import "time"

type Siswa struct {
	ID           string    `json:"id" gorm:"type:uuid;primaryKey"`
	UserID       string    `json:"user_id" gorm:"type:uuid;not null"`
	Nama         string    `json:"nama" gorm:"type:varchar(255);not null"`
	Kelas        string    `json:"kelas" gorm:"type:varchar(255);not null"`
	TanggalLahir time.Time `json:"tanggal_lahir" gorm:"type:date;not null"`
	Alamat       string    `json:"alamat" gorm:"type:text"`
	NomorInduk   string    `json:"nomor_induk" gorm:"unique;type:varchar(255);unique;not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relasi
	User User `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

func (Siswa) TableName() string {
	return "siswa"
}
