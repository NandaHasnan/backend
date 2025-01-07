package routers

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func OrderMovie(router *gin.RouterGroup) {
	router.Use(middleware.ValidasiToken())
	router.POST("", controllers.OrderMovies)
	router.GET("/cinema", controllers.CinemaFilter)
}
