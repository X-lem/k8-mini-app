package main

import (
	"log"
	"net/http"

	"github.com/X-lem/k8-app/server/routes"
	"github.com/X-lem/k8-app/server/shared"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger) // This will log the requests so you can see them coming in!

	// Attempt to get the DB
	db := shared.GetDB()

	// Setup all the routes we can
	setupRoutes(r, db)

	// Start the application
	log.Println("k8-mini-app started")
	http.ListenAndServe(":8080", r)
}

func setupRoutes(r *chi.Mux, db *pgx.Conn) {
	routes.Routes01(r)

	if db != nil {
		routes.Routes02(r, db)
	}
}
