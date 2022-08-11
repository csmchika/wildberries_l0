package postgres

import (
	"database/sql"
	"fmt"

	"log"
	"os"

	_ "github.com/lib/pq"
)

func (db *DB) Init() {
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	var err error
	db.Open, err = sql.Open("postgres", psqlconn)

	if err != nil {
		log.Fatalf("%v: Init() error: %s\n", os.Getenv("DB_NAME"), err)
	}
}
func (db *DB) Disconnect() {
	err := db.Open.Close()
	if err != nil {
		log.Panic(err)
	} else {
		log.Println("db disconnect")
	}
}

// func (db *DB) GetRaw() *sql.DB {
// 	return db.open
// }
