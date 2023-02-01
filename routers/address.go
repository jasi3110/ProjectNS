package routers

import (
	// "OnlineShop/controllers"
	"OnlineShop/controllers/masters"
	// "OnlineShop/repos/masterRepo"

	// "OnlineShop/utls"
	"net/http"

	"github.com/gorilla/mux"
)

func UserAddressRoutes(Router *mux.Router) *mux.Router {

	UserAddressController := masters.UserAddressController{}

	Router.Handle("/useraddress/create", http.HandlerFunc(UserAddressController.UserAddressCreate)).Methods(http.MethodPost)
	Router.Handle("/useraddress/getbyid/{id}", http.HandlerFunc(UserAddressController.UserAddressGetById)).Methods(http.MethodGet)
	Router.Handle("/useraddress/getall", http.HandlerFunc(UserAddressController.UserAddressGetAll)).Methods(http.MethodGet)
	Router.Handle("/useraddress/update", http.HandlerFunc(UserAddressController.UserAddressUpdate)).Methods(http.MethodPost)
	Router.Handle("/useraddress/getallcustomer/{customerid}", http.HandlerFunc(UserAddressController.UserAddressGetAllCustomer)).Methods(http.MethodGet)

	return Router
}
