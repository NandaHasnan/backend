package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pilinux/argon2"
)

type TaskResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type User struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User = []User{
	{
		Id:       1,
		Fullname: "fazz",
		Email:    "fazz@mail.com",
		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$FOMYAEWsUSM4Cmj2YNQ8Xyoa6fw11J/EysRQp2/RIMw", //  "1234"
	},
	{
		Id:       2,
		Fullname: "track",
		Email:    "track@mail.com",
		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$RPeB4d/todvjVR4QFvP3qOK2cAHpAjDZspXeEKMaJvU", //  "5678"
	},
}

func Users(*gin.Context) {
	router := gin.Default()

	router.GET("/auth", func(c *gin.Context) {
		var login User

		hasher, _ := argon2.CreateHash(login.Password, login.Password, argon2.DefaultParams)

		c.JSON(http.StatusOK, TaskResponse{
			Success: true,
			Message: "all user",
			Result:  hasher,
		})
	})

	router.POST("/auth/login", func(ctx *gin.Context) {
		var login User

		ctx.ShouldBind(&login)

		var user *User
		for i := range users {
			if users[i].Email == login.Email {
				user = &users[i]
				break
			}
		}

		if user == nil {
			ctx.JSON(http.StatusNotFound, TaskResponse{
				Success: false,
				Message: "Email not found",
			})
			return
		}

		match, _ := argon2.ComparePasswordAndHash(login.Password, login.Password, user.Password)
		if match {
			ctx.JSON(http.StatusUnauthorized, TaskResponse{
				Success: false,
				Message: "Invalid password",
			})
			return
		}

		ctx.JSON(http.StatusOK, TaskResponse{
			Success: true,
			Message: "Login successful",
			Result:  user,
		})
	})

	router.Run("localhost:8888")
}
