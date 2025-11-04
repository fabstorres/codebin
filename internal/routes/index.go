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
		a.Renderer().Render(w, "home/index.html", nil)
	}).Methods("GET")
}
