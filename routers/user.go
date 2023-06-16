package routers

import (
	"OnlineShop/controllers"
	"net/http"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) *mux.Router {

	userController := controllers.UserController{}

	router.Handle("/user/create", http.HandlerFunc(userController.UserCreate)).Methods(http.MethodPost)
	router.Handle("/user/login", http.HandlerFunc(userController.UserLogin)).Methods(http.MethodPost)
	router.Handle("/user/updatepassword",http.HandlerFunc(userController.UserUpdatePassword)).Methods(http.MethodPost)
	router.Handle("/user/updatemobileno", http.HandlerFunc(userController.UserUpdateMobileno)).Methods(http.MethodPost)
	router.Handle("/user/updateemail", http.HandlerFunc(userController.UserUpdateEmail)).Methods(http.MethodPost)
	router.Handle("/user/changepassword",http.HandlerFunc(userController.UserChangePassword)).Methods(http.MethodPost)
	router.Handle("/user/verifymobileno", http.HandlerFunc(userController.UserVerifyMobileno)).Methods(http.MethodPost)
	router.Handle("/user/updatename", http.HandlerFunc(userController.UserUpdateName)).Methods(http.MethodPost)
	router.Handle("/user/checkingpassword", http.HandlerFunc(userController.UserCheckingPassword)).Methods(http.MethodPost)
	router.Handle("/user/userverifybyid", http.HandlerFunc(userController.UserVerifyById)).Methods(http.MethodPost)
	router.Handle("/user/getbyid/{id}", http.HandlerFunc(userController.UserGetById)).Methods(http.MethodGet)
	router.Handle("/user/getall", http.HandlerFunc(userController.UserGetAll)).Methods(http.MethodGet)
	router.Handle("/user/verifyuser", http.HandlerFunc(userController.Userverify)).Methods(http.MethodPost)
	router.Handle("/user/checkotp", http.HandlerFunc(userController.UserCheckOtp)).Methods(http.MethodPost)
	router.Handle("/user/delete", http.HandlerFunc(userController.UserDelete)).Methods(http.MethodPost)

	return router
}
