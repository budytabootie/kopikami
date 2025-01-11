package services

import (
	"errors"
	"kopikami/models"
	"kopikami/repositories"
	"kopikami/utils"

	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=admin cashier"`
}

type AuthService interface {
	Register(input RegisterInput) (*models.User, error)
	Login(input LoginInput) (string, error)
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{userRepo}
}

// ✅ Fungsi Register Ditambahkan (Perbaikan)
func (s *authService) Register(input RegisterInput) (*models.User, error) {
	// Cek apakah email sudah terdaftar
	existingUser, _ := s.userRepo.FindByEmail(input.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Membuat user baru
	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	// Menyimpan user ke database
	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// ✅ Fungsi Login Sudah Ada
func (s *authService) Login(input LoginInput) (string, error) {
	user, err := s.userRepo.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Verifikasi Password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate Token JWT
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
