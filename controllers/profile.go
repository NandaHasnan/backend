package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	val, isAval := c.Get("userId")

	if isAval {
		c.String(http.StatusOK, "%s", int(val.(float64)))
	}
}
