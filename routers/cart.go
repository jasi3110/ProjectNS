package routers

import (
	"OnlineShop/controllers/masters"
	"net/http"
	"github.com/gorilla/mux"
)

func CartRoutes(Router *mux.Router) *mux.Router {

	cartController := masters.CartController{}

	Router.Handle("/cart/create", http.HandlerFunc(cartController.CartCreate)).Methods(http.MethodPost)
	Router.Handle("/cart/update", http.HandlerFunc(cartController.CartUpdate)).Methods(http.MethodPost)
	Router.Handle("/cart/productdelete", http.HandlerFunc(cartController.CartProductDelete)).Methods(http.MethodPost)
	Router.Handle("/cart/deleteall",http.HandlerFunc(cartController.CartDelete)).Methods(http.MethodPost)
	Router.Handle("/cart/getall/{id}", http.HandlerFunc(cartController.CartGetAll)).Methods(http.MethodGet)
	
	return Router
}
