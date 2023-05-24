package routers

import (
	"OnlineShop/controllers/dashboard"
	"net/http"
	"github.com/gorilla/mux"
)

func DashBoardRoutes(Router *mux.Router) *mux.Router {

	DashBoardController := dashboard.Dashboard{}
	DashBoardControllerCart:=dashboard.DashboardCart{}

	Router.Handle("/dashboard/homepage",http.HandlerFunc(DashBoardController.Homepage)).Methods(http.MethodGet)
	Router.Handle("/dashboard/cartpage/{id}",http.HandlerFunc(DashBoardControllerCart.Cartpage)).Methods(http.MethodGet)
	Router.Handle("/dashboard/allimage",http.HandlerFunc(DashBoardController.ProductImageGetAll)).Methods(http.MethodGet)
	
	return Router

}