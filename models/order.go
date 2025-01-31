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

type PaymentInfo struct {
	ID                   int        `json:"id" form:"id" db:"id"`
	OrderID              int        `json:"order_id" form:"order_id" db:"order_id"` // Foreign key to orders
	VirtualAccountNumber string     `json:"virtual_account_number" form:"virtual_account_number" db:"virtual_account_number"`
	TotalPayment         float64    `json:"total_payment" form:"total_payment" db:"total_payment"`
	DueDate              time.Time  `json:"due_date" form:"due_date" db:"due_date"`
	PaymentStatus        string     `json:"payment_status" form:"payment_status" db:"payment_status"`
	PaymentTime          *time.Time `json:"payment_time,omitempty" form:"payment_time" db:"payment_time"` // Nullable
	PaymentMethod        string     `json:"payment_method" form:"payment_method" db:"payment_method"`     // Payment method (e.g., Gopay, Visa)
}

type OrderNew struct {
	Id           *int   `json:"id" form:"id" db:"id" example:"1"`
	User_Id      int    `json:"user_id" form:"user_id" db:"user_id" example:"1"`
	Date         string `json:"date" form:"date" db:"date" example:"2006-01-02"`
	Time         string `json:"time" form:"time" db:"time" example:"22:10:33"`
	Movie_Title  string `json:"movie_title" form:"movie_title" db:"movie_title" example:"Spiderman"`
	Cinema_name  string `json:"cinema_name" form:"cinema_name" db:"cinema_name" example:"Spiderman"`
	Total_Seat   string `json:"total_seat" form:"total_seat" db:"total_seat"`
	Total_Price  int    `json:"total_price" form:"total_price" db:"total_price" example:"90000"`
	Full_name    string `json:"full_name" form:"full_name" db:"full_name" example:"doni"`
	Email        string `json:"email" form:"email" db:"email" example:"doni@mail.com"`
	Phone_number string `json:"phone_number" form:"phone_number" db:"phone_number" example:"+6232574365"`
	Payment      string `json:"payment" form:"payment" db:"payment" example:"+6232574365"`
}

type Seat struct {
	Total_Seat string `json:"total_seat" form:"total_seat" db:"total_seat"`
}

type ListSeat []Seat

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

	err = conn.QueryRow(context.Background(), `
		SELECT cinema_name, location
		FROM cinema
		WHERE cinema.id = $1
	`, data.Cinema_Id).Scan(
		&order.Cinema_name,
		&order.Location,
	)
	if err != nil {
		return order, fmt.Errorf("error fetching cinema name ")
	}

	return order, nil
}

// ==================================================================
// ==================================================================
// ==================================================================
func OrderTicketNew(data OrderNew) (OrderNew, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var order OrderNew
	// var price int

	// err := conn.QueryRow(context.Background(), `
	// 	SELECT cinema.price
	// 	FROM movie_cinema
	// 	JOIN cinema ON movie_cinema.cinema_id = cinema.id
	// 	WHERE movie_cinema.id = $1
	// `, data.Movie_Cinema_Id).Scan(&price)
	// if err != nil {
	// 	return order, fmt.Errorf("error fetching price for movie_cinema_id %d: %v", data.Movie_Cinema_Id, err)
	// }

	// totalPrice := price * len(data.Seat)
	// orderTime, err := time.Parse(time.TimeOnly, data.Time)
	// if err != nil {
	// 	return OrderData{}, fmt.Errorf("invalid time format: %v", err)
	// }

	// orderDate, err := time.Parse(time.DateOnly, data.Date)
	// if err != nil {
	// 	return OrderData{}, fmt.Errorf("invalid date format: %v", err)
	// }

	// seatArray := fmt.Sprintf(`{%s}`, strings.Join(data.Total_Seat, ",")) // Format as {A1,A2,A3}
	// fmt.Println("Seat Array for Database:", seatArray)

	err := conn.QueryRow(context.Background(), `
		INSERT INTO order_new (user_id, date, time, movie_title, cinema_name, total_seat, total_price, full_name, email, phone_number, payment)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, user_id, date, time, movie_title, cinema_name, total_seat, total_price, full_name, email, phone_number, payment
	`, data.User_Id, data.Date, data.Time, data.Movie_Title, data.Cinema_name, data.Total_Seat, data.Total_Price, data.Full_name, data.Email, data.Phone_number, data.Payment).Scan(
		&order.Id,
		&order.User_Id,
		&order.Date,
		&order.Time,
		&order.Movie_Title,
		&order.Cinema_name,
		&order.Total_Seat,
		&order.Total_Price,
		&order.Full_name,
		&order.Email,
		&order.Phone_number,
		&order.Payment,
	)
	if err != nil {
		return order, fmt.Errorf("error inserting order: %v", err)
	}

	// err = conn.QueryRow(context.Background(), `
	// 	SELECT
	// 		movie.title,
	// 		movie.image_movie,
	// 		movie.genre,
	// 		cinema.cinema_name,
	// 		cinema.location
	// 	FROM movie_cinema
	// 	JOIN movie ON movie_cinema.movie_id = movie.id
	// 	JOIN cinema ON movie_cinema.cinema_id = cinema.id
	// 	WHERE movie_cinema.id = $1
	// `, data.Movie_Cinema_Id).Scan(
	// 	&order.Title,
	// 	&order.Image_movie,
	// 	&order.Genre,
	// 	&order.Cinema_name,
	// 	&order.Location,
	// )
	// if err != nil {
	// 	return order, fmt.Errorf("error fetching movie and cinema details for movie_cinema_id %d: %v", data.Movie_Cinema_Id, err)
	// }

	return order, nil
}

// func GetAllSeat(data Seat) (Seat, error) {
// 	conn := lib.DB()
// 	defer conn.Close(context.Background())

// 	var order Seat

// 	_, err := conn.Query(context.Background(), `
// 		SELECT total_seat FROM order_new
// 	`)
// 	if err != nil {
// 		return order, fmt.Errorf("error inserting order: %v", err)
// 	}

// 	return order, nil
// }

func GetAllSeat() ([]Seat, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `SELECT total_seat FROM order_new`)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	seats, err := pgx.CollectRows(rows, pgx.RowToStructByName[Seat])
	if err != nil {
		return nil, fmt.Errorf("failed to collect rows: %w", err)
	}

	return seats, nil
}

func AddPayment(data PaymentInfo) (PaymentInfo, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var payment PaymentInfo

	virtualAccount := "12321328913829724"
	dueDate := time.Now().Add(24 * time.Hour)

	err := conn.QueryRow(context.Background(), `
		INSERT INTO payment_info (order_id, virtual_account_number, total_payment, due_date, payment_status, payment_method) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, order_id, virtual_account_number, total_payment, due_date, payment_status, payment_method, payment_time
	`, data.OrderID, virtualAccount, data.TotalPayment, dueDate, "Pending", data.PaymentMethod).Scan(
		&payment.ID,
		&payment.OrderID,
		&payment.VirtualAccountNumber,
		&payment.TotalPayment,
		&payment.DueDate,
		&payment.PaymentStatus,
		&payment.PaymentMethod,
		&payment.PaymentTime,
	)
	if err != nil {
		return payment, fmt.Errorf("error adding payment info: %v", err)
	}

	return payment, nil
}

func OrderById(iddb int) ([]OrderNew, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var orders []OrderNew

	rows, err := conn.Query(context.Background(), `
        SELECT id, date, time, movie_title, cinema_name, total_seat, total_price, payment, full_name, email, phone_number 
        FROM order_new WHERE user_id = $1;
    `, iddb)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order OrderNew
		err := rows.Scan(&order.Id, &order.Date, &order.Time, &order.Movie_Title, &order.Cinema_name, &order.Total_Seat,
			&order.Total_Price, &order.Payment, &order.Full_name, &order.Email, &order.Phone_number)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}

	if len(orders) == 0 {
		fmt.Println("No orders found for user with ID:", iddb)
	}

	return orders, nil
}
