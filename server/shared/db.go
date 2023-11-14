package shared

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func GetDB() *pgx.Conn {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbURL := fmt.Sprintf("postgres://postgres:%s@%s:5432/k8-mini-app", os.Getenv("database.password"), os.Getenv("postgres-service-name"))

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil
	}

	// Attempt to query db
	var greeting string
	err = conn.QueryRow(context.Background(), "SELECT 'Hello, From K8 Mini App DB'").Scan(&greeting)
	if err != nil {
		log.Printf("Test sql command failed: %v\n", err)
		return nil
	}

	log.Println(greeting)

	return conn
}
