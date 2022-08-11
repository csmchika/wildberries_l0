package postgres

import (
	"database/sql"
	"encoding/json"
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
		var data []byte
		rows.Scan(&i, &data)
		recievedModel := models.Model{}
		_ = json.Unmarshal(data, &recievedModel)
		m[i-1] = &recievedModel
	}
	defer rows.Close()
	return m
}

func (db *DB) AddModelDB(data []byte) (int, error) {
	var id int
	err := db.Open.QueryRow("INSERT INTO models (data) VALUES ($1) RETURNING uid", data).Scan(&id)
	return id, err
}
