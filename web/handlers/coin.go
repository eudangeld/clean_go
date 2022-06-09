package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/eudangeld/clean_go/core/coin"
	"github.com/gorilla/mux"
)

func MakeCoinHandlers(r *mux.Router, n *negroni.Negroni, service coin.UseCase) {
	r.Handle("/v1/coin", n.With(
		negroni.Wrap(getAllCoin(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/v1/coin", n.With(
		negroni.Wrap(insertCoin(service)),
	)).Methods("POST", "OPTIONS")
}



func getAllCoin(service coin.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			
			
			w.Header().Set("Content-Type", "application/json")
			all, err := service.GetAll()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			
			err = json.NewEncoder(w).Encode(all)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		})
}


func insertCoin(service coin.UseCase) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Println("Post new")
			var coin coin.Coin

			err:= json.NewDecoder(r.Body).Decode(&coin)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			err = service.Store(&coin);

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)
		})
}


