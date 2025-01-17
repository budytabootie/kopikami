package config

import (
	"fmt"
	"log"
	"os"
	"kopikami/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDatabaseConnection membuat koneksi ke database MySQL dan mengembalikan objek *gorm.DB
func SetupDatabaseConnection() *gorm.DB {
	// Membuat DSN (Data Source Name) untuk koneksi MySQL dengan environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	// Membuka koneksi dengan GORM
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// ✅ Menjalankan migrasi otomatis untuk tabel Product
	database.AutoMigrate(&models.Product{})
	log.Println("Database migrated successfully!")

	return database
}

// CloseDatabaseConnection menutup koneksi ke database dengan memanggil db.Close()
func CloseDatabaseConnection(db *gorm.DB) {
	sqlDB, _ := db.DB() // Mengambil koneksi SQL dari objek GORM
	sqlDB.Close()       // Menutup koneksi ke database
}
