package routers

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func MovieRouter(router *gin.RouterGroup) {
	router.GET("", controllers.AllMovieDB)
	router.GET("/:id", controllers.IdMovies)
}

func MovieRouter2(router *gin.RouterGroup) {
	router.Use(middleware.ValidasiToken())
	router.POST("/addmovie", controllers.AddMovies)
	router.POST("/order", controllers.OrderMovies)
	router.PATCH("/:id", controllers.EditMovies)
	router.DELETE("/:id", controllers.DeleteMovies)
}
