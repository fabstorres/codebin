package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	AddRoute(register)
}

func register(a AppContext, r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}).Methods("GET")
}
