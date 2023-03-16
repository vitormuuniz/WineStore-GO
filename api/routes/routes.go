package routes

import (
	"net/http"

	"github.com/vitormuuniz/winestore-go/api/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func Install(router *mux.Router, wineStoreRoutes WineStoreRoutes) {
	allRoutes := wineStoreRoutes.Routes()

	for _, route := range allRoutes {
		handler := middlewares.Logger(route.Handler)
		router.HandleFunc(route.Path, handler).Methods(route.Method)
	}
}
