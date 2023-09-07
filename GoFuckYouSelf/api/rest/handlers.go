package rest

import (
	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/internal/database"
	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/internal/models"
	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/pkg/auth"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	var db = database.ConnectDB()

	c.BindJSON(&user)

	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to hash password!",
		})
		return
	}
	user.Password = hashedPassword

	if err := db.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create user!",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully!",
		"user":    user,
	})
}
