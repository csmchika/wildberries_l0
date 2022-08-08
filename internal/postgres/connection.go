package postgres

import (
	"database/sql"
	"fmt"

	"log"
	"os"

	_ "github.com/lib/pq"
)

func (db *DB) Init() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	var err error
	db.open, err = sql.Open("postgres", psqlconn)

	if err != nil {
		log.Fatalf("%v: Init() error: %s\n", os.Getenv("DB_NAME"), err)
	}
}
func (db *DB) Disconnect() {
	err := db.open.Close()
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("db disconnect\n")
	}
}
