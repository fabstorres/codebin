package server

import (
	"fmt"
	"net/http"

	"github.com/fabstorres/codebin/internal/config"
	"github.com/fabstorres/codebin/internal/render"
	"github.com/fabstorres/codebin/internal/routes"
	"github.com/gorilla/mux"
)

type App struct {
	Cfg       *config.Config
	Server    *http.Server
	Templates *render.Renderer
}

func New() *App { return &App{} }

func (a *App) Config() *config.Config     { return a.Cfg }
func (a *App) Renderer() *render.Renderer { return a.Templates }

func (a *App) Initialize() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}
	a.Cfg = cfg

	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./web/static"))),
	)

	// initialize the templates once
	renderer, err := render.New("web/templates")
	if err != nil {
		return fmt.Errorf("load templates: %w", err)
	}
	a.Templates = renderer

	routes.RegisterRoutes(a, r)

	a.Server = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}
	return nil
}
