package models

import "time"

type Pengumuman struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	Judul     string    `json:"judul" gorm:"type:varchar(255);not null"`
	Konten    string    `json:"konten" gorm:"type:text;not null"`
	Tanggal   time.Time `json:"tanggal" gorm:"type:date;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Pengumuman) TableName() string {
	return "pengumuman"
}
