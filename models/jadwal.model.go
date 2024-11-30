package models

import "time"

type Hari string

const (
	Senin  Hari = "senin"
	Selasa Hari = "selasa"
	Rabu   Hari = "rabu"
	Kamis  Hari = "kamis"
	Jumat  Hari = "jumat"
	Sabtu  Hari = "sabtu"
)

type Jadwal struct {
	ID            string    `json:"id" gorm:"type:uuid;primary_key"`
	KelasID       string    `json:"kelas_id" gorm:"type:uuid;not null"`
	GuruID        string    `json:"guru_id" gorm:"type:uuid;not null"`
	MataPelajaran string    `json:"mata_pelajaran" gorm:"type:varchar(255);not null"`
	Hari          Hari      `json:"hari" gorm:"type:hari;not null"`
	WaktuMulai    time.Time `json:"waktu_mulai" gorm:"not null"`
	WaktuSelesai  time.Time `json:"waktu_selesai" gorm:"not null"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relasi
	Kelas Kelas `json:"kelas" gorm:"foreignKey:KelasID;references:ID"`
	Guru  Guru  `json:"guru" gorm:"foreignKey:GuruID;refrences:ID"`
}

func (Jadwal) TableName() string {
	return "jadwal"
}
