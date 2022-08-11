package postgres

import (
	"database/sql"
	"log"
	"wb-l0/internal/postgres/models"
)

type DB struct {
	Open *sql.DB
}

func NewDB() *DB {
	db := DB{}
	db.Init()
	return &db
}

func (db *DB) SetCache(c *Cache) map[int]*models.Model {
	m := make(map[int]*models.Model)
	rows, err := db.Open.Query("SELECT uid, data FROM models;")
	if err != nil {
		log.Panic(err)
		return make(map[int]*models.Model)
	}
	for rows.Next() {
		var i int
		var data *models.Model
		rows.Scan(&i, &data)
		m[i] = data
	}
	defer rows.Close()
	return m
}

func (db *DB) AddModelDB(data []byte) (int, error) {
	var id int
	err := db.Open.QueryRow("INSERT INTO models (data) VALUES ($1) RETURNING uid", data).Scan(&id)
	return id, err
}
