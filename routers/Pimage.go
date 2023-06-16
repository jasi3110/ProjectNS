package routers

import (
	"OnlineShop/controllers"
	"OnlineShop/controllers/masters"
	"OnlineShop/utls"
	"net/http"
	"github.com/gorilla/mux"
)

func ProductImageRoutes(Router *mux.Router) *mux.Router {
	ImageController := masters.ImageController{}

	Router.Handle("/productimage/create", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(ImageController.ImageCreate)))).Methods(http.MethodPost)
	Router.Handle("/productimage/getbyid/{id}", http.HandlerFunc(ImageController.ProdutImageById)).Methods(http.MethodGet)
	// Router.Handle("/role/getall", http.HandlerFunc(roleController.RoleGetAll)).Methods(http.MethodGet)
	// Router.Handle("/role/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(roleController.RoleUpdate)))).Methods(http.MethodPost)

	return Router
}
