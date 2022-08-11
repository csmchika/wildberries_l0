package postgres

import (
	"database/sql"
)

type DB struct {
	Open *sql.DB
	// csh  *Cache
}

func NewDB() *DB {
	db := DB{}
	db.Init()
	return &db
}
