package dashboard

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DashboardCart struct {
}

func (dashboard *DashboardCart) Cartpage(w http.ResponseWriter, r *http.Request) {
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
	result := []models.DashboardCart1{}
	nums := []models.ProductAll{}

	val := models.DashboardCart1{
		Id:       1,
		ViewType: 1,
		Data:     nums,
	}
	result = append(result, val)
	repo := repos.CartInterface(&repos.CartStruct{})
	value, status, descreption := repo.CartGetAll(&id)
	
	
	val12 := models.DashboardCart1{
		Id:       2,
		Type:     "CartData",
		ViewType: 2,
		// Data:     value.Value,
	}
	result = append(result, val12)
	val13 := models.DashboardCart1{
		Id:       3,
		ViewType: 3,
		Data:     nums,
		Items: value.Items,
	}
	result = append(result, val13)
	val14 := models.DashboardCart1{
		Id:       4,
		ViewType: 4,
		Data:     nums,
	}
	result = append(result, val14)
	
	response := models.CommanResponesDashboardCart{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
		Value: result,
	}

	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal Dashboard in Homepage Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}}