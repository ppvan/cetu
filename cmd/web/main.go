package main

import (
	"log"
	"net/http"
)

func main() {

	server := http.Server{
		Handler: routes(),
		Addr:    ":8080",
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
