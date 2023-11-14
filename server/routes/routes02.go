package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/X-lem/k8-app/server/shared"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5"
)

type CreateUserRequest struct {
	Username string `json:"username"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}

func Routes02(r *chi.Mux, db *pgx.Conn) {
	r.Post("/create-table", func(w http.ResponseWriter, r *http.Request) {

		_, err := db.Exec(context.Background(), `
			CREATE TABLE users (
				ID SERIAL NOT NULL PRIMARY KEY,
				Username VARCHAR UNIQUE NOT NULL,
				CreatedAt TIMESTAMP NOT NULL DEFAULT NOW()
			);
		`)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "users table created!")
	})

	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		var input CreateUserRequest
		if err := shared.RequestBodyParser(r.Body, &input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		row := db.QueryRow(context.Background(), `
			INSERT INTO users (username)
			VALUES ($1)
			RETURNING ID;
		`, input.Username)

		var userID int
		err := row.Scan(&userID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		row = db.QueryRow(context.Background(), `
			SELECT * FROM users
			WHERE id = $1;
		`, userID)

		var user User
		err = row.Scan(
			&user.ID,
			&user.Username,
			&user.CreatedAt,
		)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
	})

	r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(context.Background(), `SELECT * FROM users;`)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var users []User
		for rows.Next() {
			var user User
			err = rows.Scan(
				&user.ID,
				&user.Username,
				&user.CreatedAt,
			)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		json.NewEncoder(w).Encode(users)
	})
}
