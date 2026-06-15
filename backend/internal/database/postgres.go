package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresConnection(
	databaseURL string,
) *pgxpool.Pool {

	db, err := pgxpool.New(
		context.Background(),
		databaseURL,
	)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping(context.Background())

	if err != nil {
		log.Fatal("database ping failed:", err)
	}

	log.Println("connected to PostgreSQL")

	return db
}
