package rest

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/register", RegisterUser)
	router.POST("/login", LoginUser)
	router.POST("/add-book", AddBook)
}
