package masters

import (
	"OnlineShop/models"
	"OnlineShop/repos/masterRepo"
	"encoding/json"
	"fmt"
	"net/http"
	// "strconv"

	// "github.com/gorilla/mux"
)

type PriceController struct {
}

func (Prices *PriceController) PriceCreate(w http.ResponseWriter, r *http.Request) {
	requst := models.Price{}
	err := json.NewDecoder(r.Body).Decode(&requst)
	if err != nil {
		fmt.Println("Error in Decoding RoleCreate Request :", err)
	}
	repo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	status, descreption ,_:= repo.CreatePrices(&requst)
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

// func (role *RoleController) RoleUpdate(w http.ResponseWriter, r *http.Request) {
// 	request := models.Role{}
// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	if err != nil {
// 		fmt.Println("Error in Decoding RoleUpdate Request :", err)
// 	}
// 	Repo := masterrepo.RoleInterface(&masterrepo.RoleStruct{})
// 	descreption, status := Repo.RoleUpdate(&request)
// 	response := models.CommanRespones{
// 		Statuscode:  200,
// 		Status:      status,
// 		Descreption: descreption,
// 	}
// 	respone, err := json.Marshal(&response)
// 	if err != nil {
// 		fmt.Println("Error in Marshal RoleUpdate Response :", err)
// 	}
// 	w.Write(respone)
// }

// func (price *PriceController) PriceGetById(w http.ResponseWriter, r *http.Request) {
// 	request := mux.Vars(r)
// 	id, err := strconv.ParseInt(request["id"], 10, 64)
// 	priceid := strconv.FormatInt(id, 10)
// 	if err != nil {
// 		fmt.Println("Error in Decoding PriceGetById Request :", err)
// 	}
// 	priceStruct := models.Role{
// 		Id: priceid,
// 	}
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// repo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
// 	// value, status, descreption := repo.PriceById(&priceStruct)
// 	response := models.RoleResponses{
// 		Statuscode:  200,
// 		Status:      status,
// 		Value:       value,
// 		Descreption: descreption,
// 	}
// 	respone, err := json.Marshal(response)
// 	if err != nil {
// 		fmt.Println("Error in Marshal PriceGetById Response :", err)
// 	}
// 	w.Write(respone)
// }

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