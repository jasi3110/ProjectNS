package masters

import (
	"OnlineShop/models"
	masterrepo "OnlineShop/repos/masterRepo"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CategoryController struct {
}

func (category *CategoryController) CategoryCreate(w http.ResponseWriter, r *http.Request) {
	requst := models.Category{}
	err := json.NewDecoder(r.Body).Decode(&requst)
	if err != nil {
		fmt.Println("Error in Decoding CategoryCreate Request :", err)
	}
	repo := masterrepo.CategoryInterface(&masterrepo.CategoryStruct{})
	status, descreption := repo.CreateCategory(&requst)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal CategoryCreate Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (category *CategoryController) CategoryUpdate(w http.ResponseWriter, r *http.Request) {
	request := models.Category{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in Decoding CategoryUpdate Request :", err)
	}
	repo := masterrepo.CategoryInterface(&masterrepo.CategoryStruct{})
	value, descreption, status := repo.CategoryUpdate(&request)
	response := models.CategoryResponses{
		Statuscode:  200,
		Status:      status,
		Value:     value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("Error in Marshal CategoryUpdate Response:", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (category *CategoryController) CategoryGetById(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	if err != nil {
		fmt.Println("Error in Decoding  CategoryGetById Request :", err)
	}

	CategoryStruct := models.Category{
		Id:id,
	}
	if err != nil {
		fmt.Println(err)
	}
	repo := masterrepo.CategoryInterface(&masterrepo.CategoryStruct{})
	value, status, descreption := repo.CategoryById(&CategoryStruct)
	response := models.CategoryResponses{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal CategoryGetById Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (category *CategoryController) CategoryGetAll(w http.ResponseWriter, r *http.Request) {
	
	repo := masterrepo.CategoryInterface(&masterrepo.CategoryStruct{})
	value, status, descreption := repo.CategoryGetAll()
	response := models.GetAllCategoryResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal CategoryGetAll Response :",err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}


