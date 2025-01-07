package models

import (
	"backend/lib"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

type Cinema struct {
	Id          int    `json:"id" db:"id" example:"1"`
	Cinema_name string `json:"cinema_name" form:"cinema_name" db:"cinema_name" example:"Spiderman"`
	Location    string `json:"location" form:"location" db:"location" example:"Bandung"`
	Date        string `json:"date" form:"date" db:"date" example:"2006-01-02"`
	Time        string `json:"time" form:"time" db:"time" example:"22:10:33"`
}

type CinemaBody struct {
	Cinema
	Date string `json:"date" form:"date" db:"date" example:"2006-01-02"`
	Time string `json:"time" form:"time" db:"time" example:"22:10:33"`
}

func FilterCinema(date string, time string, location string) (Data4, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	// offset := (page - 1) * limit

	date = fmt.Sprintf("%%%s%%", date)
	time = fmt.Sprintf("%%%s%%", time)
	location = fmt.Sprintf("%%%s%%", location)

	query := `
		SELECT id, cinema_name, location, date, time
		FROM cinema_coba
		WHERE (date ILIKE $1)
		AND time ILIKE $2
		AND location ILIKE $3
	`

	rows, err := conn.Query(context.Background(), query, date, time, location)

	if err != nil {
		fmt.Println(err)
	}

	cinema, err := pgx.CollectRows(rows, pgx.RowToStructByName[Cinema])
	if err != nil {
		fmt.Println(err)
	}
	return cinema, nil
}

func OrderTicket(data OrderBody) (OrderData, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var order OrderData
	var price int

	err := conn.QueryRow(context.Background(), `
		SELECT cinema.price 
		FROM movie_cinema
		JOIN cinema ON movie_cinema.cinema_id = cinema.id
		WHERE movie_cinema.id = $1
	`, data.Movie_Cinema_Id).Scan(&price)
	if err != nil {
		return order, fmt.Errorf("error fetching price for movie_cinema_id %d: %v", data.Movie_Cinema_Id, err)
	}

	totalPrice := price * len(data.Seat)
	orderTime, err := time.Parse(time.TimeOnly, data.Time)
	if err != nil {
		return OrderData{}, fmt.Errorf("invalid time format: %v", err)
	}

	orderDate, err := time.Parse(time.DateOnly, data.Date)
	if err != nil {
		return OrderData{}, fmt.Errorf("invalid date format: %v", err)
	}

	seatArray := fmt.Sprintf(`{%s}`, strings.Join(data.Seat, ",")) // Format as {A1,A2,A3}
	fmt.Println("Seat Array for Database:", seatArray)

	err = conn.QueryRow(context.Background(), `
		INSERT INTO orders (movie_cinema_id, quantity, total_price, date, time, seat, user_id, movie_id, cinema_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, movie_cinema_id, quantity, total_price, date, time, seat, user_id, movie_id, cinema_id
	`, data.Movie_Cinema_Id, data.Quantity, totalPrice, orderDate, orderTime, seatArray, data.User_Id, data.Movie_Id, data.Cinema_Id).Scan(
		&order.Id,
		&order.Movie_Cinema_Id,
		&order.Quantity,
		&order.TotalPrice,
		&order.Date,
		&order.Time,
		&order.Seat,
		&order.User_Id,
		&order.Movie_Id,
		&order.Cinema_Id,
	)
	if err != nil {
		return order, fmt.Errorf("error inserting order: %v", err)
	}

	err = conn.QueryRow(context.Background(), `
		SELECT 
			title, 
			image_movie, 
			genre
		FROM movie
		WHERE movie.id = $1;
	`, data.Movie_Id).Scan(
		&order.Title,
		&order.Image_movie,
		&order.Genre,
	)
	if err != nil {
		return order, fmt.Errorf("error fetching movie details ")
	}

	err = conn.QueryRow(context.Background(), `
		SELECT 
			users.firstname , 
			user_credentials.email, 
			users.phone_number
		FROM users join user_credentials on users.user_credentials_id = user_credentials.id 
		WHERE users.id = $1;
	`, data.User_Id).Scan(
		&order.Firstname,
		&order.Email,
		&order.Phone_number,
	)
	if err != nil {
		return order, fmt.Errorf("error fetching users details ")
	}

	return order, nil
}

// func OrderTicket(data OrderBody) (OrderData, error) {
// 	conn := lib.DB()
// 	defer conn.Close(context.Background())

// 	var order OrderData
// 	var price int

// 	err := conn.QueryRow(context.Background(), `
// 		SELECT cinema.price
// 		FROM movie_cinema
// 		JOIN cinema ON movie_cinema.cinema_id = cinema.id
// 		WHERE movie_cinema.id = $1
// 	`, data.Movie_Cinema_Id).Scan(&price)
// 	if err != nil {
// 		return order, fmt.Errorf("error fetching price for movie_cinema_id %d: %v", data.Movie_Cinema_Id, err)
// 	}

// 	totalPrice := price * len(data.Seat)
// 	orderTime, err := time.Parse(time.TimeOnly, data.Time)
// 	if err != nil {
// 		return OrderData{}, fmt.Errorf("invalid time format: %v", err)
// 	}

// 	orderDate, err := time.Parse(time.DateOnly, data.Date)
// 	if err != nil {
// 		return OrderData{}, fmt.Errorf("invalid date format: %v", err)
// 	}

// 	seatArray := fmt.Sprintf(`{%s}`, strings.Join(data.Seat, ",")) // Format as {A1,A2,A3}
// 	fmt.Println("Seat Array for Database:", seatArray)

// 	err = conn.QueryRow(context.Background(), `
// 		INSERT INTO orders (movie_cinema_id, quantity, total_price, date, time, seat, user_id, movie_id, cinema_id)
// 		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
// 		RETURNING id, movie_cinema_id, quantity, total_price, date, time, seat, user_id, movie_id, cinema_id
// 	`, data.Movie_Cinema_Id, data.Quantity, totalPrice, orderDate, orderTime, seatArray, data.User_Id, data.Movie_Id, data.Cinema_Id).Scan(
// 		&order.Id,
// 		&order.Movie_Cinema_Id,
// 		&order.Quantity,
// 		&order.TotalPrice,
// 		&order.Date,
// 		&order.Time,
// 		&order.Seat,
// 		&order.User_Id,
// 		&order.Movie_Id,
// 		&order.Cinema_Id,
// 	)
// 	if err != nil {
// 		return order, fmt.Errorf("error inserting order: %v", err)
// 	}

// 	err = conn.QueryRow(context.Background(), `
// 		SELECT
// 			movie.title,
// 			movie.image_movie,
// 			movie.genre,
// 			cinema.cinema_name,
// 			cinema.location
// 		FROM movie_cinema
// 		JOIN movie ON movie_cinema.movie_id = movie.id
// 		JOIN cinema ON movie_cinema.cinema_id = cinema.id
// 		WHERE movie_cinema.id = $1
// 	`, data.Movie_Cinema_Id).Scan(
// 		&order.Title,
// 		&order.Image_movie,
// 		&order.Genre,
// 		&order.Cinema_name,
// 		&order.Location,
// 	)
// 	if err != nil {
// 		return order, fmt.Errorf("error fetching movie and cinema details for movie_cinema_id %d: %v", data.Movie_Cinema_Id, err)
// 	}

// 	return order, nil
// }
