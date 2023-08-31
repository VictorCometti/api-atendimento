package router

import (
	"github.com/VictorCometti/api-atendimento/src/router/route"
	"github.com/gorilla/mux"
)

func GetRouter() (r *mux.Router) {
	r = mux.NewRouter()
	route.ConfigureRoute(r)
	return
}
