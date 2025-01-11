package config

import (
	"fmt"
	"log"
	"os"
	"kopikami/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func SetupDatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// ✅ Menjalankan migrasi otomatis
	database.AutoMigrate(&models.Product{})
	log.Println("Database migrated successfully!")

	return database
}

func CloseDatabaseConnection(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
