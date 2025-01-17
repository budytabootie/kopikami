package models

import "time"

// User merepresentasikan data pengguna dalam sistem
// Termasuk informasi autentikasi seperti email, password, dan role

type User struct {
	ID        uint      `gorm:"primaryKey"`                          // Primary key untuk user
	Name      string    `gorm:"type:varchar(255);unique;not null" json:"name" binding:"required"` // Nama pengguna, wajib diisi dan unik
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email" binding:"required,email"` // Email pengguna, wajib diisi dan harus unik
	Password  string    `gorm:"not null" json:"password" binding:"required,min=6"`         // Password pengguna dengan validasi minimal 6 karakter
	Role      string    `gorm:"type:enum('admin','cashier');not null" json:"role" binding:"required,oneof=admin cashier"` // Role pengguna, hanya bisa 'admin' atau 'cashier'
	CreatedAt time.Time `json:"created_at"`                         // Waktu user dibuat
	UpdatedAt time.Time `json:"updated_at"`                         // Waktu terakhir data user diperbarui
}
