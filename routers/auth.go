package routers

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func LoginRouter(router *gin.RouterGroup) {
	router.POST("", controllers.AuthLogin)
}

func RegisterRouter(router *gin.RouterGroup) {
	router.POST("", controllers.Register)
}
