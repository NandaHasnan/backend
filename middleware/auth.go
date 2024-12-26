package middleware

import (
	lib "backend/lib"
	"fmt"
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

		// out := jwt.Claims{}
		out := make(map[string]interface{})

		// godotenv.Load()
		// var JWT_SECRET []byte = []byte(lib.Getmd5(os.Getenv("JWT_SECRET")))

		// err := tok.Claims(JWT_SECRET, &out)
		err := tok.Claims(lib.SECRET, &out)

		fmt.Println(out["userid"])

		ctx.Set("userid", out["userid"].(float64))

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
