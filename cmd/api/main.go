package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ppvan/cetu/internal/models"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	urlModel *models.URLModel
	config   *Config
}

func main() {
	config := ParseConfig()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(config.DSN)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	urlModel := &models.URLModel{DB: db}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		urlModel: urlModel,
		config:   &config,
	}

	server := http.Server{
		Handler:      app.routes(),
		Addr:         fmt.Sprintf("%s:%s", config.Domain, config.Port),
		IdleTimeout:  5 * time.Minute,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	infoLog.Printf("Starting server on %s", server.Addr)
	err = server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
