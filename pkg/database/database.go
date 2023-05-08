package database

import "database/sql"

type Database struct {
	dbDriver *sql.DB
}

func New(db *sql.DB) *Database {
	return &Database{
		dbDriver: db,
	}
}
