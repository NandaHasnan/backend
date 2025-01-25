package main

import (
	"backend/docs"
	"backend/routers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// docs "BACKEND/docs"
)

// @title Backend API
// @version 1.0
// @description BackEnd API

// @BasePath /

// @securityDefinitions.apikey  ApiKeyAuth
// @in 							header
// @name 						Authorization

func main() {
	router := gin.Default()

	// router.Use(cors.New(cors.Config{
	// 	AllowAllOrigins: true,
	// 	AllowHeaders:    []string{"Authorization"},
	// }))

	router.GET("/", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// router.GET("/movies", controllers.Movies)
	// router.GET("/users", controllers.Users)
	// router.MaxMultipartMemory = 2 << 20
	router.Static("/movies/image", "./upload/movies")
	router.Static("/users/image", "./upload/users")

	routers.Routers(router)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8888")
}
