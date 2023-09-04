package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/Andyyyy64/ReadSomeBookBitch/sa-ba-/internal/database"
	"github.com/Andyyyy64/ReadSomeBookBitch/sa-ba-/api/rest/handlers"
)

func Routes(router *gin.Engine) {
	router.POST("/register", handlers.RegisterUser)
}