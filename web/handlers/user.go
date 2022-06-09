package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/eudangeld/clean_go/core/user"
	"github.com/gorilla/mux"
)

func MakeUserHandler(r *mux.Router, n *negroni.Negroni, service user.UseCase) {
	r.Handle("/v1/user/login", n.With(
		negroni.Wrap(logIn(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/v1/user/signup", n.With(
		negroni.Wrap(signup(service)),
	)).Methods("POST", "OPTIONS")
}



type UserDto struct {
	Email string `json:"email"`
	Pass string `json:"pass"`
}
func logIn(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			
			
			w.Header().Set("Content-Type", "application/json")
			var user UserDto

			err := json.NewDecoder(r.Body).Decode(&user)
			all, err := service.Login(user.Email, user.Pass)
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


func signup(service user.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user user.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = service.Signup(user.Email, user.Pass,user.Role)
		if err != nil {
			w.WriteHeader(http.StatusAccepted)
			return
		}
	})
}




