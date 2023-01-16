package controllers

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	"OnlineShop/repos/masterRepo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	w.Write(resp)
}

func (Product *Product) ProductUpdate(w http.ResponseWriter, r *http.Request) {
	request := models.Product{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error in Decoding ProductUpdate Request :", err)
		response := models.CommanRespones{
			Statuscode:  300,
			Status:      false,
			Descreption: "Error in Decoding ProductUpdate Request :",
		}
		respone, err := json.Marshal(&response)
		if err != nil {
			log.Println("Error in Marshal ProductUpadate Response :", err)
		}
		w.Write(respone)
	}
	repo := repos.ProductInterface(&repos.ProductStruct{})
	result, status := repo.ProductUpdate(&request)
	response := models.CommanRespones{
		Statuscode:  300,
		Status:      status,
		Descreption: result,
	}
	respone, err := json.Marshal(&response)
	if err != nil {
		log.Println("Error in Marshal ProductUpadate Response :", err)
	}
	w.Write(respone)
}

func (Product *Product) ProductGetById(w http.ResponseWriter, r *http.Request) {

	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)

	if err != nil {
		fmt.Println("Error in Decoding ProductGetById Request :", err)
	}
	repo := repos.ProductInterface(&repos.ProductStruct{})

	Value, status, descreption := repo.GetProductById(&id)
	categoryStruct := models.Category{
		Id: Value.Category,
	}
	UnitStruct := models.Unit{
		Id: Value.Unit,
	}
	// PriceStruct:= models.Price{
	// 	Id: Value.Price,
	// }
	categoryrepo := masterRepo.CategoryInterface(&masterRepo.CategoryStruct{})
	unitrepo := masterRepo.UnitInterface(&masterRepo.UnitStruct{})

	category, _, _ := categoryrepo.CategoryById(&categoryStruct)
	unit, _, _ := unitrepo.UnityById(&UnitStruct)
	value:= models.ProductAll{
		Id:        Value.Id,
		Name:      Value.Name,
		Category:  category,
		Quantity:  Value.Quantity,
		Unit:      unit,
		Price:     Value.Price,
		CreatedOn: Value.CreatedOn,
	}
	response:= models.ProductResponses{
		Statuscode:  200,
		Status:      status,
		Value:     value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal ProductGetById response :", err)
	}
	w.Write(resp)
}

func (product *Product) ProductGetAll(w http.ResponseWriter, r *http.Request) {
	
	Repo := repos.ProductInterface(&repos.ProductStruct{})
	value, status,descreption := Repo.ProductGetAll()
	response := models.GetAllProductResponse{
		Statuscode: 200,
		Status: status,
		Value: value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		log.Println("Error in Marshal ProductGetAll Request :", err)
	}
	w.Write(resp)
}
