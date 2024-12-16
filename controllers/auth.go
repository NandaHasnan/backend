package controllers

import (
	lib "backend/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthLogin(ctx *gin.Context) {

	var login User

	ctx.ShouldBind(&login)

	var user User
	for _, i := range users {
		if i.Email == login.Email {
			user = i

		}
	}

	if user == (User{}) {
		ctx.JSON(http.StatusNotFound, TaskResponse{
			Success: false,
			Message: "Email not found",
		})
		return
	}

	// match, _ := argon2.ComparePasswordAndHash(login.Password, login.Password, user.Password)
	if lib.VerifyHash(user.Password, login.Password) {
		ctx.JSON(http.StatusUnauthorized, TaskResponse{
			Success: false,
			Message: "Invalid password",
		})
		return
	}

	token := lib.GeneratePass(struct {
		UserId int `json:"userid"`
	}{
		UserId: login.Id,
	})

	ctx.JSON(http.StatusOK, TaskResponse{
		Success: true,
		Message: "Login successful",
		Result:  token,
	})

}

func Register(ctx *gin.Context) {

	var newUser User

	err1 := ctx.ShouldBind(&newUser)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	err := validatePassword(newUser.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, TaskResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	newUser.Id = len(users) + 1
	users = append(users, newUser)

	hasher := lib.GenerateHash(newUser.Password)
	if hasher == newUser.Email {
		ctx.JSON(http.StatusUnauthorized, TaskResponse{
			Success: false,
			Message: "Invalid password",
		})
		return
	}

	ctx.JSON(http.StatusOK, TaskResponse{
		Success: true,
		Message: "Register sukses",
		Result:  hasher,
	})

}
