package main

import (
	"log"
	"net/http"

	cf "backend/internal/config"
	d "backend/internal/database"
	"backend/router"
	app "backend/internal/app"
)

func main() {

	cfg := cf.Load()

	db := d.NewPostgresConnection(cfg.DatabaseURL)
	defer db.Close()

	app := app.New(db)

	mux := http.NewServeMux()
	router.SetupRoutes(mux, app.Router)

	log.Println("server started on port", cfg.Port)

	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}