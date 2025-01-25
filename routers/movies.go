package routers

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func MovieRouter(router *gin.RouterGroup) {
	router.GET("", controllers.AllMovieDB)
	router.GET("/:id", controllers.IdMovies)
	router.GET("/detail/:id", controllers.IdMovies)
}

func MovieRouter2(router *gin.RouterGroup) {
	router.Use(middleware.ValidasiToken())
	router.POST("/addmovie", controllers.AddMovies)
	router.POST("/order", controllers.OrderMovies)
	router.PATCH("/edit/:id", controllers.EditMovies)
	router.DELETE("/delete/:id", controllers.DeleteMovies)
}
