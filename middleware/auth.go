package middleware

import (
	lib "backend/lib"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ValidasiToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		head := ctx.GetHeader("Authorization")
		if head == "" {
			ctx.JSON(http.StatusUnauthorized, Response{
				Success: false,
				Message: "Unauthorized",
			})
			ctx.Abort()
			return
		}
		token := strings.Split(head, " ")[1:][0]

		tok, _ := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})

		out := jwt.Claims{}

		err := tok.Claims(lib.SECRET, &out)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, Response{
				Success: false,
				Message: "Unauthorized",
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
