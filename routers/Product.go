package routers

import (
	"OnlineShop/controllers"
	"OnlineShop/utls"
	"net/http"

	"github.com/gorilla/mux"
)

func ProductRoutes(router *mux.Router) *mux.Router {

	productController := controllers.Product{}

	router.Handle("/product/create", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(productController.ProductCreate)))).Methods(http.MethodPost)
	router.Handle("/product/getbyid/{id}", http.HandlerFunc(productController.ProductGetById)).Methods(http.MethodGet)
	router.Handle("/product/update", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(productController.ProductUpdate)))).Methods(http.MethodPost)
	router.Handle("/product/delete", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(productController.ProductDelete)))).Methods(http.MethodPost)
	
	
	router.Handle("/product/getall", http.HandlerFunc(productController.ProductGetAll)).Methods(http.MethodGet)
	router.Handle("/product/getallbycategory/{id}", http.HandlerFunc(productController.ProductGetAllByCategory)).Methods(http.MethodGet)
	router.Handle("/product/getallbyunit/{id}", http.HandlerFunc(productController.ProductGetAllByUnit)).Methods(http.MethodGet)

	return router
}
