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
	Router.Handle("/cart/update", http.HandlerFunc(cartController.CartUpdate)).Methods(http.MethodPost)
	Router.Handle("/cart/getall/{id}", http.HandlerFunc(cartController.CartGetAll)).Methods(http.MethodGet)
	Router.Handle("/cart/productdelete", http.HandlerFunc(cartController.CartProductDelete)).Methods(http.MethodPost)
	Router.Handle("/cart/delete", http.HandlerFunc(cartController.CartDelete)).Methods(http.MethodPost)

	return Router
}
