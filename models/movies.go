package models

import (
	"backend/lib"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type Allmovie struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" form:"title" db:"title"`
	Image_movie string `json:"image_movie" form:"image_movie" db:"image_movie"`
	Genre       string `json:"genre" form:"genre" db:"genre"`
}

type Movie struct {
	Id           int       `json:"id" db:"id"`
	Title        string    `json:"title" form:"title" db:"title"`
	Image_movie  string    `json:"image_movie" form:"image_movie" db:"image_movie"`
	Genre        string    `json:"genre" form:"genre" db:"genre"`
	Release_date time.Time `json:"release_date" form:"release_date" db:"release_date"`
	Duration     string    `json:"duration" form:"duration" db:"duration"`
	Director     string    `json:"director" form:"director" db:"director"`
	Cast_actor   string    `json:"cast_actor" form:"cast_actor" db:"cast_actor"`
	Synopsis     string    `json:"synopsis" form:"synopsis" db:"synopsis"`
	// CreateAt     time.Time  `json:"createdAt" db:"created_at"`
	// UpdateAt     *time.Time `json:"updatedAt" db:"updated_at"`
	// Description string     `json:"description" form:"description"`
}

type Movie_cinema struct {
	Id           int       `json:"id" db:"id"`
	Title        string    `json:"title" form:"title" db:"title"`
	Image_movie  string    `json:"image_movie" form:"image_movie" db:"image_movie"`
	Genre        string    `json:"genre" form:"genre" db:"genre"`
	Release_date time.Time `json:"release_date" form:"release_date" db:"release_date"`
	Duration     string    `json:"duration" form:"duration" db:"duration"`
	Director     string    `json:"director" form:"director" db:"director"`
	Cast_actor   string    `json:"cast_actor" form:"cast_actor" db:"cast_actor"`
	Synopsis     string    `json:"synopsis" form:"synopsis" db:"synopsis"`
	Cinema_name  string    `json:"cinema_name" form:"cinema_name" db:"cinema_name"`
	Price        int       `json:"price" form:"price" db:"price"`
	Location     string    `json:"location" form:"location" db:"location"`
	Date         time.Time `json:"date" form:"date" db:"date"`
	Time         string    `json:"time" form:"time" db:"time"`
}

type Order struct {
	Id              int    `json:"id" db:"id"`
	Movie_Cinema_Id int    `json:"movie_cinema_id" form:"movie_cinema_id" db:"movie_cinema_id"`
	Quantity        int    `json:"quantity" form:"quantity" db:"quantity"`
	TotalPrice      int    `json:"total_price" form:"total_price" db:"total_price"`
	Date            string `json:"date" form:"date" db:"date"`
	Time            string `json:"time" form:"time" db:"time"`
	// MovieId    int    `json:"movie_id"`
	// CinemaId   int    `json:"cinema_id"`
}

type Data []Movie
type Data2 []Allmovie

// type ListUsers []User

func MovieAll(page int, limit int, search string, sort string) Data2 {
	conn := lib.DB()
	defer conn.Close(context.Background())

	offset := (page - 1) * limit

	search = fmt.Sprintf("%%%s%%", search)

	query := fmt.Sprintf(`
		SELECT id, title, image_movie, genre
		FROM movie
		WHERE title ILIKE $3
		ORDER BY id %s
		LIMIT $1 OFFSET $2
	`, sort)

	rows, err := conn.Query(context.Background(), query, limit, offset, search)

	if err != nil {
		fmt.Println(err)
	}

	movies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Allmovie])
	if err != nil {
		fmt.Println(err)
	}
	return movies
}

func CountMovie(search string) int {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var total int

	search = fmt.Sprintf("%%%s%%", search)

	conn.QueryRow(context.Background(), `
	SELECT COUNT(id) FROM movie WHERE title ILIKE $1
	`, search).Scan(&total)

	return total
}

func MovieById(iddb int) Movie_cinema {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var id int
	var title string
	var image_movie string
	var genre string
	var release_date time.Time
	var duration string
	var director string
	var cast_actor string
	var synopsis string
	var cinema_name string
	var location string
	var date time.Time
	var time string

	// var detail Movie_cinema

	err := conn.QueryRow(context.Background(), `
	select movie.id, movie.title, movie.image_movie, movie.genre, 
	movie.release_date, movie.duration, movie.director, 
	movie.cast_actor, movie.synopsis, cinema.cinema_name, 
	cinema."location", cinema."date", cinema."time"
from movie join movie_cinema on movie_cinema.movie_id = movie.id 
join cinema on movie_cinema.cinema_id = cinema.id where movie.id = $1; 
	`, iddb).Scan(
		&id,
		&title,
		&image_movie,
		&genre,
		&release_date,
		&duration,
		&director,
		&cast_actor,
		&synopsis,
		&cinema_name,
		&location,
		&date,
		&time,
	)

	if err != nil {
		// ctx.JSON(http.StatusBadRequest, controllers.TaskResponse{
		// 	Success: false,
		// 	Message: "invalid all user",
		// 	// Result:  id,
		// })

		fmt.Println(err)
	}
	return Movie_cinema{
		Id:           id,
		Title:        title,
		Image_movie:  image_movie,
		Genre:        genre,
		Release_date: release_date,
		Duration:     duration,
		Director:     director,
		Cast_actor:   cast_actor,
		Synopsis:     synopsis,
		Cinema_name:  cinema_name,
		Location:     location,
		Date:         date,
		Time:         time,
	}
}

// func UserByEmail(email string) User {

// 	conn := lib.DB()
// 	defer conn.Close(context.Background())

// 	var user User

// 	conn.QueryRow(context.Background(), `
// 	SELECT id, email, password FROM users WHERE email = $1
// 	`, email).Scan(&user.Id, &user.Email, &user.Password)

// 	return user

// }

func InsertMovie(data Movie) Movie {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movieInsert Movie

	err := conn.QueryRow(context.Background(), `
		INSERT INTO movie (title, image_movie, genre, release_date, duration, director, cast_actor, synopsis) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, title, image_movie, genre, release_date, duration, director, cast_actor, synopsis
	`, data.Title, data.Image_movie, data.Genre, data.Release_date, data.Duration, data.Director, data.Cast_actor, data.Synopsis).
		Scan(
			&movieInsert.Id,
			&movieInsert.Title,
			&movieInsert.Image_movie,
			&movieInsert.Genre,
			&movieInsert.Release_date,
			&movieInsert.Duration,
			&movieInsert.Director,
			&movieInsert.Cast_actor,
			&movieInsert.Synopsis,
		)

	if err != nil {
		fmt.Println(err)
	}

	return movieInsert
}

// func UpdateMovie(movie Movie) Movie {
// 	conn := lib.DB()
// 	defer conn.Close(context.Background())

// 	var movieUpdate Movie

// 	conn.QueryRow(context.Background(), `
// 		UPDATE movie
// 		SET title = $1, image = $2, description = $3, created_at = $4, updated_at = $5
// 		WHERE id = $6
// 		RETURNING id, title, image, description, created_at, updated_at
// 	`, movie.Title, movie.Image, movie.Description, movie.CreateAt, movie.UpdateAt, movie.Id).Scan(
// 		&movieUpdate.Id,
// 		&movieUpdate.Title,
// 		&movieUpdate.Image,
// 		&movieUpdate.Description,
// 		&movieUpdate.CreateAt,
// 		&movieUpdate.UpdateAt,
// 	)

// 	return movieUpdate
// }

// func DeleteMovie(iddb int) Movie {
// 	conn := lib.DB()
// 	defer conn.Close(context.Background())

// 	var movieDelete Movie

// 	conn.QueryRow(context.Background(), `
// 	DELETE FROM movie WHERE id = $1
// 	RETURNING  id, title, image, description, created_at, updated_at
// 	`, iddb).Scan(
// 		&movieDelete.Id,
// 		&movieDelete.Title,
// 		&movieDelete.Image,
// 		&movieDelete.Description,
// 		&movieDelete.CreateAt,
// 		&movieDelete.UpdateAt,
// 	)

// 	return movieDelete
// }

func OrderTicket(data Order) Order {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var order Order
	var price int

	// err := conn.QueryRow(context.Background(), `
	// 	SELECT cinema.price FROM movie_cinema
	// 	JOIN cinema ON movie_cinema.cinema_id = cinema.id
	// 	WHERE movie_cinema.id = $1
	// `, data.Movie_Cinema_Id).Scan(&price)
	// if err != nil {
	// 	fmt.Printf("Error fetching price for movie_cinema_id %d: %v\n", data.Movie_Cinema_Id, err)
	// }

	totalPrice := price * data.Quantity

	err := conn.QueryRow(context.Background(), `
		INSERT INTO orders (movie_cinema_id, quantity, total_price, date, time) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, movie_cinema_id, quantity, total_price, date, time
	`, data.Movie_Cinema_Id, data.Quantity, totalPrice, data.Date, data.Time).Scan(
		&order.Id,
		&order.Movie_Cinema_Id,
		&order.Quantity,
		&order.TotalPrice,
		&order.Date,
		&order.Time,
	)
	if err != nil {
		fmt.Printf("Error inserting order: %v\n", err)
	}

	return order
}
