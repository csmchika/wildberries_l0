package serve

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"wb-l0/internal/postgres"

	"github.com/gorilla/mux"
)

type Server struct {
	Cache  *postgres.Cache
	Router *mux.Router
}

type ModelsToHtml struct {
	Model string
	Count int
}

func NewServer(c *postgres.Cache) *Server {
	Server := Server{}
	Server.Init(c)
	return &Server
}

func (s *Server) Init(c *postgres.Cache) {
	s.Cache = c
	s.Router = mux.NewRouter()

	s.Router.HandleFunc("/", s.GetModels).Methods("GET")

	log.Printf("Запуск сервера на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}

func (s *Server) GetModels(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("serve/tmpl/index.html")
	if err != nil {
		log.Fatalf("Ошибка в открытии шаблона %v", err)
		return
	}
	value := r.FormValue("model_id")
	model_id, ok := strconv.Atoi(value)
	temp_data := ModelsToHtml{"", s.Cache.CountElems()}
	if value == "" {
		temp_data.Model = ""
	} else if ok != nil {
		temp_data.Model = "Похоже, что введено не число"
	} else if model_id > s.Cache.CountElems() || model_id <= 0 {
		temp_data.Model = "Такого id ещё нет\n" +
			fmt.Sprintf("Введите ID в диапазоне от 1 до %d", s.Cache.CountElems())
	} else {
		temp_data.Model = s.Cache.ToString(model_id - 1)
	}
	_ = t.Execute(w, temp_data)
}

func (s *Server) Close() {
	log.Printf("Сервера нет, это же просто роутер...")
}
