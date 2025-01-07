package controllers

import (
	lib "backend/lib"
	"backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Schemes
// @Description Login
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @param email formData string true "Email"
// @param password formData string true "Password"
// @Success 200 {object} TaskResponse2{result=models.User_credentials}
// @Router /auth/login [post]
func AuthLogin(c *gin.Context) {
	var login models.User_credentials
	err := c.ShouldBind(&login)
	if err != nil {
		if strings.Contains(err.Error(), "Key: 'User_credentials.Email' Error:Field validation for 'Email' failed on the 'email' tag") {
			c.JSON(http.StatusBadRequest, TaskResponse2{
				Success: false,
				Message: "invalid email format",
				Result:  nil,
			})
			return
		}

		if strings.Contains(err.Error(), "Key: 'User_credentials.Password' Error:Field validation for 'Password' failed") {
			c.JSON(http.StatusBadRequest, TaskResponse2{
				Success: false,
				Message: "invalid password",
				Result:  nil,
			})
			return
		}
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

// Register godoc
// @Schemes
// @Description Register
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @param email formData string true "Email"
// @param password formData string true "Password"
// @Success 200 {object} TaskResponse2{}
// @Router /auth/register [post]
func Register(ctx *gin.Context) {

	var newUser models.User_credentials
	err := ctx.ShouldBind(&newUser)
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
				Message: "password min 8",
				Result:  nil,
			})
			return
		}

	}

	if newUser.Password != "" {
		newUser.Password = lib.GenerateHash(newUser.Password)
	}

	_, err = models.InserUser(newUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: err.Error(),
			Result:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "Register sukses",
		// Result:  UserN,
	})

}
