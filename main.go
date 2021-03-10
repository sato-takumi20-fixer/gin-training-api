package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sato-takumi20-fixer/gin-training-api/api"
)

func main() {
	router := gin.Default()
	api.Declare(router)		
	router.Run()
}
