package models

import (
	"backend/lib"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type User struct {
	Id           int    `json:"id"`
	Firstname    string `json:"firstname" form:"firstname"`
	Lastname     string `json:"lastname" form:"lastname"`
	Phone_number string `json:"phone_number" form:"phone_number"`
	Image        string `json:"image" form:"image"`
	// Email        string `json:"email" form:"email"`
	// Password     string `json:"password" form:"password"`
	// Point int `json:"point" form:"point"`
}

type User_credentials struct {
	Id int `json:"id"`
	// User_id  int    `json:"user_id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Gabung struct {
	Id           *int   `json:"id" db:"id"`
	Firstname    string `json:"firstname" form:"firstname" db:"firstname"`
	Lastname     string `json:"lastname" form:"lastname" db:"lastname"`
	Phone_number string `json:"phone_number" form:"phone_number" db:"phone_number"`
	Image        string `json:"image" form:"image" db:"image"`
	Email        string `json:"email" form:"email" db:"email" binding:"required"`
	Password     string `json:"password" form:"password" db:"password"`
}

type ListUsers []User
type ListUsersGabung []Gabung

func UserAll() ListUsersGabung {
	conn := lib.DB()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `
		SELECT users.id, users.firstname, users.lastname, users.phone_number, users.image,
		user_credentials.email, user_credentials.password
		FROM users RIGHT JOIN user_credentials ON users.user_credentials_id = user_credentials.id
	`)
	if err != nil {
		fmt.Println("failed to execute query: %w", err)
	}

	userAll, err := pgx.CollectRows(rows, pgx.RowToStructByName[Gabung])
	if err != nil {
		fmt.Println("failed to collect rows: %w", err)
	}

	return userAll
}

// func UserAll() ListUsers {
// 	conn := lib.DB()
// 	defer conn.Close(context.Background())

// 	// var userList ListUsers

// 	rows, err := conn.Query(context.Background(), `
// 		SELECT id, firstname, lastname, phone_number, image,
// 		point, created_at, updated_at FROM users
// 	`)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	userall, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return userall
// }

func UserById(iddb int) Gabung {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var id int
	var firstname string
	var lastname string
	var phone_number string
	var image string
	var email string
	var password string
	// var point int
	// var user Gabung

	err := conn.QueryRow(context.Background(), `
	SELECT users.id, COALESCE(users.firstname, ''), COALESCE(users.lastname, ''), COALESCE(users.phone_number, ''), COALESCE(users.image, ''),
		COALESCE(user_credentials.email, ''), COALESCE(user_credentials.password, '')
		FROM users RIGHT JOIN user_credentials ON users.user_credentials_id = user_credentials.id WHERE users.id = $1
	`, iddb).Scan(&id, &firstname, &lastname, &phone_number, &image, &email, &password)

	if err != nil {
		// ctx.JSON(http.StatusBadRequest, controllers.TaskResponse{
		// 	Success: false,
		// 	Message: "invalid all user",
		// 	// Result:  id,
		// })

		fmt.Println(err)
	}
	// return user
	return Gabung{
		Id:           &id,
		Firstname:    firstname,
		Lastname:     lastname,
		Phone_number: phone_number,
		Image:        image,
		Email:        email,
		Password:     password,
		// Point: point,
	}
}

func UserByEmail(email string) User_credentials {

	conn := lib.DB()
	defer conn.Close(context.Background())

	var user User_credentials

	conn.QueryRow(context.Background(), `
	SELECT id, email, password FROM user_credentials WHERE email = $1
	`, email).Scan(&user.Id, &user.Email, &user.Password)

	return user

}

func InserUser(user User_credentials) User_credentials {
	conn := lib.DB()
	defer conn.Close(context.Background())
	var userNew User_credentials

	err := conn.QueryRow(context.Background(), `
	INSERT INTO user_credentials (email, password) 
	VALUES  ($1, $2)
	RETURNING  id, email, password
	`, user.Email, user.Password).Scan(
		&userNew.Id,
		&userNew.Email,
		&userNew.Password,
	)

	if err != nil {
		fmt.Println("inser error")
	}

	return userNew
}

func UpdateUser(user Gabung) Gabung {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var updatedUser Gabung

	if user.Email != "" || user.Password != "" {
		_, err := conn.Exec(context.Background(), `
			UPDATE user_credentials
			SET email = COALESCE($1, email),
			    password = COALESCE($2, password)
			WHERE id = $3
		`, user.Email, user.Password, user.Id)
		if err != nil {
			fmt.Println("Error updating user_credentials:", err)
			return Gabung{}
		}
	}

	var userCredentialsId int
	err := conn.QueryRow(context.Background(), `
		SELECT id FROM user_credentials WHERE id = $1
	`, user.Id).Scan(&userCredentialsId)

	fmt.Println("Received ID for user update:", user.Id)
	if err != nil || userCredentialsId == 0 {
		fmt.Println("Error: user_credentials not found or invalid ID:", user.Id)
		return Gabung{}
	}

	var count int
	err = conn.QueryRow(context.Background(), `
		SELECT COUNT(*) FROM users
		WHERE user_credentials_id = $1
	`, userCredentialsId).Scan(&count)
	if err != nil {
		fmt.Println("Error checking users table:", err)
		return Gabung{}
	}

	if count == 0 {
		_, err = conn.Exec(context.Background(), `
			INSERT INTO users (firstname, lastname, phone_number, image, user_credentials_id)
			VALUES ($1, $2, $3, $4, $5)
		`, user.Firstname, user.Lastname, user.Phone_number, user.Image, userCredentialsId)
		if err != nil {
			fmt.Println("Error inserting into users:", err)
			return Gabung{}
		}
	}

	if user.Firstname != "" || user.Lastname != "" || user.Phone_number != "" || user.Image != "" || user.Email != "" {
		err = conn.QueryRow(context.Background(), `
			UPDATE users
			SET firstname = COALESCE($1, firstname),
			    lastname = COALESCE($2, lastname),
			    phone_number = COALESCE($3, phone_number),
			    image = COALESCE($4, image)
			FROM user_credentials
			WHERE users.user_credentials_id = $5
			AND user_credentials.id = $5
			RETURNING users.id, users.firstname, users.lastname, users.phone_number, users.image, user_credentials.email, user_credentials.password
		`, user.Firstname, user.Lastname, user.Phone_number, user.Image, userCredentialsId).Scan(
			&updatedUser.Id,
			&updatedUser.Firstname,
			&updatedUser.Lastname,
			&updatedUser.Phone_number,
			&updatedUser.Image,
			&updatedUser.Email,
			&updatedUser.Password,
		)
		if err != nil {
			fmt.Println("Error updating users:", err)
			return Gabung{}
		}
	}

	return updatedUser
}

// func UpdateUser(user User) User {
// 	conn := lib.DB()
// 	defer conn.Close(context.Background())

// 	var updatedUser User

// 	// Menjalankan query update dengan penanganan error
// 	conn.QueryRow(context.Background(), `
// 		UPDATE users
// 		SET firstname = $1, lastname = $2, phone_number = $3, image = $4
// 		WHERE id = $5
// 		RETURNING id, firstname, lastname, phone_number, image, point, created_at, updated_at
// 	`, user.Firstname, user.Lastname, user.Phone_number, user.Image, user.Id).Scan(
// 		&updatedUser.Id,
// 		&updatedUser.Firstname,
// 		&updatedUser.Lastname,
// 		&updatedUser.Phone_number,
// 		&updatedUser.Image,
// 		// &updatedUser.Email,
// 		// &updatedUser.Password,
// 		// &updatedUser.Point,
// 	)

// 	return updatedUser
// }

func DeleteUser(iddb int) (Gabung, error) {
	conn := lib.DB()
	defer conn.Close(context.Background())

	var userDelete Gabung

	err := conn.QueryRow(context.Background(), `
		SELECT 
			users.id, 
			users.firstname, 
			users.lastname, 
			users.phone_number, 
			users.image, 
			user_credentials.email, 
			user_credentials.password
		FROM 
			users
		JOIN 
			user_credentials 
		ON 
			users.user_credentials_id = user_credentials.id
		WHERE 
			users.id = $1
	`, iddb).Scan(
		&userDelete.Id,
		&userDelete.Firstname,
		&userDelete.Lastname,
		&userDelete.Phone_number,
		&userDelete.Image,
		&userDelete.Email,
		&userDelete.Password,
	)

	if err != nil {
		log.Println("Error fetching user data:", err)
		return Gabung{}, fmt.Errorf("failed to fetch user data: %w", err)
	}

	_, err = conn.Exec(context.Background(), `
		DELETE FROM 
			user_credentials 
		USING 
			users 
		WHERE 
			user_credentials.id = users.user_credentials_id 
			AND users.id = $1
	`, iddb)

	if err != nil {
		log.Println("Error deleting user credentials:", err)
		return Gabung{}, fmt.Errorf("failed to delete user credentials: %w", err)
	}

	_, err = conn.Exec(context.Background(), `
		DELETE FROM 
			users 
		WHERE 
			id = $1
	`, iddb)

	if err != nil {
		log.Println("Error deleting user:", err)
		return Gabung{}, fmt.Errorf("failed to delete user: %w", err)
	}

	return userDelete, nil
}
