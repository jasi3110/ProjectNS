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
	// Router.Handle("/category/getbyid/{id}", http.HandlerFunc(categoryController.CategoryGetById)).Methods(http.MethodGet)
	Router.Handle("/category/getall", http.HandlerFunc(cartController.CartGetAll)).Methods(http.MethodGet)
	// Router.Handle("/category/update", http.HandlerFunc(categoryController.CategoryUpdate)).Methods(http.MethodPost)

	return Router
}
