package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	server := http.Server{
		Handler: app.routes(),
		Addr:    *addr,
	}

	infoLog.Printf("Starting server on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
