package main

import (
	// "database/sql"
	// "fmt"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"wb-l0/cmd/config"
	postgres "wb-l0/internal/postgres"
	"wb-l0/internal/postgres/models"

	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"

	"github.com/nats-io/stan.go"
)

func main() {
	config.ConfigSetup()
	dbObject := postgres.NewDB()
	cache := postgres.NewCache(dbObject)
	log.Printf("count elems in cache %d", len(cache.CacheModel))
	connect, er := stan.Connect(os.Getenv("NATS_CLUSTER_ID"), os.Getenv("NATS_CLIENT_ID"))
	if er != nil {
		log.Fatal("не удалось подключиться к nats-streaming-server")
	}
	log.Printf("connect good")
	_, err := connect.Subscribe(os.Getenv("NATS_SUBJECT"),
		func(m *stan.Msg) {
			log.Printf("You received a message!\n")
			if messageHandler(cache, dbObject, m.Data) {
				err := m.Ack() // в случае успешного сохранения msg уведомляем NATS.
				if err != nil {
					log.Printf("ack() err: %s", err)
				}
			}
		},
		stan.AckWait(time.Duration(30)*time.Second),      // Интервал тайм-аута - AckWait (30 сек default) - ожидание уведомления NATS о чтении сообщения
		stan.DeliverAllAvailable(),                       // DeliverAllAvailable доставит все доступные сообщения
		stan.DurableName(os.Getenv("NATS_DURABLE_NAME")), // долговечные подписки позволяют клиентам назначить постоянное имя подписке
		// Это приводит к тому, что сервер потоковой передачи NATS отслеживает последнее подтвержденное сообщение для этого clientID + постоянное имя,
		// так что клиенту будут доставлены только сообщения с момента последнего подтвержденного сообщения.
		stan.SetManualAckMode(), // ручной режим подтверждения приема сообщения для подписки
		stan.MaxInflight(10))
	if err != nil {
		panic(err)
	}
	log.Printf("subscribe end")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func CheckError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
func messageHandler(c *postgres.Cache, db *postgres.DB, data []byte) bool {
	recievedModel := models.Model{}
	var Validator = validator.New()
	log.Printf("Создал, но не обработал")
	err := json.Unmarshal(data, &recievedModel)
	if err != nil {
		log.Printf("error, %v\n", err)
	}
	err = Validator.Struct(recievedModel)
	if err != nil {
		log.Printf("messageHandler() error, %v\n", err)
		// ошибка формата присланных данных. Пропускаем, сообщив серверу, что сообщение получили
		return false
	}
	log.Printf("unmarshal Order to struct")
	var id int64
	err = db.Open.QueryRow("INSERT INTO models (data) VALUES ($1) RETURNING uid", data).Scan(&id)
	if err != nil {
		log.Printf("Error adding model %v\n", err)
		return false
	}
	c.AddModel(id, &recievedModel)
	log.Printf("There are %d elems in cache", c.CountElems())
	return true
}
