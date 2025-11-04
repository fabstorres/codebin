package routes

import (
	"github.com/fabstorres/codebin/internal/config"
	"github.com/fabstorres/codebin/internal/render"
	"github.com/gorilla/mux"
)

type AppContext interface {
	Config() *config.Config
	Renderer() *render.Renderer
}

type RouteRegistrar func(AppContext, *mux.Router)

var registry []RouteRegistrar

func AddRoute(fn RouteRegistrar) {
	registry = append(registry, fn)
}

func RegisterRoutes(a AppContext, r *mux.Router) {
	for _, reg := range registry {
		reg(a, r)
	}
}
