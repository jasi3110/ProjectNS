package controllers

import (
	"OnlineShop/models"
	"OnlineShop/repos"

	// "OnlineShop/repos/masterRepo"

	// "OnlineShop/repos/masterRepo"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"github.com/gorilla/mux"
)

type Product struct {
}

func (Product *Product) ProductCreate(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.Product{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding Product Create Request :", err)

		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = false
		response.Descreption = "Something Went Wrong"

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal  Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.ProductInterface(&repos.ProductStruct{})
		descreption, status := repo.ProductCreate(&request)

		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = status
		response.Descreption = descreption

		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal Product Create Response :", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	}
}

func (Product *Product) ProductUpdate(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.Product{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding Product Update Request :", err)

		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = false
		response.Descreption = "Something Went Wrong"

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal Product Update Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.ProductInterface(&repos.ProductStruct{})
		descreption, status := repo.ProductUpdate(&request)

		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = status
		response.Descreption = descreption

		respone, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal Product Update Response : ", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(respone)
	}
}

func (Product *Product) ProductDelete(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := models.Product{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Panic("Error in Decoding Product Delete Request :", err)

		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = false
		response.Descreption = "Something Went Wrong"

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal Product Delete Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.ProductInterface(&repos.ProductStruct{})
		status, descreption := repo.ProductDelete(&request)

		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = status
		response.Descreption = descreption

		respone, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal Product Delete Response :", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(respone)
	}
}

func (Product *Product) ProductGetById(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)

	if err != nil {
		log.Panic("Error in Decoding ProductGetById Request :", err)

		response := models.CommanRespones{}
		response.Statuscode = 200
		response.Status = false
		response.Descreption = "Something Went Wrong"

		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal ProductGetById Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.ProductInterface(&repos.ProductStruct{})
		value, status, descreption := repo.GetProductById(&id)

		response := models.ProductResponses{}
		response.Statuscode = 200
		response.Status = status
		response.Value = value
		response.Descreption = descreption

		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal ProductGetById response :", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (product *Product) ProductGetAll(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	Repo := repos.ProductInterface(&repos.ProductStruct{})
	value, status, descreption := Repo.ProductGetAll()

	response := models.GetAllProductResponse{}
	response.Statuscode = 200
	response.Status = status
	response.Value = value
	response.Descreption = descreption

	resp, err := json.Marshal(response)
	if err != nil {
		log.Panic("Error in Marshal ProductGetAll Request :", err)
	}

	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (Product *Product) ProductGetAllByCategory(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	if err != nil {
		log.Panic("Error in Decoding ProductGetALLByCategory  Request :", err)
		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Something Went Wrong"
			
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal ProductGetALLByCategory Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.ProductInterface(&repos.ProductStruct{})
		value, status, descreption := repo.ProductGetAllByCategory(&id)

		response := models.GetAllProductResponse{}
		response.Statuscode=200
		response.Status=status
		response.Value=value
		response.Descreption=descreption
		
		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal ProductGetByCategory response :", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	}
}

func (Product *Product) ProductGetAllByUnit(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)

	if err != nil {
		log.Panic("Error in Decoding ProductGetALLByUnit  Request :", err)
		response := models.CommanRespones{}
		response.Statuscode=200
		response.Status=false
		response.Descreption="Something Went Wrong"
		
		resp, err := json.Marshal(&response)
		if err != nil {
			log.Panic("Error in Marshal ProductGetAllByUnit Response:", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)

	} else {

		repo := repos.ProductInterface(&repos.ProductStruct{})
		value, status, descreption := repo.ProductGetAllByUnit(&id)

		response := models.GetAllProductResponse{}
		response.Statuscode=200
		response.Status=status
		response.Value=value
		response.Descreption=descreption
		
		resp, err := json.Marshal(response)
		if err != nil {
			log.Panic("Error in Marshal ProductGetByUnit response :", err)
		}

		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}
}

func (Product *Product) ProductSearchBar(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	request := mux.Vars(r)
	values := url.Values{}
	for k, v := range request {
		values.Set(k, v)
	}
	searchBar := values.Encode()
	result:=searchBar[5:]
	repo := repos.ProductInterface(&repos.ProductStruct{})
	value, status := repo.ProductSearchBar(result)

	response := models.ProductSearchResponses{}
	response.Statuscode=200
	response.Status=status
	response.Value=value
	
	resp, err := json.Marshal(response)
	if err != nil {
		log.Panic("Error in Marshal ProductSearchBar response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
