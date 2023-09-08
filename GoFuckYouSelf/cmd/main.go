package main

import (
	"time"

	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/api/rest"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowCredentials: false,
		MaxAge:           24 * time.Hour,
	}))
	rest.Routes(r)

	r.Run(":8080")
}
