package main

import (
	"io/ioutil"
	"log"
	"os"
	"wb-l0/cmd/config"

	"github.com/nats-io/stan.go"
)

func ReadAll(path string) (res *[][]byte, err error) {
	res = new([][]byte)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, v := range dir {
		if !v.IsDir() {
			file, err := ioutil.ReadFile(path + "/" + v.Name())
			if err == nil {
				*res = append(*res, file)
			}
		}
	}
	return res, err
}

func main() {
	config.ConfigSetup()
	repeat := 1
	models, err := ReadAll("./jsons")
	if err != nil {
		log.Panicf("Ошибка во время чтения файлов %v", err)
		return
	}
	connect, _ := stan.Connect(os.Getenv("NATS_CLUSTER_ID"), os.Getenv("NATS_CLIENT_ID"))
	for repeat > 0 {
		for _, v := range *models {
			err = connect.Publish(os.Getenv("NATS_SUBJECT"), v)
			if err != nil {
				log.Panicf("Ошибка отправки %v", err)
				return
			}
			log.Println("\rОтправил файлы")
		}
		repeat--
	}
}
