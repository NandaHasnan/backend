package models

import (
	"backend/lib"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type Allmovie struct {
	Id           int       `json:"id" db:"id" example:"1"`
	Title        string    `json:"title" form:"title" db:"title" example:"Spiderman"`
	Image_movie  string    `json:"image_movie" form:"image_movie" db:"image_movie" example:"b00db012-1a27-43b3-895a-abd3f540362e.jpg"`
	Genre        string    `json:"genre" form:"genre" db:"genre" example:"Action, Advanture"`
	Release_date time.Time `json:"release_date" form:"release_date" db:"release_date" example:"2006-01-02"`
	Duration     string    `json:"duration" form:"duration" db:"duration" example:"22:10:33"`
}

type Movie struct {
	Id           int       `json:"id" db:"id" example:"1"`
	Title        string    `json:"title" form:"title" db:"title" example:"Spiderman"`
	Image_movie  string    `db:"image_movie" example:"b00db012-1a27-43b3-895a-abd3f540362e.jpg"`
	Genre        string    `json:"genre" form:"genre" db:"genre" example:"Action, Advanture"`
	Release_date time.Time `db:"release_date" example:"2006-01-02"`
	Duration     string    `db:"duration" example:"22:10:33"`
	Director     string    `json:"director" form:"director" db:"director" example:"joko anwar"`
	Cast_actor   string    `json:"cast_actor" form:"cast_actor" db:"cast_actor" example:"warsito"`
	Synopsis     string    `json:"synopsis" form:"synopsis" db:"synopsis" example:"Lorem ipsum dolor sit amet"`
	// CreateAt     time.Time  `json:"createdAt" db:"created_at"`
	// UpdateAt     *time.Time `json:"updatedAt" db:"updated_at"`
	// Description string     `json:"description" form:"description"`
}

type Movie_body struct {
	Movie
	// Id           int    `json:"id" `
	// Title        string `json:"title" form:"title"`
	// Image_movie  string `json:"image_movie" form:"image_movie" `
	// Genre        string `json:"genre" form:"genre" `
	Release_date string `json:"release_date" form:"release_date" example:"2006-01-02"`
	Duration     string `json:"duration" form:"duration" example:"22:10:33"`
	// Director     string `json:"director" form:"director" `
	// Cast_actor   string `json:"cast_actor" form:"cast_actor" `
	// Synopsis     string `json:"synopsis" form:"synopsis" `
}

type Movie_Data struct {
	Movie
	Release_date time.Time `db:"release_date"`
	Duration     time.Time `db:"duration"`
}

type Movie_cinema struct {
	Id           int       `json:"id" db:"id" example:"1"`
	Title        string    `json:"title" form:"title" db:"title" example:"Spiderman"`
	Image_movie  string    `json:"image_movie" form:"image_movie" db:"image_movie" example:"b00db012-1a27-43b3-895a-abd3f540362e.jpg"`
	Genre        string    `json:"genre" form:"genre" db:"genre" example:"Action, Advanture"`
	Release_date time.Time `json:"release_date" form:"release_date" db:"release_date" example:"2025-10-10"`
	Duration     string    `json:"duration" form:"duration" db:"duration" example:"02:09:09"`
	Director     string    `json:"director" form:"director" db:"director" example:"Steven Spielberg"`
	Cast_actor   string    `json:"cast_actor" form:"cast_actor" db:"cast_actor" example:"Tom Holland, Andrew Garfield, Tobey Maguire, Zendaya"`
	Synopsis     string    `json:"synopsis" form:"synopsis" db:"synopsis" example:"Lorem ipsum dolor sit amet"`
	// Cinema_name  string    `json:"cinema_name" form:"cinema_name" db:"cinema_name" example:"cinemas, XXI"`
	// Location     string    `json:"location" form:"location" db:"location" example:"semarang, solo"`
	// Date         time.Time `json:"date" form:"date" db:"date" example:"2025-12-01"`
	// Time         string    `json:"time" form:"time" db:"time" example:"01:35:00"`
}

type order struct {
	Movie_Cinema_Id int      `json:"movie_cinema_id" form:"movie_cinema_id" db:"movie_cinema_id" example:"1"`
	Title           string   `json:"title" form:"title" db:"title" example:"Spiderman"`
	Image_movie     string   `json:"image_movie" form:"image_movie" db:"image_movie" example:"b00db012-1a27-43b3-895a-abd3f540362e.jpg"`
	Genre           string   `json:"genre" form:"genre" db:"genre" example:"Action, Advanture"`
	Quantity        int      `json:"quantity" form:"quantity" db:"quantity" example:"5"`
	TotalPrice      int      `json:"total_price" form:"total_price" db:"total_price" example:"90000"`
	Cinema_name     string   `json:"cinema_name" form:"cinema_name" db:"cinema_name" example:"Hiflix"`
	Location        string   `json:"location" form:"location" db:"location" example:"Bandung"`
	Seat            []string `json:"seat" form:"seat[]" db:"seat"`
	User_Id         int      `json:"user_id" form:"user_id" db:"user_id" example:"1"`
	Movie_Id        int      `json:"movie_id" form:"movie_id" db:"movie_id" example:"1"`
	Cinema_Id       int      `json:"cinema_id" form:"cinema_id" db:"cinema_id" example:"1"`
	Firstname       string   `json:"firstname" form:"firstname" db:"firstname" example:"doni"`
	Email           string   `json:"email" form:"email" db:"email" example:"doni@mail.com"`
	Phone_number    string   `json:"phone_number" form:"phone_number" db:"phone_number" example:"+6232574365"`
}

type OrderBody struct {
	// Id              int       `json:"id" db:"id"`
	order
	Date string `json:"date" form:"date" example:"2006-01-02"`
	Time string `json:"time" form:"time" example:"22:10:33"`
	// MovieId    int    `json:"movie_id"`
	// CinemaId   int    `json:"cinema_id"`
}

type OrderData struct {
	Id int `db:"id"`
	order
	Date time.Time `db:"date"`
	Time time.Time `db:"time"`
	// MovieId    int    `json:"movie_id"`
	// CinemaId   int    `json:"cinema_id"`
}

type OrderSagger struct {
	Movie_Cinema_Id int      `json:"movie_cinema_id" form:"movie_cinema_id" db:"movie_cinema_id" example:"1"`
	Quantity        int      `json:"quantity" form:"quantity" db:"quantity" example:"5"`
	Date            string   `json:"date" form:"date" example:"2006-01-02"`
	Time            string   `json:"time" form:"time" example:"22:10:33"`
	Seat            []string `json:"seat" form:"seat[]" db:"seat"`
}

type Data []Movie
type Data2 []Allmovie
type Data3 []Movie_cinema
type Data4 []Cinema

// type ListUsers []User

func MovieAll(page int, limit int, search string, sort string) Data2 {
	conn := lib.DB()
	defer conn.Close(context.Background())

	offset := (page - 1) * limit

	search = fmt.Sprintf("%%%s%%", search)

	query := fmt.Sprintf(`
		SELECT id, title, image_movie, genre, release_date, duration
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
	// var cinema_name string
	// var location string
	// var date time.Time
	// var time string

	// var detail Movie_cinema

	err := conn.QueryRow(context.Background(), `
	select movie.id, movie.title, movie.image_movie, movie.genre, 
	movie.release_date, movie.duration, movie.director, 
	movie.cast_actor, movie.synopsis
from movie where movie.id = $1; 
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
		// &cinema_name,
		// &location,
		// &date,
		// &time,
	)

	// 	select movie.id, movie.title, movie.image_movie, movie.genre,
	// 	movie.release_date, movie.duration, movie.director,
	// 	movie.cast_actor, movie.synopsis, cinema.cinema_name,
	// 	cinema."location", cinema."date", cinema."time"
	// from movie join movie_cinema on movie_cinema.movie_id = movie.id
	// join cinema on movie_cinema.cinema_id = cinema.id where movie.id = $1;

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
		// Cinema_name:  cinema_name,
		// Location:     location,
		// Date:         date,
		// Time:         time,
	}
}

func MovieById2(iddb int) (Movie, error) {
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

	// var detail Movie_cinema

	err := conn.QueryRow(context.Background(), `
	select id, title, image_movie, genre,
	release_date, duration, director,
	cast_actor, synopsis
from movie where id = $1;
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
	)

	return Movie{
		Id:           id,
		Title:        title,
		Image_movie:  image_movie,
		Genre:        genre,
		Release_date: release_date,
		Duration:     duration,
		Director:     director,
		Cast_actor:   cast_actor,
		Synopsis:     synopsis,
	}, err
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

// func InsertMovie(data Movie_body) (Movie_Data, error) {
// 	conn := lib.DB()
// 	defer conn.Close(context.Background())

// 	var movieInsert Movie_Data

// 	// if data.Release_date == "" {
// 	// 	return Movie_Data{}, fmt.Errorf("release date is required")
// 	// }
// 	movieDate, err := time.Parse("2006-01-02", data.Release_date)
// 	if err != nil {
// 		return Movie_Data{}, fmt.Errorf("invalid release date format, expected YYYY-MM-DD")
// 	}

// 	log.Println(data.Release_date)
// 	// if data.Duration == "" {
// 	// 	return Movie_Data{}, fmt.Errorf("duration is required")
// 	// }
// 	movieDuration, err := time.Parse("15:04:05", data.Duration)
// 	if err != nil {
// 		return Movie_Data{}, fmt.Errorf("invalid duration format, expected HH:mm:ss: %v", err)
// 	}

// 	err = conn.QueryRow(context.Background(), `
// 		INSERT INTO movie (title, image_movie, genre, release_date, duration, director, cast_actor, synopsis)
// 		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
// 		RETURNING id, title, image_movie, genre, release_date, duration, director, cast_actor, synopsis
// 	`, data.Title, data.Image_movie, data.Genre, movieDate, movieDuration, data.Director, data.Cast_actor, data.Synopsis).
// 		Scan(
// 			&movieInsert.Id,
// 			&movieInsert.Title,
// 			&movieInsert.Image_movie,
// 			&movieInsert.Genre,
// 			&movieInsert.Release_date,
// 			&movieInsert.Duration,
// 			&movieInsert.Director,
// 			&movieInsert.Cast_actor,
// 			&movieInsert.Synopsis,
// 		)

// 	if err != nil {
// 		return Movie_Data{}, fmt.Errorf("failed to insert movie: %v", err)
// 	}

// 	return movieInsert, nil
// }

func InsertMovie(data Movie_body) (Movie_Data, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movieInsert Movie_Data

	// if data.Release_date == "" {
	// 	return Movie_Data{}, fmt.Errorf("release date is required")
	// }
	movieDate, err := time.Parse("2006-01-02", data.Release_date)
	if err != nil {
		return Movie_Data{}, fmt.Errorf("invalid release date format, expected YYYY-MM-DD")
	}

	log.Println(data.Release_date)
	// if data.Duration == "" {
	// 	return Movie_Data{}, fmt.Errorf("duration is required")
	// }
	movieDuration, err := time.Parse("15:04:05", data.Duration)
	if err != nil {
		return Movie_Data{}, fmt.Errorf("invalid duration format, expected HH:mm:ss: %v", err)
	}

	err = conn.QueryRow(context.Background(), `
		INSERT INTO movie ( image_movie, title, genre, release_date, duration, director, cast_actor, synopsis) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, title, image_movie, genre, release_date, duration, director, cast_actor, synopsis
	`, data.Image_movie, data.Title, data.Genre, movieDate, movieDuration, data.Director, data.Cast_actor, data.Synopsis).
		Scan(
			&movieInsert.Id,
			&movieInsert.Image_movie,
			&movieInsert.Title,
			&movieInsert.Genre,
			&movieInsert.Release_date,
			&movieInsert.Duration,
			&movieInsert.Director,
			&movieInsert.Cast_actor,
			&movieInsert.Synopsis,
		)

	if err != nil {
		return Movie_Data{}, fmt.Errorf("failed to insert movie: %v", err)
	}

	return movieInsert, nil
}

func UpdateMovie(movie Movie_body) (Movie_Data, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movieUpdate Movie_Data

	// Check if Release_date is empty
	if movie.Release_date == "" {
		return Movie_Data{}, fmt.Errorf("release date cannot be empty")
	}

	movieDate, err := time.Parse("2006-01-02", movie.Release_date)
	if err != nil {
		return Movie_Data{}, fmt.Errorf("invalid release date format, expected YYYY-MM-DD")
	}

	log.Println(movie)

	movieDuration, err := time.Parse("15:04:05", movie.Duration)
	if err != nil {
		return Movie_Data{}, fmt.Errorf("invalid duration format, expected HH:mm:ss: %v", err)
	}

	if movie.Image_movie != "" || movie.Title != "" || movie.Genre != "" || movie.Release_date != "" || movie.Duration != "" || movie.Director != "" || movie.Cast_actor != "" || movie.Synopsis != "" {
		conn.QueryRow(context.Background(), `
		UPDATE movie
		SET image_movie = $1, title = $2, genre = $3, release_date = $4, duration = $5, director = $6, cast_actor = $7, synopsis = $8
		WHERE id = $9
		RETURNING id, image_movie, title, genre,  release_date, duration, director, cast_actor, synopsis
	`, movie.Image_movie, movie.Title, movie.Genre, movieDate, movieDuration, movie.Director, movie.Cast_actor, movie.Synopsis, movie.Id).Scan(
			&movieUpdate.Id,
			&movieUpdate.Image_movie,
			&movieUpdate.Title,
			&movieUpdate.Genre,
			&movieUpdate.Release_date,
			&movieUpdate.Duration,
			&movieUpdate.Director,
			&movieUpdate.Cast_actor,
			&movieUpdate.Synopsis,
		)
	}
	return movieUpdate, nil
}

func DeleteMovie(iddb int) Movie {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movieDelete Movie

	conn.QueryRow(context.Background(), `
	DELETE FROM movie WHERE id = $1
	RETURNING  id, title, image_movie, genre, release_date, duration, director, cast_actor, synopsis
	`, iddb).Scan(
		&movieDelete.Id,
		&movieDelete.Title,
		&movieDelete.Image_movie,
		&movieDelete.Genre,
		&movieDelete.Release_date,
		&movieDelete.Duration,
		&movieDelete.Director,
		&movieDelete.Cast_actor,
		&movieDelete.Synopsis,
	)

	return movieDelete
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

// 	totalPrice := price * data.Quantity
// 	fmt.Println(data.Time)
// 	fmt.Println(data.Date)
// 	orderTime, err := time.Parse(time.TimeOnly, data.Time)

// 	if err != nil {
// 		fmt.Println(err)
// 		return OrderData{}, err
// 	}

// 	orderDate, err := time.Parse(time.DateOnly, data.Date)
// 	if err != nil {
// 		fmt.Println(err)
// 		return OrderData{}, err
// 	}

// 	err = conn.QueryRow(context.Background(), `
// 		INSERT INTO orders (movie_cinema_id, quantity, total_price, date, time)
// 		VALUES ($1, $2, $3, $4, $5)
// 		RETURNING id, movie_cinema_id, quantity, total_price, date, time
// 	`, data.Movie_Cinema_Id, data.Quantity, totalPrice, orderDate, orderTime).Scan(
// 		&order.Id,
// 		&order.Movie_Cinema_Id,
// 		&order.Quantity,
// 		&order.TotalPrice,
// 		&order.Date,
// 		&order.Time,
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

func Update(data Movie_body) (Movie_Data, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var movieUpdate Movie_Data
	var movieDate time.Time
	log.Println(data.Release_date)

	if data.Release_date != "" {
		parsedDate, err := time.Parse("2006-01-02", data.Release_date)
		if err != nil {
			return Movie_Data{}, fmt.Errorf("invalid release date format, expected YYYY-MM-DD")
		}
		movieDate = parsedDate
	}

	var movieDuration string
	if data.Duration != "" {
		movieDuration = data.Duration
	}

	log.Println(data.Duration)

	err := conn.QueryRow(context.Background(), `
		UPDATE movie
		SET title = COALESCE($1, title),
		    image_movie = COALESCE($2, image_movie),
		    genre = COALESCE($3, genre),
		    release_date = COALESCE($4, release_date),
		    duration = COALESCE($5, duration),
		    director = COALESCE($6, director),
		    cast_actor = COALESCE($7, cast_actor),
		    synopsis = COALESCE($8, synopsis)
		WHERE id = $9
		RETURNING id, title, image_movie, genre, release_date, duration, director, cast_actor, synopsis
	`, data.Title, data.Image_movie, data.Genre, movieDate, movieDuration, data.Director, data.Cast_actor, data.Synopsis, data.Id).Scan(
		&movieUpdate.Id,
		&movieUpdate.Title,
		&movieUpdate.Image_movie,
		&movieUpdate.Genre,
		&movieUpdate.Release_date,
		&movieUpdate.Duration,
		&movieUpdate.Director,
		&movieUpdate.Cast_actor,
		&movieUpdate.Synopsis,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return Movie_Data{}, fmt.Errorf("movie with ID %d not found for update", data.Id)
		}
		return Movie_Data{}, fmt.Errorf("failed to update movie: %v", err)
	}

	return movieUpdate, nil
}

// 15:04:05 +0000 UTC
