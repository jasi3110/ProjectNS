package routers

import (
	"OnlineShop/controllers"
	"OnlineShop/controllers/masters"
	"OnlineShop/utls"
	"net/http"
	"github.com/gorilla/mux"
)

func AdImagesRoutes(Router *mux.Router) *mux.Router {
	ImageController := masters.AdImageController{}

	Router.Handle("/adimages/create", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(ImageController.AdImageCreate)))).Methods(http.MethodPost)
	Router.Handle("/adimages/getall", http.HandlerFunc(ImageController.AdImageGetall)).Methods(http.MethodGet)
	// Router.Handle("/role/getall", http.HandlerFunc(roleController.RoleGetAll)).Methods(http.MethodGet)
	// Router.Handle("/role/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(roleController.RoleUpdate)))).Methods(http.MethodPost)

	return Router
}
