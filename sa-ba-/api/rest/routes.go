package rest

import (
	"github.com/gin-gonic/gin"
	"sa-ba-/internal/database"
	"sa-ba-/api/rest/handlers"
)

func Routes(router *gin.Engine) {
	router.POST("/register", handlers.RegisterUser)
}