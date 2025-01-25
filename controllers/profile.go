package controllers

import (
	"backend/models"
	"net/http"
	"strconv"

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
	// if ctx.Request.Header.Get("Authorization") == "" {

	// }
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

// Profile godoc
// @Schemes
// @Description Profile Users
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "Detail User"
// @Success 200 {object} TaskResponse2{result=models.Gabung}
// @Security ApiKeyAuth
// @Router /users/detail/{id} [get]
func UserProfile2(ctx *gin.Context) {
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

	profile := models.UserById(iddb)

	// if isAvail {
	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "User Id",
		Result:  profile,
	})
	// }
}
