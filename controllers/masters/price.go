package masters

import (
	"OnlineShop/models"
	"OnlineShop/repos/masterRepo"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "strconv"
	// "github.com/gorilla/mux"
)

type PriceController struct {
}

func (Price *PriceController) PriceCreate(w http.ResponseWriter, r *http.Request) {
	requst := models.Price{}
	err := json.NewDecoder(r.Body).Decode(&requst)
	if err != nil {
		fmt.Println("Error in Decoding RoleCreate Request :", err)
	}
	repo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	status, descreption ,_:= repo.CreatePrice(&requst)
	response := models.CommanRespones{
		Statuscode:  http.StatusOK,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal RoleCreate Response :", err)
	}
	w.Write(resp)
}

func (price *PriceController) PriceUpdate(w http.ResponseWriter, r *http.Request) {
	request := models.Price{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in Decoding PriceUpdate Request :", err)
	}
	Repo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	descreption, status := Repo.PriceUpdate(&request)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	respone, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("Error in Marshal PriceUpdate Response :", err)
	}
	w.Write(respone)
}

func (price *PriceController) PriceGetById(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	
	if err != nil {
		fmt.Println("Error in Decoding PriceGetById Request :", err)
	}
	priceStruct := models.Price{
		Id: id,
	}
	if err != nil {
		fmt.Println(err)
	}
	repo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	value, status, descreption := repo.PriceById(&priceStruct)
	response := models.PriceResponses{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	respone, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal PriceGetById Response :", err)
	}
	w.Write(respone)
}

func (price *PriceController) PriceGetByDate(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	productid, err := strconv.ParseInt(request["productid"], 10, 64) 
	pricedate := strconv.FormatInt(productid, 10)
	
	if err != nil {
		fmt.Println("Error in Decoding PriceGetById Request :", err)
	}
	priceStruct := models.Price{
		Createdon: pricedate,
	}
	if err != nil {
		fmt.Println(err)
	}
	repo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	value, status, descreption := repo.PriceById(&priceStruct)
	response := models.PriceResponses{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	respone, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal PriceGetById Response :", err)
	}
	w.Write(respone)
}


// func (role *RoleController) RoleGetAll(w http.ResponseWriter, r *http.Request) {

// 	repo := masterrepo.RoleInterface(&masterrepo.RoleStruct{})
// 	value, status, descreption := repo.RoleGetAll()
// 	response := models.GetAllRoleResponse{
// 		Statuscode:  200,
// 		Status:   Prices status,
// 		Value:       value,
// 		Descreption: descreption,
// 	}
// 	respone, err := json.Marshal(response)
// 	if err != nil {
// 		fmt.Println("Error in Marshal RoleGetAll Response :", err)
// 	}
// 	w.Write(respone)
// }