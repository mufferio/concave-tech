package utilities

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// NewDB returns a ready-to-use connection pool.
func NewDB() *sql.DB {
	// Allow the host name to come from an env-var so the same binary
	// works both on your laptop (localhost) and inside Docker (postgres service).
	host := os.Getenv("PGHOST")
	if host == "" {
		host = "localhost"
	}

	dsn := fmt.Sprintf(
		"postgres://devuser:devpass@%s:5432/mydatabase?sslmode=disable",
		host,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("opening DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("ping DB: %v", err)
	}
	return db
}
