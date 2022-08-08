package postgres

import (
	"database/sql"
)

type DB struct {
	open *sql.DB
	// csh  *Cache
}

func NewDB() *DB {
	db := DB{}
	db.Init()
	return &db
}
