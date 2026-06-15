package main

import (
	"log"
	"net/http"

	cf "backend/internal/config"
	d "backend/internal/database"
)

func main() {

	cfg := cf.Load()

	db := d.NewPostgresConnection(
		cfg.DatabaseURL,
	)

	defer db.Close()

	mux := http.NewServeMux()

	log.Println("server started on port", cfg.Port)


	err := http.ListenAndServe(":"+cfg.Port, mux)

	if err != nil {
		log.Fatal(err)
	}
}
