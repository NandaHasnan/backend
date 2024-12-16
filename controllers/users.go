package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AllUsers(ctx *gin.Context) {

	// var login User
	// ctx.ShouldBind(&login)

	// var user *User
	// for i := range users {
	// 	ctx.Header("X-Logged-User", "true")
	// 	if users[i].Email == login.Email {
	// 		user = &users[i]
	// 		break
	// 	}
	// }

	// if user != nil {
	// 	ctx.JSON(401, TaskResponse{
	// 		Success: false,
	// 		Message: "Unauthorized",
	// 	})
	// 	return
	// }

	// hasher, _ := argon2.CreateHash(login.Password, login.Password, argon2.DefaultParams)
	// var passUser User
	// hasher := lib.GenerateHash(passUser.Password)
	// if hasher == passUser.Email {
	// 	ctx.JSON(http.StatusUnauthorized, TaskResponse{
	// 		Success: false,
	// 		Message: "Invalid password",
	// 	})
	// 	return
	// }

	ctx.JSON(http.StatusOK, TaskResponse{
		Success: true,
		Message: "all user",
		Result:  users,
	})

}

func IdUsers(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	user := users
	if err != nil {
		ctx.JSON(http.StatusBadRequest, TaskResponse{
			Success: false,
			Message: "invalid all user",
			Result:  id,
		})
		return
	}

	var Users User

	for _, data := range user {
		if data.Id == id {
			Users = data
		}
	}

	if Users == (User{}) {
		ctx.JSON(http.StatusNotFound, task{
			Success: false,
			Message: "all image not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "all users",
		Result:  Users,
	})

}

func AddUser(ctx *gin.Context) {

	var newUser User

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.Id = len(users) + 1
	users = append(users, newUser)

	ctx.JSON(http.StatusOK, TaskResponse{
		Success: true,
		Message: "add users sukses",
		Result:  newUser,
	})

}

func EditUser(ctx *gin.Context) {

	id, err := strconv.Atoi(ctx.Param("id"))
	user := users
	if err != nil {
		ctx.JSON(http.StatusBadRequest, TaskResponse{
			Success: false,
			Message: "invalid add user",
			Result:  id,
		})
		return
	}

	var updateUser User
	if err := ctx.ShouldBind(&updateUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, users := range user {
		if users.Id == id {
			if updateUser.Fullname != "" {
				user[i].Fullname = updateUser.Fullname
			}
			if updateUser.Email != "" {
				user[i].Email = updateUser.Email
			}
			if updateUser.Password != "" {
				user[i].Password = users.Password
			}
		}

		ctx.JSON(http.StatusOK, TaskResponse{
			Success: true,
			Message: "Update user sukses",
			Result:  user[i],
		})
	}

}

func DeleteUser(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	user := users
	if err != nil {
		ctx.JSON(http.StatusBadRequest, TaskResponse{
			Success: false,
			Message: "invalid user delete",
			Result:  id,
		})
		return
	}

	for i, Users := range user {
		if Users.Id == id {
			user = append(user[:i], user[i+1:]...)
		}
	}

	ctx.JSON(http.StatusOK, TaskResponse{
		Success: true,
		Message: "Delete user sukses",
		Result:  user,
	})

}

// router.Run("localhost:8888")
