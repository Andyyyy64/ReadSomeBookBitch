package auth

import (
	"fmt"
	"os"
	"time"
	"strings"

	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/internal/database"
	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/internal/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJwt(user models.User) (string, error) {
	expire := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   user.Username,
		ExpiresAt: expire.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return token, nil
}

func GetUserFromToken(tokenStr string) (*models.User, error) {

	if len(tokenStr) > 7 && strings.ToUpper(tokenStr[:7]) == "BEARER " {
		tokenStr = tokenStr[7:]
	}

	token, err := ValidateToken(tokenStr)
	fmt.Println("token")
	fmt.Println(token)
	fmt.Println(tokenStr)
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	db := database.ConnectDB()
	if err := db.Where("username = ?", claims.Subject).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserFromUserId(userId uint) (*models.User, error) {
	var user models.User
	db := database.ConnectDB()
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func LoginUser(db *gorm.DB, email, password string) (*models.User, string, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, "", err
	}
	if !CheckPasswordHash(password, user.Password) {
		return nil, "Invalid password", db.Error
	}
	token, err := GenerateJwt(user)
	if err != nil {
		return nil, "", err
	}
	return &user, token, nil
}
