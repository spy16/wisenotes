package storage

import (
	"database/sql"
	_ "embed"

	"github.com/spy16/wisenotes/storage/db"
)

//go:embed schema.sql
var schema string

// Open opens a new database connection, runs migrations and returns a new
// Queries instance.
func Open(spec string) (*db.Queries, func(), error) {
	dbConn, err := sql.Open("libsql", spec)
	if err != nil {
		return nil, nil, err
	}

	if err := migrate(dbConn); err != nil {
		_ = dbConn.Close()
		return nil, nil, err
	}

	return db.New(dbConn), func() { _ = dbConn.Close() }, nil
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(schema)
	return err
}
