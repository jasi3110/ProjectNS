package controllers

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type SaleController struct {
}

func (sale *SaleController) SaleEntry(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.Invoice{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding Sale Entry Requst :", err)
		
		
		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Something Went Wrong"
			
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal Sale Entry Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}

	Repo := repos.SaleInterface(&repos.SaleStruct{})
	status, descreption, value := Repo.CreateSale(&request)

	response := models.SaleCommanRespones{}

	response.Statuscode = 200
	response.Status = status
	response.Value = value
	response.Descreption = descreption

	resp, err := json.Marshal(response)
	if err != nil {
		log.Panic("Error in Marshal sale Entry :", err)
	}

	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (sale *SaleController) SaleInvoiceGetAll(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	Repo := repos.SaleInterface(&repos.SaleStruct{})
	value, status, descreption := Repo.InvoiceGetall()
	response := models.GetAllInvoiceResponse{}

	response.Statuscode=200
	response.Status=status
	response.Value=value
	response.Descreption=descreption
	
	resp, err := json.Marshal(response)
	if err != nil {
		log.Panic("Error in Marshal SaleInvoiceGetAll Request :", err)
	}

	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (sale *SaleController) InvoiceGetallByUserid(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	if err != nil {
		log.Panic("Error in Decoding Invoice Getall By Userid Request :", err)

		
		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Something Went Wrong"
			
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal  Invoice Getall By Userid Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}

	Repo := repos.SaleInterface(&repos.SaleStruct{})
	value, status,descreption := Repo.InvoiceGetallByUserid(&id)
	response := models.GetAllInvoiceResponse{}

	response.Statuscode=200
	response.Status=status
	response.Value=value
	response.Descreption=descreption
		
	resp, err := json.Marshal(response)
	if err != nil {
		log.Panic("Error in Marshal Invoice Getall By Userid :", err)
	}

	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (sale *SaleController) GetSaleByInvoiceId(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	if err != nil {
		log.Panic("Error in Decoding SaleGetByInvoiceId Request :", err)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Something Went Wrong"
		
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal SaleGetByInvoiceId Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.SaleInterface(&repos.SaleStruct{})
		value, status, descreption := repo.GetSaleByInvoiceid(&id)

		response := models.GetSaleByInvoiceIdResponse{}
		response.Statuscode=200
		response.Status=status
		response.Value=value
		response.Descreption=descreption

		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal SaleGetByInvoiceId response :", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (sale *SaleController) InvoiceByDateRange(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.InvoiceByDateRange{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding  InvoiceByDateRange Requst :", err)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Something Went Wrong"
		
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal InvoiceByDateRange Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		if err != nil {
			log.Panic("Error in Decoding InvoiceByDateRange Request :", err)
		}

		repo := repos.SaleInterface(&repos.SaleStruct{})
		value, status, descreption := repo.InvoiceByDateRange(&request)

		response := models.GetAllSaleInvoiceByDateRangeResponse{}
		response.Statuscode=200
		response.Status=status
		response.Value=value
		response.Descreption=descreption

		
		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal ProductGetById response :", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (sale *SaleController) SaleDelete(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	billid := models.Invoice{
		Id: id,
	}

	if err != nil {
		log.Panic("Error in Decoding Sale Delete Request :", err)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Something Went Wrong"
		
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal CartUpdate Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.SaleInterface(&repos.SaleStruct{})
		status, descreption := repo.SaleDelete(&billid)

		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=status
		response.Descreption=descreption

		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal Sale Delete Response : ", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}
