package main

import (
	"database/sql"
	"flag"
	"learning/pkg/models/mysql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // New import
)

type application struct {
	errorlog *log.Logger
	infolog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	addr := flag.String("addr", ":4000", "HTTP network port")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL DSN")
	flag.Parse()

	infoLog.Println("Starting server")
	log.Printf("starting server on %s", *addr)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	app := &application{
		errorlog: errorLog, infolog: infoLog, snippets: &mysql.SnippetModel{DB: db},
	}
	defer db.Close()

	srv := http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
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
