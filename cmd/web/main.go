package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorlog *log.Logger
	infolog  *log.Logger
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{
		errorlog: errorLog, infolog: infoLog,
	}
	addr := flag.String("addr", ":4000", "HTTP network port")
	flag.Parse()

	infoLog.Println("Starting server")
	log.Printf("starting server on %s", *addr)

	srv := http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
