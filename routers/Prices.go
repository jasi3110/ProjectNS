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

	Router.Handle("/price/create", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(priceController.PriceCreate)))).Methods(http.MethodPost)
	Router.Handle("/price/getbyid/{id}", http.HandlerFunc(priceController.PriceGetById)).Methods(http.MethodGet)
	Router.Handle("/Price/getbydate", http.HandlerFunc(priceController.PriceGetByDate)).Methods(http.MethodPost)
	Router.Handle("/Price/priceproductgetall", http.HandlerFunc(priceController.PriceProductGetAll)).Methods(http.MethodPost)
	Router.Handle("/Price/getall", http.HandlerFunc(priceController.PriceGetAll)).Methods(http.MethodGet)
	Router.Handle("/Price/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(priceController.PriceUpdate)))).Methods(http.MethodPost)

	return Router
}
