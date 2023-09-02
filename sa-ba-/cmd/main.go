package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Andyyyy64/ReadSomeBookBitch/sa-ba-/api/rest/routes"	
)

func main() {
	r := gin.Default()

	rest.Routes(r)

	router.Run(":8080")
}