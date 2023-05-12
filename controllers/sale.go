package controllers

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	"strconv"

	// "OnlineShop/utls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type SaleController struct {
}

func (sale *SaleController) SaleEntry(w http.ResponseWriter, r *http.Request) {
	request:=models.Invoice{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err!=nil {
		fmt.Println("Error in Decoding SaleEntry Requst :",err)
	}

Repo:=repos.SaleInterface(&repos.SaleStruct{})
status,descreption,value:=Repo.CreateSale(&request)
response:=models.SaleCommanRespones{
	Statuscode:  http.StatusAccepted,
	Status:      status,
	Value: value,
	Descreption:descreption,
}

resp,err:=json.Marshal(response)
if err!=nil {
	fmt.Println("Error in Marshal saleEntry :",err)
}
w.Header().Set("Content-Type", "Application/json")
w.Write(resp)
}

func (sale *SaleController) SaleInvoiceGetAll(w http.ResponseWriter, r *http.Request) {

	Repo := repos.SaleInterface(&repos.SaleStruct{})
	value, status:= Repo.InvoiceGetall()
	response := models.GetAllSaleInvoiceResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption:"",
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal SaleGetAll Request :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (sale *SaleController) InvoiceGetallByCustomerid(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	if err != nil {
		fmt.Println("Error in Decoding SaleGetAll Request :", err)
	}
	Repo := repos.SaleInterface(&repos.SaleStruct{})
	value, status:= Repo.InvoiceGetallByCustomerid(&id)
	response := models.GetAllSaleInvoiceResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption:"ssvs",
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal SaleGetAll Request :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (sale *SaleController) SaleGetByBillId(w http.ResponseWriter, r *http.Request) {

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)

	if err != nil {
		fmt.Println("Error in Decoding SaleGetByBillId Request :", err)
	}
	repo := repos.SaleInterface(&repos.SaleStruct{})

	value, status, descreption := repo.SaleGetByBillid(&id)
	
	response := models. GetAllSaleInvoiceGetByBillIdResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal ProductGetById response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (sale *SaleController) SaleGetByCustomerid(w http.ResponseWriter, r *http.Request) {

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)

	if err != nil {
		fmt.Println("Error in Decoding SaleGetByBillId Request :", err)
	}
	repo := repos.SaleInterface(&repos.SaleStruct{})

	value, status, descreption := repo.SaleGetByCustomerid(&id)
	
	response := models. GetAllSaleInvoiceGetByBillIdResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal ProductGetById response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (sale *SaleController) GetUserReportByDateRange(w http.ResponseWriter, r *http.Request) {

	request := models.GetUserReportByDateRange{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err!=nil {
		fmt.Println("Error in Decoding  GetUserReportByDateRange Requst :",err)
	}
	

	if err != nil {
		fmt.Println("Error in Decoding GetUserReportByDateRange Request :", err)
	}
	repo := repos.SaleInterface(&repos.SaleStruct{})

	value, status, descreption := repo.GetUserReportByDateRange(&request)
	
	response := models.  GetAllSaleInvoiceByDateRangeResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal ProductGetById response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (sale *SaleController) SaleDelete(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	billid := models.Invoice{
		Id: id,
	}
	if err != nil {
		fmt.Println("Error in Decoding UserGetById Request :", err)
	}
	repo := repos.SaleInterface(&repos.SaleStruct{})

	 status, descreption := repo.SaleDelete(&billid)
	
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)

	if err != nil {
		fmt.Println("Error in Marshal UserGetById Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}