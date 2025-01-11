package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);unique;not null" json:"name" binding:"required"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email" binding:"required,email"`
	Password  string    `gorm:"not null" json:"password" binding:"required,min=6"`
	Role      string    `gorm:"type:enum('admin','cashier');not null" json:"role" binding:"required,oneof=admin cashier"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
