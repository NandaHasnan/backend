package routers

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func OrderMovie(router *gin.RouterGroup) {
	router.Use(middleware.ValidasiToken())
	router.POST("", controllers.OrderMoviesNew)
	// router.POST("", controllers.OrderMovies)
	router.GET("/cinema", controllers.CinemaFilter)
	router.POST("/payment", controllers.Payment)
}
