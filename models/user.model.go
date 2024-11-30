package models

import "time"

type Roles string

const (
	Admin    Roles = "admin"
	Siswa    Roles = "siswa"
	Guru     Roles = "guru"
	OrangTua Roles = "orang_tua"
)

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primary_key"`
	Username  string    `json:"username" gorm:"type:varchar(255);not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Role      Roles     `json:"role" gorm:"type:user_role;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
