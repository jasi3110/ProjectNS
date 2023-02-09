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

func DiscountRoutes(Router *mux.Router) *mux.Router {
	DiscountController := masters.DiscountController{}

	Router.Handle("/discount/create", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(DiscountController.DiscountCreate)))).Methods(http.MethodPost)
	// Router.Handle("/role/getbyid/{id}", http.HandlerFunc(roleController.RoleGetById)).Methods(http.MethodGet)
	Router.Handle("/discount/getall", http.HandlerFunc(DiscountController.DiscountGetAll)).Methods(http.MethodGet)
	// Router.Handle("/role/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(roleController.RoleUpdate)))).Methods(http.MethodPost)

	return Router
}
