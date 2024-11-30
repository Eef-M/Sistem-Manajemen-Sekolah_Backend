package models

import "time"

type Kelas struct {
	ID          string    `json:"id" gorm:"type:uuid;primary_key"`
	NamaKelas   string    `json:"nama_kelas" gorm:"type:varchar(255);not null"`
	WaliKelasID string    `json:"wali_kelas_id" gorm:"type:uuid;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relasi
	Guru Guru `json:"guru" gorm:"foreignKey:WaliKelasID;references:ID"`
}

func (Kelas) TableName() string {
	return "kelas"
}
