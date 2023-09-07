package main

import (
	"github.com/Andyyyy64/ReadSomeBookBitch/GoFuckYouSelf/api/rest"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()

	rest.Routes(r)

	r.Run(":8080")
}
