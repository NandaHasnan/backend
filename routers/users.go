package routers

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	router.Use(middleware.ValidasiToken())
	router.GET("", controllers.AllUsersDB)
	// router.GET("/:id", controllers.IdUsers)
	router.POST("/add_User", controllers.InsertUser)
	// router.PATCH("/:id", controllers.EditUser)
	router.PATCH("", controllers.EditUser)
	router.PATCH("/edit/:id", controllers.EditUserAdmin)
	router.DELETE("", controllers.DeleteUser)
	router.DELETE("/delete/:id", controllers.DeleteUserAdmin)
	router.GET("/detail/:id", controllers.UserProfile2)
}
