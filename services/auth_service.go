package services

import (
	"errors"
	"kopikami/models"
	"kopikami/repositories"
	"kopikami/utils"

	"golang.org/x/crypto/bcrypt"
)

// LoginInput mendefinisikan struktur data untuk input login
// Menggunakan validasi JSON binding untuk memastikan data yang dikirim sesuai

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterInput mendefinisikan struktur data untuk input pendaftaran pengguna
// Terdapat validasi role dan panjang password minimal

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=admin cashier"`
}

// AuthService mendefinisikan kontrak untuk layanan autentikasi
// Meliputi fungsi untuk registrasi dan login pengguna
type AuthService interface {
	Register(input RegisterInput) (*models.User, error)   // Mendaftarkan user baru
	Login(input LoginInput) (string, error)               // Login user dan mengembalikan token JWT
}

// authService adalah implementasi dari AuthService yang menggunakan UserRepository
// untuk melakukan operasi pada data pengguna
type authService struct {
	userRepo repositories.UserRepository
}

// NewAuthService membuat instance baru dari authService
func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

// Register memungkinkan pendaftaran pengguna baru dengan validasi data
func (s *authService) Register(input RegisterInput) (*models.User, error) {
	// Cek apakah email sudah terdaftar
	existingUser, _ := s.userRepo.FindByEmail(input.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hashing password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Membuat user baru dengan data yang sudah tervalidasi
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	// Menyimpan data user ke dalam database
	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Login memproses autentikasi pengguna berdasarkan email dan password
func (s *authService) Login(input LoginInput) (string, error) {
	// Cek apakah email sudah terdaftar
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Verifikasi Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate Token JWT setelah verifikasi berhasil
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
