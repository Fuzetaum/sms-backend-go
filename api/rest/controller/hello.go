package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func InitHello(ginEngine *gin.Engine) {
	ginEngine.GET("/hello", helloWorld)
}
