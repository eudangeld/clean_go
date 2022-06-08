package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/eudangeld/clean_go/core/coin"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main () {
	db, err := sql.Open("sqlite3", "./data/coin.db")	

	if err != nil {
		log.Fatalf("Db connect failed: %s", err.Error())
	}

	defer db.Close()


	service := coin.NewService(db)

	r:= mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	r.Handle("/v1/coin", n.With(
		negroni.Wrap(coinHandler(service)),
	)).Methods("GET", "OPTIONS")
	http.Handle("/", r)
	srv := &http.Server{
		ReadTimeout: 30* time.Second,
		WriteTimeout: 30* time.Second,
		Addr: ":4000",
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr,"Logger", log.Lshortfile),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
  
}


func coinHandler(service coin.UseCase) http.Handler {
	return http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		books, _ := service.GetAll()
		for _, b := range books {
			fmt.Println(b)
		}
	})
}
