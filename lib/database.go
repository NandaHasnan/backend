package lib

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func DB() *pgx.Conn {
	godotenv.Load()
	// DB := "postgresql://postgres:1@localhost:5432/faztix?sslmode=disable"
	// DB := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	DB, err := pgx.ParseConfig("")
	if err != nil {
		log.Println(err)
	}
	conn, err := pgx.Connect(context.Background(), DB.ConnString())
	if err != nil {
		fmt.Println(err)
	}
	return conn

	// godotenv.Load()
	// config, err := pgx.ParseConfig("")
	// if err != nil {
	// 	log.Println(err)
	// }

	// conn, err := pgx.Connect(context.Background(), config.ConnString())
	// if err != nil {
	// 	log.Println(err)
	// }

	// return conn
}
