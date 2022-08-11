package postgres

import (
	"database/sql"
	"log"
	"wb-l0/internal/postgres/models"
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

func (db *DB) SetCache(c *Cache) map[int64]*models.Model {
	m := make(map[int64]*models.Model)
	rows, err := db.Open.Query("SELECT uid, data FROM models;")
	if err != nil {
		log.Panic(err)
		return make(map[int64]*models.Model)
	}
	for rows.Next() {
		var i int64
		var data *models.Model
		rows.Scan(&i, &data)
		log.Printf("%d, id %+v data", i, data)
		m[i] = data
		// err = query.Scan(&modelData)
		// if err != nil {
		// 	log.Panic(err)
		// }
		// bk.JModelSlice.Lock()
		// err := bk.JModelSlice.AddFromData(jsonData)
		// bk.JModelSlice.Unlock()
		// if err != nil {
		// 	log.Println(err)
		// 	return
		// }
	}
	defer rows.Close()
	return m
}
