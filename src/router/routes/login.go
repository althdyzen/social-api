package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Rota{
	URI:      "/login",
	Method:   http.MethodPost,
	Function: controllers.Login,
	NeedAuth: false,
}
