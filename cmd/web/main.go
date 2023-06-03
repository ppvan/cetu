package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ppvan/cetu/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	urlModel *models.URLModel
}

func main() {

	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "cetu:cetu@/cetu?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	urlModel := &models.URLModel{DB: db}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		urlModel: urlModel,
	}

	server := http.Server{
		Handler: app.routes(),
		Addr:    *addr,
	}

	infoLog.Printf("Starting server on %s", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
