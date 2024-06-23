package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type health struct {
	Status string `json:"status"`
	Messages []string `json:"messages"`
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
			w.Write(b)
			w.WriteHeader(http.StatusOK)
	})

	s := http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 5 * time.Second,
		writeTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
