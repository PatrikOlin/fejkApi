package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/PatrikOlin/fejkApi/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	t := time.Now()
	logpath := filepath.Join(".", "logs")
	os.MkdirAll(logpath, os.ModePerm)
	logFile, err := os.OpenFile(logpath+"fejkApi_"+t.Format("20060102")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	err = db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	api := r.PathPrefix("/v1").Subrouter()
	api.HandleFunc("/articles", getArticle).Methods(http.MethodGet)
	api.HandleFunc("/people", getPerson).Methods(http.MethodGet)
	api.HandleFunc("/companies", getCompany).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8124", handlers.CombinedLoggingHandler(logFile, handlers.CORS()(r))))

}
