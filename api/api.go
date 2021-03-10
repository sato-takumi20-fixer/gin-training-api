package api

import "github.com/gin-gonic/gin"

func Declare(router *gin.Engine) {
	DeclareCalcApi(router)
}