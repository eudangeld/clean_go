package web

import (
	"database/sql"
	"log"

	"github.com/codegangsta/negroni"
	"github.com/eudangeld/clean_go/core/coin"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main () {
	db, err := sql.Open("sqlite3", "./data/coin_test.db")	

	if err != nil {
		log.Fatal("Db connect failed: %s", err.Error())
	}

	defer db.Close()


	service := coin.NewService(db)

	r:= mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	r.Handle("/v1/coin", n.With()
  // n.UseHandler(r)
  // n.Run(":3000")




}