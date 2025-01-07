package routers

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func Profile(router *gin.RouterGroup) {
	router.Use(middleware.ValidasiToken())
	router.GET("", controllers.UserProfile)
}
