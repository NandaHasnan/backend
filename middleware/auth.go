package middleware

import (
	"backend/controllers"
	lib "backend/lib"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

func ValidasiToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		head := ctx.GetHeader("Authorization")
		// fmt.Println(head)
		token := strings.Split(head, " ")[1:][0]

		tok, _ := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})

		out := jwt.Claims{}

		err := tok.Claims(lib.SECRET, &out)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, controllers.TaskResponse{
				Success: false,
				Message: "Unauthorized",
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
