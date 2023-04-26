package routers

import (
	"OnlineShop/controllers/dashboard"
	"net/http"
	"github.com/gorilla/mux"
)

func DashBoardRoutes(Router *mux.Router) *mux.Router {

	DashBoardController := dashboard.Dashboard{}

	Router.Handle("/dashboard/homepage",http.HandlerFunc(DashBoardController.Homepage)).Methods(http.MethodGet)
	
	return Router

}