package repositories

import (
	"errors"
	"kopikami/models"
	"gorm.io/gorm"
)

// UserRepository mendefinisikan kontrak untuk operasi data pada entitas User
type UserRepository interface {
	Create(user *models.User) error                     // Membuat pengguna baru di database
	FindByEmail(email string) (*models.User, error)    // Mencari pengguna berdasarkan email
}

// userRepository adalah implementasi dari UserRepository
// Menggunakan GORM sebagai ORM untuk mengakses database
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// Create menambahkan data pengguna baru ke dalam database
func (r *userRepository) Create(user *models.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FindByEmail mencari data pengguna berdasarkan email yang diberikan
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("user not found")
	}
	return &user, err
}
