package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI    string
	Method string
	Func   func(w http.ResponseWriter, r *http.Request)
}

func ConfigureRoute(r *mux.Router) (router *mux.Router) {
	routes := serviceRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}
	return

}
