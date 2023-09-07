package rest

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/register", RegisterUser)
}
