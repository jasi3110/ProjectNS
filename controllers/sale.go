package controllers

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	// "OnlineShop/utls"
	"encoding/json"
	"fmt"
	"net/http"
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
status,descreption:=Repo.CreateSale(request)
response:=models.CommanRespones{
	Statuscode:  http.StatusAccepted,
	Status:      status,
	Descreption:descreption,
}

resp,err:=json.Marshal(response)
if err!=nil {
	fmt.Println("Error in Marshal saleEntry :",err)
}
w.Write(resp)
}

