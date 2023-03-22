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
	Router.Handle("/sale/invoicegetallbycustomerid/{id}", http.HandlerFunc(saleController.InvoiceGetallByCustomerid)).Methods(http.MethodGet)
	
	Router.Handle("/sale/getbybillid/{id}", http.HandlerFunc(saleController.SaleGetByBillId)).Methods(http.MethodGet)
	Router.Handle("/sale/getbycustomerid/{id}", http.HandlerFunc(saleController.SaleGetByCustomerid)).Methods(http.MethodGet)
	Router.Handle("/saleInvoice/getall", http.HandlerFunc(saleController.SaleInvoiceGetAll)).Methods(http.MethodGet)
	Router.Handle("/sale/ GetAllSaleByDateRange", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(saleController.GetUserReportByDateRange)))).Methods(http.MethodPost)
	Router.Handle("/sale/delete", utls.Authorize(controllers.CheckAuthenticLogin(http.HandlerFunc(saleController.SaleDelete)))).Methods(http.MethodPost)

	return Router
}