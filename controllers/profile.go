package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Profile godoc
// @Schemes
// @Description Profile Users
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} TaskResponse2{result=models.Gabung}
// @Security ApiKeyAuth
// @Router /profile [get]
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
