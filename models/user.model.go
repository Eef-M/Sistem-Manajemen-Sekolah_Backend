package models

import "time"

type Roles string

const (
	RoleAdmin    Roles = "admin"
	RoleSiswa    Roles = "siswa"
	RoleGuru     Roles = "guru"
	RoleOrangTua Roles = "orang tua"
)

type User struct {
	ID        string    `json:"id" gorm:"type:uuid;primaryKey"`
	Username  string    `json:"username" gorm:"type:varchar(255);not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	Role      Roles     `json:"role" gorm:"type:roles;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
