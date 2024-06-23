package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request){
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
