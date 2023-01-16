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

func RoleRoutes(Router *mux.Router) *mux.Router {
	roleController := masters.RoleController{}

	Router.Handle("/role/create", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(roleController.RoleCreate)))).Methods(http.MethodPost)
	Router.Handle("/role/getbyid/{id}", http.HandlerFunc(roleController.RoleGetById)).Methods(http.MethodGet)
	Router.Handle("/role/getall", http.HandlerFunc(roleController.RoleGetAll)).Methods(http.MethodGet)
	Router.Handle("/role/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(roleController.RoleUpdate)))).Methods(http.MethodPost)

	return Router
}
