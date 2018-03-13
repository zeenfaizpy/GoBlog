package config

import (
	"database/sql"

	_ "github.com/lib/pq" // Pgsql driver import
)

// Env ..
type Env struct {
	DB *sql.DB
}

// NewDB ..
func NewDB(driverSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", driverSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
