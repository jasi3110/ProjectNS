package routers

import (
	// "OnlineShop/controllers"
	"OnlineShop/controllers"
	"OnlineShop/controllers/masters"
	"OnlineShop/utls"

	// "OnlineShop/repos/masterRepo"

	// "OnlineShop/utls"
	"net/http"

	"github.com/gorilla/mux"
)

func CartRoutes(Router *mux.Router) *mux.Router {

	cartController := masters.CartController{}

	Router.Handle("/cart/create", http.HandlerFunc(cartController.CartCreate)).Methods(http.MethodPost)
	Router.Handle("/cart/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(cartController.CartUpdate)))).Methods(http.MethodPost)
	Router.Handle("/cart/productdelete", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(cartController.CartProductDelete)))).Methods(http.MethodPost)
	Router.Handle("/cart/deleteall",http.HandlerFunc(cartController.CartDelete)).Methods(http.MethodPost)

	Router.Handle("/cart/getall/{id}", http.HandlerFunc(cartController.CartGetAll)).Methods(http.MethodGet)
	
	return Router
}
