package routers

import (
	// "OnlineShop/controllers"
	"OnlineShop/controllers/masters"
	// "OnlineShop/repos/masterRepo"

	// "OnlineShop/utls"
	"net/http"

	"github.com/gorilla/mux"
)

func CategoryRoutes(Router *mux.Router) *mux.Router {

	categoryController := masters.CategoryController{}

	Router.Handle("/category/create", http.HandlerFunc(categoryController.CategoryCreate)).Methods(http.MethodPost)
	Router.Handle("/category/getbyid/{id}", http.HandlerFunc(categoryController.CategoryGetById)).Methods(http.MethodGet)
	Router.Handle("/category/getall", http.HandlerFunc(categoryController.CategoryGetAll)).Methods(http.MethodGet)
	Router.Handle("/category/update", http.HandlerFunc(categoryController.CategoryUpdate)).Methods(http.MethodPost)

	return Router
}
