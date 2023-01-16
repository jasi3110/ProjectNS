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

type UnitController struct {
}

func (unit *UnitController) UnitCreate(w http.ResponseWriter, r *http.Request) {
	requst := models.Unit{}
	err := json.NewDecoder(r.Body).Decode(&requst)
	if err != nil {
		fmt.Println("Error in Decoding Create Unit Request :", err)
	}
	repo := masterrepo.UnitInterface(&masterrepo.UnitStruct{})
	status, descreption := repo.CreateUnit(&requst)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	respone, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal Create Unit Response :", err)
	}
	w.Write(respone)
}

func (unit *UnitController) UnitGetById(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	unitid := strconv.FormatInt(id, 10)
	if err != nil {
		fmt.Println("Error in Decoding Unit GetId Request :", err)
	}

	UnitStruct := models.Unit{
		Id: unitid,
	}
	if err != nil {
		fmt.Println(err)
	}
	repo := masterrepo.UnitInterface(&masterrepo.UnitStruct{})
	value, status, descreption := repo.UnityById(&UnitStruct)
	response := models.UnitResponses{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	respone, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal unit GetId Response :", err)
	}
	w.Write(respone)
}

func (unit *UnitController) UnitGetAll(w http.ResponseWriter, r *http.Request) {
	
	repo := masterrepo.UnitInterface(&masterrepo.UnitStruct{})
	value, status, descreption := repo.UnitByAll()
	response := models.GetAllUnitResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	respone, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal Unit GetAll Response :", err)
	}
	w.Write(respone)
}

func (unit *UnitController) UnitUpdate(w http.ResponseWriter, r *http.Request) {
	request := models.Unit{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in Decoding Unit Update Request :", err)
	}
	repo := masterrepo.UnitInterface(&masterrepo.UnitStruct{})
	descreption, status := repo.UnitUpdate(&request)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	respone, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("Error in Marshal Unit Update Response :", err)
	}
	w.Write(respone)
}
