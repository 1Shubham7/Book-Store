package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"

	"github.com/gorilla/mux"
	"net/http"
)

type health struct {
	Status string `json:"status"`
	Messages []string `json:"messages"`
}

type jsonError struct{
	Code string `json:"code"`
	Msg string `json:"msg"`
}

func main() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error making DB connected: %s", err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error making DB driver: %s", err.Error())
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Error making migration engine: %s", err.Error())
	}
	migrator.Steps(2)

	
	r := mux.NewRouter()

	r.HandleFunc(
		"/healthcheck", 
		func(w http.ResponseWriter, r *http.Request){
			h := health{
				Status : "OK",
				Messages: []string{},
			}
			b, _ := json.Marshal(h)
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			
	})

	r.HandleFunc(
		"/book/{isbn}",
		func(w http.ResponseWriter, r *http.Request){
			v := mux.Vars(r)
			e := jsonError{
				code: "001",
				Msg: fmt.Sprintf("No books with ISBN %s", v["isbn"]),
			}

			b,_ := json.Marshal(e)
			q := jsonError
			w.WriteHeader(http.StatusNotFound)
			w.Write(b)
		},
	)

	s := http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 5 * time.Second,
		writeTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
