package routers

import (
	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
	// 	AllowHeaders:     []string{"Authorization", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))
	router.Use(func(ctx *gin.Context) {

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type, Accept")
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin")
		ctx.Header("Access-Control-Allow-Credentials", "true")

		ctx.Next()
	})

	router.Use(func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			ctx.Header("Access-Control-Allow-Origin", "http://localhost:5173")
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	})

	MovieRouter(router.Group("/movies"))
	MovieRouter2(router.Group("/movies"))
	UsersRouter(router.Group("/users"))
	LoginRouter(router.Group("/auth/login"))
	RegisterRouter(router.Group("/auth/register"))
	// OrderMovie(router.Group("/order"))
	OrderMovie(router.Group("/order"))
	Profile(router.Group("/profile"))

}
