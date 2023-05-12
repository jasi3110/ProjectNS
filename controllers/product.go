package controllers

import (
	"OnlineShop/models"
	"OnlineShop/repos"

	// "OnlineShop/repos/masterRepo"

	// "OnlineShop/repos/masterRepo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
}

func (Product *Product) ProductCreate(w http.ResponseWriter, r *http.Request) {
	request := models.Product{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in Decoding ProductCreate Request :", err)
	}
	repo := repos.ProductInterface(&repos.ProductStruct{})
	descreption, status := repo.ProductCreate(&request)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}

	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal ProductCreate Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (Product *Product) ProductUpdate(w http.ResponseWriter, r *http.Request) {
	request := models.Product{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error in Decoding ProductUpdate Request :", err)
	}
	repo := repos.ProductInterface(&repos.ProductStruct{})
	result, status := repo.ProductUpdate(&request)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: result,
	}
	respone, err := json.Marshal(&response)
	if err != nil {
		log.Println("Error in Marshal ProductUpadate Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(respone)
}

func (Product *Product) ProductDelete(w http.ResponseWriter, r *http.Request) {
	request := models.Product{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error in Decoding ProductDelete Request :", err)
	}
	repo := repos.ProductInterface(&repos.ProductStruct{})
	 status,result := repo.ProductDelete(&request)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: result,
	}
	respone, err := json.Marshal(&response)
	if err != nil {
		log.Println("Error in Marshal ProductDelete Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(respone)
}

func (Product *Product) ProductGetById(w http.ResponseWriter, r *http.Request) {

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)

	if err != nil {
		fmt.Println("Error in Decoding ProductGetById Request :", err)
	}
	repo := repos.ProductInterface(&repos.ProductStruct{})

	value, status, descreption := repo.GetProductById(&id)
	
	response := models.ProductResponses{
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

func (product *Product) ProductGetAll(w http.ResponseWriter, r *http.Request) {

	Repo := repos.ProductInterface(&repos.ProductStruct{})
	value, status, descreption := Repo.ProductGetAll()
	response := models.GetAllProductResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		log.Println("Error in Marshal ProductGetAll Request :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (Product *Product) ProductGetAllByCategory(w http.ResponseWriter, r *http.Request) {

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)

	if err != nil {
		fmt.Println("Error in Decoding ProductGetALLByCategory  Request :", err)
	}
	repo := repos.ProductInterface(&repos.ProductStruct{})

	value, status, descreption := repo.ProductGetAllByCategory(&id)
	
	response := models.GetAllProductResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal ProductGetByCategory response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (Product *Product) ProductGetAllByUnit(w http.ResponseWriter, r *http.Request) {

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)

	if err != nil {
		fmt.Println("Error in Decoding ProductGetALLByUnit  Request :", err)
	}
	repo := repos.ProductInterface(&repos.ProductStruct{})

	value, status, descreption := repo.ProductGetAllByUnit(&id)
	
	response := models.GetAllProductResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal ProductGetByUnit response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (Product *Product) ProductSearchBar(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	values := url.Values{}
    for k, v := range request {
        values.Set(k, v)
    }
    searchBar := values.Encode()
	repo := repos.ProductInterface(&repos.ProductStruct{})
	value, status:= repo.ProductSearchBar(searchBar[5:])
	
	response := models.ProductSearchResponses{
		Statuscode:  200,
		Status:      status,
		Value:       value,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal ProductSearchBar response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}