package routers

import "github.com/gin-gonic/gin"

func Routers(router *gin.Engine) {

	MovieRouter(router.Group("/movies"))
	MovieRouter2(router.Group("/movies"))
	UsersRouter(router.Group("/users"))
	LoginRouter(router.Group("/auth/login"))
	RegisterRouter(router.Group("/auth/register"))
	OrderMovie(router.Group("/order"))
	Profile(router.Group("/profile"))

}
