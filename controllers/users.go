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

// Users godoc
// @Schemes
// @Description Add New Users
// @Tags Users
// @Accept mpfd
// @Produce json
// @param email formData string false "Email"
// @param password formData string false "password"
// @param firstname formData string false "Firstname"
// @param lastname formData string false "Lastname"
// @param phone_number formData string false "Phone Number"
// @param image formData file false "Image"
// @Success 200 {object} TaskResponse2{result=models.Gabung}
// @Security ApiKeyAuth
// @Router /users/add_User [post]
func InsertUser(ctx *gin.Context) {

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

	if password != "" {
		password = lib.GenerateHash(password)
	}

	newUser := models.Gabung{
		Firstname:    firstname,
		Lastname:     lastname,
		Phone_number: phone_number,
		Image:        fileName + "." + strings.Split(file.Filename, ".")[1],
		Email:        email,
		Password:     password,
	}

	result, err := models.AddUser(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, TaskResponse2{
			Success: false,
			Message: "Error adding user",
		})
		return
	}

	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "User added successfully",
		Result:  result,
	})
}

// Users godoc
// @Schemes
// @Description Edit Profile Users
// @Tags Profile
// @Accept mpfd
// @Produce json
// @param firstname formData string false "Firstname"
// @param lastname formData string false "Lastname"
// @param phone_number formData string false "Phone Number"
// @param image formData file false "Image"
// @param email formData string false "Email"
// @param password formData string false "password"
// @Success 200 {object} TaskResponse2{result=models.Gabung}
// @Security ApiKeyAuth
// @Router /users [patch]
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

	// err = ctx.ShouldBind(&updatedUser)
	// if err != nil {
	// 	if strings.Contains(err.Error(), "Key: 'Gabung.Email' Error:Field validation for 'Email' failed on the 'email' tag") {
	// 		ctx.JSON(http.StatusBadRequest, TaskResponse2{
	// 			Success: false,
	// 			Message: "invalid email format",
	// 			Result:  nil,
	// 		})
	// 		return
	// 	}

	// 	if strings.Contains(err.Error(), "Key: 'Gabung.Password' Error:Field validation for 'Password' failed") {
	// 		ctx.JSON(http.StatusBadRequest, TaskResponse2{
	// 			Success: false,
	// 			Message: "invalid password",
	// 			Result:  nil,
	// 		})
	// 		return
	// 	}
	// }

	result, _ := models.UpdateUser(updatedUser)

	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "User updated successfully",
		Result:  result,
	})
}

// Users godoc
// @Schemes
// @Description Edit Profile Users
// @Tags Users
// @Accept mpfd
// @Produce json
// @Param id path int true "Edit User"
// @param firstname formData string false "Firstname"
// @param lastname formData string false "Lastname"
// @param phone_number formData string false "Phone Number"
// @param image formData file false "Image"
// @param email formData string false "Email"
// @param password formData string false "password"
// @Success 200 {object} TaskResponse2{result=models.Gabung}
// @Security ApiKeyAuth
// @Router /users/{id} [patch]
func EditUserAdmin(ctx *gin.Context) {
	// val, isAvail := ctx.Get("userid")
	// if !isAvail {
	// 	ctx.JSON(http.StatusUnauthorized, TaskResponse2{
	// 		Success: false,
	// 		Message: "Unauthorized",
	// 	})
	// 	return
	// }

	// userId := int(val.(float64))

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

	ctx.ShouldBind(&user)

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
			Message: "File size is too large, max 2 mb",
		})
		return
	}

	var fileName string
	if file.Filename != "" {
		splitFile := strings.Split(file.Filename, ".")
		if len(splitFile) < 2 {
			ctx.JSON(http.StatusBadRequest, TaskResponse2{
				Success: false,
				Message: "Invalid file format",
			})
			return
		}

		fileExt := strings.ToLower(splitFile[len(splitFile)-1])

		allowedExtensions := map[string]bool{
			"jpg": true,
			"png": true,
		}

		if !allowedExtensions[fileExt] {
			ctx.JSON(http.StatusBadRequest, TaskResponse2{
				Success: false,
				Message: "Only .jpg and .png files are allowed",
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
		Id:           &iddb,
		Firstname:    firstname,
		Lastname:     lastname,
		Phone_number: phone_number,
		Image:        fileName + "." + strings.Split(file.Filename, ".")[1],
		Email:        email,
		Password:     password,
	}

	result, err := models.UpdateUser(updatedUser)
	if err != nil {
		if strings.Contains(err.Error(), "Key: 'User_credentials.Email' Error:Field validation for 'Email' failed on the 'email' tag") {
			ctx.JSON(http.StatusBadRequest, TaskResponse2{
				Success: false,
				Message: "invalid email format",
				Result:  nil,
			})
			return
		}

		if strings.Contains(err.Error(), "Key: 'User_credentials.Password' Error:Field validation for 'Password' failed") {
			ctx.JSON(http.StatusBadRequest, TaskResponse2{
				Success: false,
				Message: "invalid password",
				Result:  nil,
			})
			return
		}
	}

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

// Users godoc
// @Schemes
// @Description Delete User
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} TaskResponse2{result=models.Gabung}
// @Security ApiKeyAuth
// @Router /users [delete]
func DeleteUser(ctx *gin.Context) {
	// iddb, _ := strconv.Atoi(ctx.Param("id"))
	// user := models.UserById(iddb)
	// if user == (models.Gabung{}) {
	// 	ctx.JSON(http.StatusBadRequest, TaskResponse2{
	// 		Success: false,
	// 		Message: "invalid add user",
	// 		Result:  iddb,
	// 	})
	// 	return
	// }

	val, isAvail := ctx.Get("userid")
	userId := int(val.(float64))

	// ctx.ShouldBind(&user)

	DeleteUser, err := models.DeleteUser(userId)
	if err != nil {
		log.Println(err)
	}

	if isAvail {
		ctx.JSON(http.StatusOK, TaskResponse2{
			Success: true,
			Message: "Delete User sukses",
			Result:  DeleteUser,
		})
	}

}

// Users godoc
// @Schemes
// @Description Delete User
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "Delete User"
// @Success 200 {object} TaskResponse2{result=models.Gabung}
// @Security ApiKeyAuth
// @Router /users/{id} [delete]
func DeleteUserAdmin(ctx *gin.Context) {
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

	// val, isAvail := ctx.Get("userid")
	// userId := int(val.(float64))

	ctx.ShouldBind(&user)

	DeleteUser, err := models.DeleteUser(iddb)
	if err != nil {
		log.Println(err)
	}

	// if isAvail {
	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "Delete User sukses",
		Result:  DeleteUser,
	})
	// }

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

// func UserProfile(ctx *gin.Context) {
// 	val, isAvail := ctx.Get("userid")
// 	userId := int(val.(float64))

// 	profile := models.UserById(userId)

// 	if isAvail {
// 		ctx.JSON(http.StatusOK, TaskResponse2{
// 			Success: true,
// 			Message: "User Id",
// 			Result:  profile,
// 		})
// 	}
// }

// Users godoc
// @Schemes
// @Description Get All Users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} TaskResponse2{result=models.ListUsersGabung}
// @Security ApiKeyAuth
// @Router /users [get]
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
