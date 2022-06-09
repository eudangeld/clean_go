package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/eudangeld/clean_go/core/coin"
	"github.com/eudangeld/clean_go/core/user"
	"github.com/eudangeld/clean_go/web/handlers"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main () {
	db, err := sql.Open("sqlite3", "./data/coin.db")	

	if err != nil {
		log.Fatalf("Db connect failed: %s", err.Error())
	}

	defer db.Close()


	coinService := coin.NewService(db)
	userService := user.NewUserService(db)

	r:= mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	
	//init handler for cois
	handlers.MakeCoinHandlers(r, n, coinService)
	handlers.MakeUserHandler(r, n, userService)
	http.Handle("/", r)

	port:=os.Getenv("PORT")

	if port == "" {
		log.Fatal("Server port not defined exitting")
	}
	
	
	srv := &http.Server{
		ReadTimeout: 30* time.Second,
		WriteTimeout: 30* time.Second,
		Addr: ":"+port,
		Handler: http.DefaultServeMux,
		ErrorLog: log.New(os.Stderr,"Logger", log.Lshortfile),
	}
	
	log.Println("Server started at:localhost:" + port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	

  
}

