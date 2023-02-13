package routers

import (
	// "OnlineShop/controllers"
	"OnlineShop/controllers/masters"
	// "OnlineShop/repos/masterRepo"

	// "OnlineShop/utls"
	"net/http"

	"github.com/gorilla/mux"
)

func CartRoutes(Router *mux.Router) *mux.Router {

	cartController := masters.CartController{}

	Router.Handle("/cart/create", http.HandlerFunc(cartController.CartCreate)).Methods(http.MethodPost)
	Router.Handle("/category/update", http.HandlerFunc(cartController.CartUpdate)).Methods(http.MethodGet)
	Router.Handle("/category/getall/{id}", http.HandlerFunc(cartController.CartGetAll)).Methods(http.MethodGet)
	Router.Handle("/category/delete", http.HandlerFunc(cartController.CartDelete)).Methods(http.MethodPost)

	return Router
}
