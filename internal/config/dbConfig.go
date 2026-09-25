package config

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBconfig() (*pgxpool.Pool, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	source := os.Getenv("DB_DATA")
	dbPort := os.Getenv("DB_PORT")

	connect := "postgres://" + user + ":" + pass + "@localhost:" + dbPort + "/" + source

	db, err := pgxpool.New(context.Background(), connect)

	if err != nil {
		log.Println("gagal connect :", err)
		return nil, err
	}

	return db, nil
}
