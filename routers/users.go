package routers

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	router.Use(middleware.ValidasiToken())
	router.GET("", controllers.AllUsers)
	router.GET("/:id", controllers.IdUsers)
	router.POST("/", controllers.AddUser)
	router.PATCH("/:id", controllers.EditUser)
	router.DELETE("/:id", controllers.DeleteUser)
}
