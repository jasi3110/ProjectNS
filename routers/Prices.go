package routers

import (
	"OnlineShop/controllers"
	"OnlineShop/controllers/masters"
	"OnlineShop/utls"
	"net/http"

	"github.com/gorilla/mux"
)

func PricesRoutes(Router *mux.Router) *mux.Router {
	priceController := masters.PriceController{}

	Router.Handle("/prices/create", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(priceController.PriceCreate)))).Methods(http.MethodPost)
	// Router.Handle("/price/getbyid/{id}", http.HandlerFunc(priceController.PriceGetById)).Methods(http.MethodGet)
	// Router.Handle("/role/getall", http.HandlerFunc(roleController.RoleGetAll)).Methods(http.MethodGet)
	// Router.Handle("/role/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(roleController.RoleUpdate)))).Methods(http.MethodPost)

	return Router
}
