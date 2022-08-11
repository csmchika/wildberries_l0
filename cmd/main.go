package main

import (
	// "database/sql"
	// "fmt"

	"encoding/json"
	"log"
	"os"
	"time"
	"wb-l0/cmd/config"
	postgres "wb-l0/internal/postgres"
	"wb-l0/internal/postgres/models"
	"wb-l0/server"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"

	"github.com/nats-io/stan.go"
)

func main() {
	config.ConfigSetup()
	dbObject := postgres.NewDB()
	cache := postgres.NewCache(dbObject)
	log.Printf("В кеше сейчас %d элементов", cache.CountElems())

	connect, er := stan.Connect(os.Getenv("NATS_CLUSTER_ID"), os.Getenv("NATS_CLIENT_ID"))
	if er != nil {
		log.Fatal("Не удалось подключиться к nats-streaming-server")
	}

	log.Printf("Подключение к БД и nats-streaming успешно")
	_, err := connect.Subscribe(os.Getenv("NATS_SUBJECT"),
		func(m *stan.Msg) {
			log.Printf("Ты получил сообщение!\n")
			if messageHandler(cache, dbObject, m.Data) {
				err := m.Ack()
				if err != nil {
					log.Printf("Ошибка ответа: %s", err)
				}
			}
		},
		stan.AckWait(time.Duration(30)*time.Second),
		stan.DeliverAllAvailable(),
		stan.DurableName(os.Getenv("NATS_DURABLE_NAME")),
		stan.SetManualAckMode(),
		stan.MaxInflight(10))
	if err != nil {
		log.Fatal("Ошибка подписки на канал")
	}
	Server := server.NewServer(cache)
	defer Server.Close()
	defer dbObject.Disconnect()
}

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

func messageHandler(c *postgres.Cache, db *postgres.DB, data []byte) bool {
	recievedModel := models.Model{}
	var Validator = validator.New()
	err := json.Unmarshal(data, &recievedModel)
	if err != nil {
		log.Printf("Ошибка во время приведения типов %v\n", err)
	}
	err = Validator.Struct(recievedModel)
	if err != nil {
		log.Printf("Ошибка валидации данных %v\n", err)
		return true
	}
	id, err := db.AddModelDB(data)
	if err != nil {
		log.Printf("Ошибка во время добавления данных в БД %v\n", err)
		return false
	}
	c.AddModelCache(id, &recievedModel)
	log.Printf("В кеше сейчас %d элементов", c.CountElems())
	return true
}
