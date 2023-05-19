package masters

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type DiscountController struct {
}

func (discount *DiscountController) DiscountCreate(w http.ResponseWriter, r *http.Request) {
	requst := models.RDiscount{}
	err := json.NewDecoder(r.Body).Decode(&requst)
	if err != nil {
		fmt.Println("Error in Decoding RoleCreate Request :", err)
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
	repo := repos.DiscountInterface(&repos.DiscountStruct{})
	status, descreption := repo.CreateDiscount(&requst)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal RoleCreate Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
}

func (discount *DiscountController) DiscountUpdate(w http.ResponseWriter, r *http.Request) {
	request:= models.RDiscount{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in Decoding RoleUpdate Request :", err)
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
	Repo := repos.DiscountInterface(&repos.DiscountStruct{})
	status,descreption:= Repo.DiscountUpdate(&request)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	respone, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("Error in Marshal RoleUpdate Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(respone)
}
}

func (discount *DiscountController) DiscountGetById(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	
	if err != nil {
		fmt.Println("Error in Decoding Discount Product GetById Request :", err)
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
	
	if err != nil {
		fmt.Println(err)
	}
	Repo := repos.DiscountInterface(&repos.DiscountStruct{})
	value,status,descreption := Repo.DiscountProductById(&id)
	response := models.ProductResponses{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	respone, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal Discpunt Product GetById Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(respone)
}
}

func (discount *DiscountController) DiscountGetAll(w http.ResponseWriter, r *http.Request) {
	
	repo := repos.DiscountInterface(&repos.DiscountStruct{})
	
	value, status,descreption := repo.DiscountGetAll()
	response := models.GetAllDiscountResponse{
		Statuscode:  200,
		Status: status,
		Value:     value,
		Descreption: descreption,
	}
	respone, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal RoleGetAll Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(respone)
}