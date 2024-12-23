package lib

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func DB() *pgx.Conn {
	// godotenv.Load()
	DB := "postgresql://postgres:1@localhost:5432/faztix?sslmode=disable"
	// DB := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	conn, err := pgx.Connect(context.Background(), DB)
	if err != nil {
		fmt.Println(err)
	}
	return conn
}
