package route

import (
	"net/http"

	"github.com/VictorCometti/api-atendimento/src/controllers"
)

var serviceRoutes = []Route{
	{
		URI:    "/service",
		Method: http.MethodGet,
		Func:   controllers.FindServiceByIdentifier,
	},
}
