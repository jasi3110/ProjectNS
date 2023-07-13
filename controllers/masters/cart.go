package masters

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CartController struct {
}

func (cart *CartController) CartCreate(w http.ResponseWriter, r *http.Request) {
	requst := models.RCart{}
	err := json.NewDecoder(r.Body).Decode(&requst)
	if err != nil {
		fmt.Println("Error in Decoding CartCreate Request :", err)
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{
	repo := repos.CartInterface(&repos.CartStruct{})
	status, descreption := repo.Createcart(&requst)
	response := models.CartRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal CartCreate Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
}

func (cart *CartController) CartUpdate(w http.ResponseWriter, r *http.Request) {
	request := models.GetCart{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in Decoding CartUpdate Request :", err)
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{
	
	repo := repos.CartInterface(&repos.CartStruct{})
	descreption, status := repo.CartUpdate(&request)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("Error in Marshal CartUpdate Response:", err)
	}
	w.Write(resp)
}}

func (cart *CartController) CartGetAll(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	if err != nil {
		fmt.Println("Error in Decoding  CartAll Request :", err)
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{
	repo := repos.CartInterface(&repos.CartStruct{})
	value, status, descreption := repo.CartGetAll(&id)
	response := models.GetAllCartResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value.Value,
		Items: value.Items,
		// Productprice: value.Productprice,
		// Productdiscoiunt: value.Productdiscoiunt,
		// Total: value.Total,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal CartGetAll Response :",err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}}

func (cart *CartController) CartProductDelete(w http.ResponseWriter, r *http.Request) {
	request := models.RCart{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in Decoding Cart Product Delete Request :", err)
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{
	repo := repos.CartInterface(&repos.CartStruct{})

	 status, descreption := repo.CartProductDelete(&request)
	
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshal Cart Product Delete Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}}

func (cart *CartController) CartDelete(w http.ResponseWriter, r *http.Request) { 
	request := models.RCart{}
	err := json.NewDecoder(r.Body).Decode(&request)
	
	if err != nil {
		log.Println("Error in Decoding Cart Delete Request :", err)
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal CartUpdate Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{
	
	repo := repos.CartInterface(&repos.CartStruct{})

	 status, descreption := repo.CartDelete(&request)
	
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)

	if err != nil {
		log.Println("Error in Marshal Cart Delete Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}}
