package models

import "time"

// Product mendefinisikan struktur data untuk produk dalam sistem
// Termasuk informasi dasar seperti nama, harga, stok, dan waktu pembuatan

type Product struct {
	ID        uint      `gorm:"primaryKey"`                          // Primary Key untuk produk
	Name      string    `gorm:"type:varchar(255);unique;not null" json:"name" binding:"required"` // Nama produk, unik dan wajib diisi
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price" binding:"required,gt=0"`  // Harga produk dengan validasi minimal > 0
	Stock     int       `gorm:"not null;default:0" json:"stock" binding:"required,gte=0"`        // Stok produk dengan validasi minimal >= 0
	CreatedAt time.Time `json:"created_at"`                         // Waktu produk dibuat
	UpdatedAt time.Time `json:"updated_at"`                         // Waktu produk diperbarui terakhir kali
}
