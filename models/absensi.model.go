package models

import "time"

type StatusKehadiran string

const (
	Hadir StatusKehadiran = "hadir"
	Sakit StatusKehadiran = "sakit"
	Izin  StatusKehadiran = "izin"
	Alfa  StatusKehadiran = "alfa"
)

type Absensi struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	KelasID   string    `json:"kelas_id" gorm:"type:uuid;not null"`
	Tanggal   time.Time `'json:"tanggal" gorm:"type:date;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relasi
	Kelas Kelas `json:"kelas" gorm:"foreignKey:KelasID;references:ID"`
}

type AbsensiDetail struct {
	ID        string          `json:"id" gorm:"type:uuid;primaryKey"`
	AbsensiID string          `json:"absensi_id" gorm:"type:uuid;not null"`
	Status    StatusKehadiran `json:"status" gorm:"type:status_kehadiran;not null"`
	CreatedAt time.Time       `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"autoUpdateTime"`

	// Relasi
	Absensi Absensi `json:"absensi" gorm:"foreignKey:AbsensiID;references:ID"`
}

func (Absensi) TableName() string {
	return "absensi"
}
