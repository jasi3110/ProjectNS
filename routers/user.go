package routers

import (
	"OnlineShop/controllers"
	"OnlineShop/utls"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) *mux.Router {

	userController := controllers.UserController{}

	router.Handle("/user/create", http.HandlerFunc(userController.UserCreate)).Methods(http.MethodPost)
	router.Handle("/user/login", http.HandlerFunc(userController.UserLogin)).Methods(http.MethodPost)
	router.Handle("/user/updatepassword", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(userController.UserUpdatePassword)))).Methods(http.MethodPost)
	router.Handle("/user/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(userController.UserUpdate)))).Methods(http.MethodPost)
	router.Handle("/user/getbyid/{id}", http.HandlerFunc(userController.UserGetById)).Methods(http.MethodGet)
	router.Handle("/user/getall", http.HandlerFunc(userController.UserGetAll)).Methods(http.MethodGet)

	return router
}
