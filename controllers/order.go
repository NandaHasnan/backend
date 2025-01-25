package controllers

import (
	"backend/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Order Movie godoc
// @Schemes
// @Description Order Movie
// @Tags Order
// @Accept x-www-form-urlencoded
// @Produce json
// @param movie_cinema_id formData string true "Movie Name"
// @param quantity formData int true "Quantity"
// @param date formData string true "Date"
// @param time formData string true "Time"
// @param seat[] formData array true "Seat"
// @Success 200 {object} TaskResponse2{result=models.OrderBody}
// @Security ApiKeyAuth
// @Router /order [post]
func OrderMovies(ctx *gin.Context) {

	var orderMovie models.OrderBody
	err := ctx.ShouldBind(&orderMovie)

	if err != nil {
		fmt.Println(err)
		return
	}

	order, err := models.OrderTicket(orderMovie)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "Order tiket sukses",
		Result:  order,
	})

}

// Order Movie godoc
// @Schemes
// @Description Order Movie
// @Tags Order
// @Accept json
// @Produce json
// @param date query string true "Choose Date "
// @param time query string true "Choose Time"
// @param location query string true "Choose Location"
// @Success 200 {object} TaskResponse2{result=models.Cinema}
// @Security ApiKeyAuth
// @Router /order/cinema [get]
func CinemaFilter(ctx *gin.Context) {
	// pageParams, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	// limitParams, err := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	// if err != nil {
	// 	log.Println(err)
	// }
	date := ctx.Query("date")
	time := ctx.Query("time")
	location := ctx.Query("location")
	// sort := ctx.DefaultQuery("sort", "ASC")

	// if sort != "ASC" {
	// 	sort = "DESC"
	// }

	// var movies models.Data4

	movies, err := models.FilterCinema(date, time, location)
	if err != nil {
		log.Printf("Error filtering cinema: %v\n", err)
		ctx.JSON(http.StatusBadRequest, TaskResponse2{
			Success: false,
			Message: "infalid cinema",
			Result:  nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, TaskResponse2{
		Success: true,
		Message: "All Movie",
		Result:  movies,
	})

}

func Payment(ctx *gin.Context) {
	var addpayment models.PaymentInfo
	err := ctx.ShouldBind(&addpayment)
	if err != nil {
		fmt.Println(err)
		return
	}

	payment, err := models.AddPayment(addpayment)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "Payment tiket sukses",
		Result:  payment,
	})
}

func OrderMoviesNew(ctx *gin.Context) {

	var orderMovie models.OrderNew
	err := ctx.ShouldBind(&orderMovie)

	if err != nil {
		fmt.Println(err)
		return
	}

	order, err := models.OrderTicketNew(orderMovie)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, task{
		Success: true,
		Message: "Order tiket sukses",
		Result:  order,
	})

}
