package main

import (
	"backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.GET("/movies", controllers.Movies)
	// router.GET("/users", controllers.Users)

	routers.Routers(router)

	router.Run(":8888")
}
