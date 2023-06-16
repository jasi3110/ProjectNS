package routers

import (
	"OnlineShop/controllers"
	"OnlineShop/controllers/masters"
	"OnlineShop/utls"
	"net/http"
	"github.com/gorilla/mux"
)

func UnitRoutes(Router *mux.Router) *mux.Router {

	unitController := masters.UnitController{}

	Router.Handle("/unit/create", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(unitController.UnitCreate)))).Methods(http.MethodPost)
	Router.Handle("/unit/getbyid/{id}", http.HandlerFunc(unitController.UnitGetById)).Methods(http.MethodGet)
	Router.Handle("/unit/getall", http.HandlerFunc(unitController.UnitGetAll)).Methods(http.MethodGet)
	Router.Handle("/unit/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(unitController.UnitUpdate)))).Methods(http.MethodPost)

	return Router

}
