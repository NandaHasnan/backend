package controllers

import (
	lib "backend/lib"
	"backend/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// var users models.ListUsers = models.ListUsers{
// 	{
// 		Id:       1,
// 		Fullname: "fazz",
// 		Email:    "fazz@mail.com",
// 		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$FOMYAEWsUSM4Cmj2YNQ8Xyoa6fw11J/EysRQp2/RIMw", //  "1234"
// 	},
// 	{
// 		Id:       2,
// 		Fullname: "track",
// 		Email:    "track@mail.com",
// 		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$RPeB4d/todvjVR4QFvP3qOK2cAHpAjDZspXeEKMaJvU", //  "5678"
// 	},
// 	{
// 		Id:       3,
// 		Fullname: "endra",
// 		Email:    "endra@mail.com",
// 		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$FOMYAEWsUSM4Cmj2YNQ8Xyoa6fw11J/EysRQp2/RIMw", //  "1234"
// 	},
// 	{
// 		Id:       4,
// 		Fullname: "adiv",
// 		Email:    "adiv@mail.com",
// 		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$RPeB4d/todvjVR4QFvP3qOK2cAHpAjDZspXeEKMaJvU", //  "5678"
// 	},
// 	{
// 		Id:       5,
// 		Fullname: "rinaldi",
// 		Email:    "rinaldi@mail.com",
// 		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$FOMYAEWsUSM4Cmj2YNQ8Xyoa6fw11J/EysRQp2/RIMw", //  "1234"
// 	},
// 	{
// 		Id:       6,
// 		Fullname: "rama",
// 		Email:    "rama@mail.com",
// 		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$RPeB4d/todvjVR4QFvP3qOK2cAHpAjDZspXeEKMaJvU", //  "5678"
// 	},
// 	{
// 		Id:       7,
// 		Fullname: "alwi",
// 		Email:    "alwi@mail.com",
// 		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$FOMYAEWsUSM4Cmj2YNQ8Xyoa6fw11J/EysRQp2/RIMw", //  "1234"
// 	},
// 	{
// 		Id:       8,
// 		Fullname: "joko",
// 		Email:    "joko@mail.com",
// 		Password: "$argon2id$v=19$m=65536,t=1,p=2$Zm9vYmFyYmF6$RPeB4d/todvjVR4QFvP3qOK2cAHpAjDZspXeEKMaJvU", //  "5678"
// 	},
// }

// func AllUsers(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, TaskResponse{
// 		Success: true,
// 		Message: "all user",
// 		Result:  users,
// 	})

// }

// func IdUsers(ctx *gin.Context) {

// 	id, err := strconv.Atoi(ctx.Param("id"))
// 	user := users
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, TaskResponse{
// 			Success: false,
// 			Message: "invalid all user",
// 			Result:  id,
// 		})
// 		return
// 	}

// 	var Users models.User

// 	for _, data := range user {
// 		if data.Id == id {
// 			Users = data
// 		}
// 	}

// 	if Users == (models.User{}) {
// 		ctx.JSON(http.StatusNotFound, task{
// 			Success: false,
// 			Message: "all image not found",
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, task{
// 		Success: true,
// 		Message: "all users",
// 		Result:  Users,
// 	})

// }

func EditUser(ctx *gin.Context) {
	val, isAvail := ctx.Get("userid")
	if !isAvail {
		ctx.JSON(http.StatusUnauthorized, TaskResponse2{
			Success: false,
			Message: "Unauthorized",
		})
		return
	}

	userId := int(val.(float64))

	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: "No file provided",
		})
		return
	}

	maxFile := 2 * 1024 * 1024
	if file.Size > int64(maxFile) {
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: "File size is too large",
		})
		return
	}

	var fileName string
	if file.Filename != "" {
		splitFile := strings.Split(file.Filename, ".")
		if len(splitFile) < 2 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid file format",
			})
			return
		}

		fileExt := strings.ToLower(splitFile[len(splitFile)-1])

		allowedExtensions := map[string]bool{
			"jpg": true,
			"png": true,
		}

		if !allowedExtensions[fileExt] {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Only .jpg and .png files are allowed",
			})
			return
		}

		fileName = uuid.New().String()
		filePath := fmt.Sprintf("upload/users/%s.%s", fileName, fileExt)

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, TaskResponse2{
				Success: false,
				Message: "Error saving the file",
			})
			return
		}
	}

	firstname := ctx.PostForm("firstname")
	lastname := ctx.PostForm("lastname")
	phone_number := ctx.PostForm("phone_number")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	if password != "" && !strings.Contains(password, "$argon2i$v=19$m=65536,t=1,p=2$") {
		password = lib.GenerateHash(password)
	}

	updatedUser := models.Gabung{
		Id:           &userId,
		Firstname:    firstname,
		Lastname:     lastname,
		Phone_number: phone_number,
		Image:        fileName + "." + strings.Split(file.Filename, ".")[1],
		Email:        email,
		Password:     password,
	}

	result := models.UpdateUser(updatedUser)

	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "User updated successfully",
		Result:  result,
	})
}

// func EditUser(ctx *gin.Context) {

// 	// iddb, _ := strconv.Atoi(ctx.Param("id"))
// 	// user := models.UserById(iddb)
// 	// if user == (models.Gabung{}) {
// 	// 	ctx.JSON(http.StatusBadRequest, TaskResponse{
// 	// 		Success: false,
// 	// 		Message: "invalid add user",
// 	// 		Result:  iddb,
// 	// 	})
// 	// 	return
// 	// }

// 	// ctx.ShouldBind(&user)

// 	val, isAvail := ctx.Get("userid")
// 	if !isAvail {
// 		ctx.JSON(http.StatusUnauthorized, TaskResponse2{
// 			Success: false,
// 			Message: "Unauthorized",
// 		})
// 		return
// 	}
// 	userId := models.UserById(int(val.(float64)))

// 	file, err := ctx.FormFile("image")
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, TaskResponse2{
// 			Success: false,
// 			Message: "No file provided",
// 		})
// 		return
// 	}

// 	maxFile := 2 * 1024 * 1024
// 	if file.Size > int64(maxFile) {
// 		ctx.JSON(http.StatusBadRequest, TaskResponse2{
// 			Success: false,
// 			Message: "File size is too large",
// 		})
// 		return
// 	}

// 	var fileName string
// 	if file.Filename != "" {

// 		splitFile := strings.Split(file.Filename, ".")
// 		if len(splitFile) < 2 {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Invalid file format",
// 			})
// 			return
// 		}

// 		fileExt := strings.ToLower(splitFile[len(splitFile)-1])

// 		allowedExtensions := map[string]bool{
// 			"jpg": true,
// 			"png": true,
// 		}

// 		if !allowedExtensions[fileExt] {
// 			ctx.JSON(http.StatusBadRequest, gin.H{
// 				"error": "Only .jpg and .png files are allowed",
// 			})
// 			return
// 		}

// 		fileName = uuid.New().String()
// 		filePath := fmt.Sprintf("upload/users/%s.%s", fileName, fileExt)

// 		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
// 			ctx.JSON(http.StatusInternalServerError, TaskResponse2{
// 				Success: false,
// 				Message: "Error saving the file",
// 			})
// 			return
// 		}
// 	}

// 	if !strings.Contains(userId.Password, "$argon2i$v=19$m=65536,t=1,p=2$") {
// 		if userId.Password != "" {
// 			userId.Password = lib.GenerateHash(userId.Password)
// 		}
// 	}

// 	UpdateUser := models.UpdateUser(userId)

// 	ctx.JSON(http.StatusOK, TaskResponse2{
// 		Success: true,
// 		Message: "Update User sukses",
// 		Result:  UpdateUser,
// 	})

// }

func DeleteUser(ctx *gin.Context) {
	iddb, _ := strconv.Atoi(ctx.Param("id"))
	user := models.UserById(iddb)
	if user == (models.Gabung{}) {
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: "invalid add user",
			Result:  iddb,
		})
		return
	}

	// ctx.ShouldBind(&user)

	DeleteUser, err := models.DeleteUser(iddb)
	if err != nil {
		log.Println(err)
	}

	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "Delete User sukses",
		Result:  DeleteUser,
	})

}

// func UsersId(ctx *gin.Context) {
// 	iddb, err := strconv.Atoi(ctx.Param("id"))
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, TaskResponse2{
// 			Success: false,
// 			Message: "invalid user",
// 		})
// 		return
// 	}
// 	user := models.UserById(iddb)
// 	// if user == nil {
// 	// 	ctx.JSON(http.StatusNotFound, TaskResponse2{
// 	// 		Success: false,
// 	// 		Message: "id tidak ada",
// 	// 	})
// 	// 	return
// 	// }

// 	ctx.JSON(http.StatusOK, TaskResponse2{
// 		Success: true,
// 		Message: "User Id",
// 		Result:  user,
// 	})

// }

func UserProfile(ctx *gin.Context) {
	val, isAvail := ctx.Get("userid")
	userId := int(val.(float64))

	profile := models.UserById(userId)

	if isAvail {
		ctx.JSON(http.StatusOK, TaskResponse2{
			Success: true,
			Message: "User Id",
			Result:  profile,
		})
	}
}

func AllUsersDB(ctx *gin.Context) {
	user := models.UserAll()
	// fmt.Println(user)
	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "All User",
		Result:  user,
	})

}

// router.Run("localhost:8888")
