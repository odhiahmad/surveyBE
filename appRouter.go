package main

import (
	"survey/api/connect"
	"survey/api/controllers"
	"survey/api/controllers/handler"
	"survey/api/manager"
	"survey/api/middlewares"
	"survey/api/utils/httpParse"
	"survey/api/utils/httpResponse"

	"github.com/gorilla/mux"
)

type appRouter struct {
	app                  *briApp
	parse                *httpParse.JsonParse
	responder            httpResponse.IResponder
	connect              connect.Connect
	logRequestMiddleware *middlewares.LogRequestMiddleware
}

type appRoutes struct {
	centerRoutes controllers.IDelivery
	mdw          []mux.MiddlewareFunc
}

func (r *appRouter) InitMainRoutes() {
	r.app.router.Use(r.logRequestMiddleware.Log)
	serviceManager := manager.NewServiceManager(r.connect)
	appRoutes := []appRoutes{
		{
			centerRoutes: handler.NewUserController(r.app.router, r.parse, r.responder, serviceManager.UserUseCase()),
			mdw:          nil,
		},
	}

	for _, r := range appRoutes {
		r.centerRoutes.InitRoute(r.mdw...)
	}
}

func NewAppRouter(app *briApp) *appRouter {
	return &appRouter{
		app,
		httpParse.NewJsonParse(),
		httpResponse.NewJSONResponder(),
		app.connect,
		middlewares.NewLogRequestMiddleware(),
	}
}
