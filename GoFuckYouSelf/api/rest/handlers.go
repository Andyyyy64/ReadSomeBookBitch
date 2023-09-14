package rest

import (
	"strconv"

	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/internal/database"
	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/internal/models"
	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/pkg/auth"
	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/pkg/books"
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

func LoginUser(c *gin.Context) {
	var loginInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginInfo); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid input data",
		})
		return
	}

	db := database.ConnectDB()
	user, token, err := auth.LoginUser(db, loginInfo.Email, loginInfo.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Authentication failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"user":    user,
		"token":   token,
	})
}

func AddBook(c *gin.Context) {
	var bookDetails models.Books
	var categoryStr string

	c.BindJSON(&bookDetails)

	categoryStr = c.Query("category_id")
	// convert string to int
	categoryID, err := strconv.Atoi(categoryStr)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid category ID",
		})
		return
	}

	token := c.Request.Header.Get("Authorization")

	user, err := auth.GetUserFromToken(token)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid token",
		})
		return
	}
	if err := books.AddBook(*user, bookDetails, categoryID); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to add book",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Book added successfully",
	})

}

func AddCategory(c *gin.Context) {
	var categoryDetails models.Category

	c.BindJSON(&categoryDetails)

	token := c.Request.Header.Get("Authorization")

	user, err := auth.GetUserFromToken(token)
	if err != nil {
		c.JSON(400, gin.H{
			"message" : "Invalid token",
		})
		return
	}
	if err := books.AddCategory(*user, categoryDetails); err != nil {
		return c.JSON(500, gin.H{
			"message" : "Failed to add category",
		})
		return
	}

	c.JSON(200, gin.H{
		"message" : "Category added successfully",
	})
}

func GetAllUserInfo(c *gin.Context) {
	token := c.Requse.Header.Get("Authorization")

	user, err := auth.GetUserFromToken(token)
	if err != nil {
		c.JSON(400, gin.H{
			"message" : "Invaild token",
		})
		return
	}
	if err := books.GetAllUserInfo(*user); err != nil {
		return c.JSON(500, gin.H{
			"message" : "Failed to get you",
		})
		return
	}

	c.JSON(200, gin.H{
		"message" : "successfully get you!"
	})
}
