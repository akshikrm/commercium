package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *storage {
	return &storage{
		db: db,
	}
}
