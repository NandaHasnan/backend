package controllers

import (
	lib "backend/lib"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthLogin(c *gin.Context) {
	var login models.User_credentials

	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: "Invalid request payload",
			Result:  nil,
		})
		return
	}

	user := models.UserByEmail(login.Email)
	if user == (models.User_credentials{}) {
		c.JSON(http.StatusNotFound, TaskResponse2{
			Success: false,
			Message: "Email not found",
			Result:  nil,
		})
		return
	}

	if !lib.VerifyHash(user.Password, login.Password) {
		c.JSON(http.StatusUnauthorized, TaskResponse2{
			Success: false,
			Message: "Invalid password",
			Result:  nil,
		})
		return
	}

	token := lib.GeneratePass(struct {
		UserId int `json:"userid"`
	}{
		UserId: user.Id,
	})

	c.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "Login successful",
		Result:  token,
	})
}

func Register(ctx *gin.Context) {

	var newUser models.User_credentials
	ctx.ShouldBind(&newUser)
	// fmt.Printf("Received User Data: %+v\n", newUser)

	err := validatePassword(newUser.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	if newUser.Password != "" {
		newUser.Password = lib.GenerateHash(newUser.Password)
	}

	models.InserUser(newUser)

	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "Register sukses",
		// Result:  UserN,
	})

}
