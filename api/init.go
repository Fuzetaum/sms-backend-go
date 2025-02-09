package api

import (
	"sms-backend-go/api/rest/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	ginEngine := gin.Default()
	controller.InitHello(ginEngine)
	ginEngine.Run()
}
