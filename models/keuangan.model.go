package models

import "time"

type StatusPembayaran string

const (
	Lunas      StatusPembayaran = "lunas"
	BelumLunas StatusPembayaran = "belum lunas"
)

type Keuangan struct {
	ID              string    `json:"id" gorm:"type:uuid;primaryKey"`
	SiswaID         string    `json:"siswa_id" gorm:"type:uuid;not null"`
	JenisPembayaran string    `json:"jenis_pembayaran" gorm:"type:varchar(255);not null"`
	Status          string    `json:"status" gorm:"type:status_pembayaran;not null"`
	Jumlah          int       `json:"jumlah" gorm:"not null"`
	Tanggal         time.Time `json:"tanggal" gorm:"type:date;not null"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	// Relasi
	Siswa Siswa `json:"siswa" gorm:"foreignKey:SiswaID;references:ID"`
}

func (Keuangan) TableName() string {
	return "keuangan"
}
