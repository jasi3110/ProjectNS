package masters

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	masterrepo "OnlineShop/repos/masterRepo"
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
	}
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
	w.Write(resp)
}

func (discount *DiscountController) DiscountUpdate(w http.ResponseWriter, r *http.Request) {
	request:= models.Role{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in Decoding RoleUpdate Request :", err)
	}
	Repo := masterrepo.RoleInterface(&masterrepo.RoleStruct{})
	descreption, status := Repo.RoleUpdate(&request)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	respone, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("Error in Marshal RoleUpdate Response :", err)
	}
	w.Write(respone)
}

func (discount *DiscountController) DiscountGetById(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	roleid := strconv.FormatInt(id, 10)
	if err != nil {
		fmt.Println("Error in Decoding RoleGetById Request :", err)
	}
	roleStruct := models.Role{
		Id: roleid,
	}
	if err != nil {
		fmt.Println(err)
	}
	repo := masterrepo.RoleInterface(&masterrepo.RoleStruct{})
	value, status, descreption := repo.RoleById(&roleStruct)
	response := models.RoleResponses{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	respone, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal RoleGetById Response :", err)
	}
	w.Write(respone)
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
	w.Write(respone)
}


