package routers


import (
	"OnlineShop/controllers"
	"OnlineShop/utls"
	"net/http"
	"github.com/gorilla/mux"
)

func SaleRoutes(Router *mux.Router) *mux.Router {

	saleController := controllers.SaleController{}

	Router.Handle("/sale/createsale", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(saleController.SaleEntry)))).Methods(http.MethodPost)
	Router.Handle("/sale/invoicegetallbyUserid/{id}", http.HandlerFunc(saleController.InvoiceGetallByUserid)).Methods(http.MethodGet)
	Router.Handle("/sale/getbyinvoiceid/{id}", http.HandlerFunc(saleController.GetSaleByInvoiceId)).Methods(http.MethodGet)
	Router.Handle("/saleInvoice/getall", http.HandlerFunc(saleController.SaleInvoiceGetAll)).Methods(http.MethodGet)
	Router.Handle("/sale/invoiceByDateRange", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(saleController.InvoiceByDateRange)))).Methods(http.MethodPost)
	Router.Handle("/sale/delete", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(saleController.SaleDelete)))).Methods(http.MethodPost)

	return Router
}