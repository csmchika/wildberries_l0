package main

import (
	// "database/sql"
	// "fmt"

	// _ "github.com/lib/pq"
	"log"
	"os"
	"wb-l0/cmd/config"

	"github.com/nats-io/stan.go"
)

func main() {
	config.ConfigSetup()
	// dbObject := postgres.NewDB()
	connect, er := stan.Connect(os.Getenv("NATS_CLUSTER_ID"), os.Getenv("NATS_CLIENT_ID"))
	if er != nil {
		log.Fatal("не удалось подключиться к nats-streaming-server")
	}
	_, err := connect.Subscribe(os.Getenv("NATS_SUBJECT"),
		func(m *stan.Msg) {
			log.Printf("YOU received a message! %s\n", m.Data)
			return
		})
	if err != nil {
		panic(err)
	}
}

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
