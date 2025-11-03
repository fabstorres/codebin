package server

import (
	"fmt"
	"net/http"

	"github.com/fabstorres/codebin/internal/config"
	"github.com/fabstorres/codebin/internal/routes"
	"github.com/gorilla/mux"
)

type App struct {
	Cfg    *config.Config
	Server *http.Server
}

func New() *App {
	return &App{}
}

func (a *App) Config() *config.Config {
	return a.Cfg
}

func (a *App) Initalize() error {

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	routes.RegisterRoutes(a, r)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}

	a.Cfg = cfg
	a.Server = srv

	return nil
}
