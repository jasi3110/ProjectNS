package routers


import (
	// "OnlineShop/controllers"
	"OnlineShop/controllers"
	// "OnlineShop/controllers/masters"
	"OnlineShop/utls"

	// "OnlineShop/repos/masterRepo"

	// "OnlineShop/utls"
	"net/http"

	"github.com/gorilla/mux"
)

func SaleRoutes(Router *mux.Router) *mux.Router {
	saleController := controllers.SaleController{}

	Router.Handle("/sale/createsale", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(saleController.SaleEntry)))).Methods(http.MethodPost)
	// Router.Handle("/role/getbyid/{id}", http.HandlerFunc(roleController.RoleGetById)).Methods(http.MethodGet)
	// Router.Handle("/role/getall", http.HandlerFunc(roleController.RoleGetAll)).Methods(http.MethodGet)
	// Router.Handle("/role/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(roleController.RoleUpdate)))).Methods(http.MethodPost)

	return Router
}