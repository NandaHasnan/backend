package routers

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func MovieRouter(router *gin.RouterGroup) {

	router.GET("", controllers.AllMovies)
	router.GET("/:id", controllers.IdMovies)
	router.GET("/search", controllers.SearchMovies)
	router.GET("/paging", controllers.Paging)
}

func MovieRouter2(router *gin.RouterGroup) {
	router.Use(middleware.ValidasiToken())
	router.POST("/", controllers.AddMovies)
	router.PATCH("/:id", controllers.EditMovies)
	router.DELETE("/:id", controllers.DeleteMovies)
}
