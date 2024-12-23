package controllers

import (
	lib "backend/lib"
	"backend/models"
	"fmt"
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
	// Ambil file gambar dari request
	file, err := ctx.FormFile("image")
	var fileName string
	if err == nil && file != nil {
		// Simpan file jika ada
		fileName = uuid.New().String()
		splitFile := strings.Split(file.Filename, ".")
		fileExt := splitFile[len(splitFile)-1]
		filePath := fmt.Sprintf("upload/movies/%s.%s", fileName, fileExt)

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, TaskResponse{
				Success: false,
				Message: "Error saving the file",
				Result:  nil,
			})
			return
		}
	}

	// Ambil data dari form-data
	firstname := ctx.PostForm("firstname")
	lastname := ctx.PostForm("lastname")
	phone_number := ctx.PostForm("phone_number")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	// Validasi email
	if email == "" {
		ctx.JSON(http.StatusBadRequest, TaskResponse{
			Success: false,
			Message: "Email is required",
			Result:  nil,
		})
		return
	}

	// Periksa apakah user ada di database
	user := models.UserByEmail(email)
	if user == (models.User_credentials{}) {
		ctx.JSON(http.StatusNotFound, TaskResponse{
			Success: false,
			Message: "Email not found",
			Result:  nil,
		})
		return
	}

	// Hash password jika ada perubahan
	if password != "" && !strings.Contains(password, "$argon2i$v=19$m=65536,t=1,p=2$") {
		password = lib.GenerateHash(password)
	}

	// Siapkan data untuk diperbarui
	updatedUser := models.Gabung{
		Firstname:    firstname,
		Lastname:     lastname,
		Phone_number: phone_number,
		Image:        fileName + "." + strings.Split(file.Filename, ".")[1],
		Email:        email,
		Password:     password,
	}

	// Update data pengguna di database
	result := models.UpdateUser(updatedUser)

	// Kirim respons berhasil
	ctx.JSON(http.StatusOK, TaskResponse{
		Success: true,
		Message: "User updated successfully",
		Result:  result,
	})
}

// func EditUser(ctx *gin.Context) {

// 	iddb, _ := strconv.Atoi(ctx.Param("id"))
// 	user := models.UserById(iddb)
// 	if user == (models.Gabung{}) {
// 		ctx.JSON(http.StatusBadRequest, TaskResponse{
// 			Success: false,
// 			Message: "invalid add user",
// 			Result:  iddb,
// 		})
// 		return
// 	}

// 	ctx.ShouldBind(&user)

// 	if !strings.Contains(user.Password, "$argon2i$v=19$m=65536,t=1,p=2$") {
// 		if user.Password != "" {
// 			user.Password = lib.GenerateHash(user.Password)
// 		}
// 	}

// 	UpdateUser := models.UpdateUser(user)

// 	ctx.JSON(http.StatusOK, TaskResponse{
// 		Success: true,
// 		Message: "Update User sukses",
// 		Result:  UpdateUser,
// 	})

// }

// func DeleteUser(ctx *gin.Context) {
// 	iddb, _ := strconv.Atoi(ctx.Param("id"))
// 	user := models.UserById(iddb)
// 	if user == (models.User{}) {
// 		ctx.JSON(http.StatusBadRequest, TaskResponse{
// 			Success: false,
// 			Message: "invalid add user",
// 			Result:  iddb,
// 		})
// 		return
// 	}

// 	// ctx.ShouldBind(&user)

// 	DeleteUser := models.DeleteUser(iddb)

// 	ctx.JSON(http.StatusOK, TaskResponse{
// 		Success: true,
// 		Message: "Delete User sukses",
// 		Result:  DeleteUser,
// 	})

// }

func UsersId(ctx *gin.Context) {
	iddb, _ := strconv.Atoi(ctx.Param("id"))
	user := models.UserById(iddb)
	ctx.JSON(http.StatusOK, TaskResponse{
		Success: true,
		Message: "User Id",
		Result:  user,
	})

}

func AllUsersDB(ctx *gin.Context) {
	user := models.UserAll()
	// fmt.Println(user)
	ctx.JSON(http.StatusOK, TaskResponse{
		Success: true,
		Message: "All User",
		Result:  user,
	})

}

// router.Run("localhost:8888")
