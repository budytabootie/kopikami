package repositories

import (
	"errors"
	"kopikami/models"
	"gorm.io/gorm"
)

// UserRepository mendefinisikan kontrak untuk operasi data pada entitas User
type UserRepository interface {
	Create(user *models.User) error                     // Membuat pengguna baru di database
	FindByEmail(email string) (*models.User, error)     // Mencari pengguna berdasarkan email
	FindByID(userID uint) (*models.User, error)         // Mencari pengguna berdasarkan ID
	FindAll() ([]models.User, error)                    // Mengambil semua pengguna
	Update(user *models.User) error                     // Memperbarui data pengguna
	Delete(userID uint) error                           // Menghapus pengguna berdasarkan ID
}

// userRepository adalah implementasi dari UserRepository
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository membuat instance baru dari userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// Create menambahkan data pengguna baru ke dalam database
func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
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

// FindByID mencari data pengguna berdasarkan ID
func (r *userRepository) FindByID(userID uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, userID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("user not found")
	}
	return &user, err
}

// FindAll mengambil semua pengguna dari database
func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

// Update memperbarui data pengguna di database
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete menghapus pengguna berdasarkan ID dari database
func (r *userRepository) Delete(userID uint) error {
	result := r.db.Delete(&models.User{}, userID)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return result.Error
}
