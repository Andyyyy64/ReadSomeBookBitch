package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
