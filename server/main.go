package main

import (
	"log"
	"net/http"

	"github.com/X-lem/k8-app/server/routes"
	"github.com/X-lem/k8-app/server/shared"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger) // This will log the requests so you can see them coming in!
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
	}))

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
