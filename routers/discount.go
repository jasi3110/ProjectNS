package routers

import (
	"OnlineShop/controllers"
	"OnlineShop/controllers/masters"
	"OnlineShop/utls"
	"net/http"
	"github.com/gorilla/mux"
)

func DiscountRoutes(Router *mux.Router) *mux.Router {
	DiscountController := masters.DiscountController{}

	Router.Handle("/discount/create", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(DiscountController.DiscountCreate)))).Methods(http.MethodPost)
	Router.Handle("/discount/getbyid/{id}", http.HandlerFunc(DiscountController.DiscountGetById)).Methods(http.MethodGet)
	Router.Handle("/discount/getall", http.HandlerFunc(DiscountController.DiscountGetAll)).Methods(http.MethodGet)
	Router.Handle("/discount/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(DiscountController.DiscountUpdate)))).Methods(http.MethodPost)

	return Router
}
