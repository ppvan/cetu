package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/ppvan/cetu/internal/models"
)

var version = "1.0.0"

type config struct {
	env    string
	port   int
	domain string
	db     struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  time.Duration
	}
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	urlModel *models.URLModel
	config   *config
}

func main() {

	var config config
	flag.StringVar(&config.domain, "domain", "localhost", "Domain name")
	flag.StringVar(&config.env, "env", "development", "Environment (development|production)")
	flag.IntVar(&config.port, "port", 4000, "Port number")

	// DB settings
	flag.StringVar(&config.db.dsn, "dsn", os.Getenv("CETU_DB"), "Postgres data source name")
	flag.IntVar(&config.db.maxOpenConns, "db-max-open-conns", 25, "Postgres max open connections")
	flag.IntVar(&config.db.maxIdleConns, "db-max-idle-conns", 25, "Postgres max idle connections")
	flag.DurationVar(&config.db.maxIdleTime, "db-max-idle-time", 15*time.Minute, "Postgres max connection idle time")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(config)
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
		Addr:         fmt.Sprintf("%s:%d", config.domain, config.port),
		IdleTimeout:  5 * time.Minute,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	infoLog.Printf("Starting server on %s", server.Addr)
	// err = server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	err = server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	fmt.Println(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
