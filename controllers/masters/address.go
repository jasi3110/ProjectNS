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

type UserAddressController struct {
}

func (UserAddress *UserAddressController) UserAddressCreate(w http.ResponseWriter, r *http.Request) {
	requst := models.UserAddress{}
	err := json.NewDecoder(r.Body).Decode(&requst)
	if err != nil {
		fmt.Println("Error in Decoding UserAddressCreate Request :", err)
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
	
	repo := masterrepo.UserAddressInterface(&masterrepo.UserAddressStruct{})
	status, descreption := repo.UserAddressCreate(&requst)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal UserAddressCreate Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}}

func (UserAddress *UserAddressController) UserAddressUpdate(w http.ResponseWriter, r *http.Request) {
	request := models.UserAddress{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in Decoding UserAddressUpdate Request :", err)
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
	repo := masterrepo.UserAddressInterface(&masterrepo.UserAddressStruct{})
	value, descreption, status := repo.UserAddressUpdate(&request)
	response := models.UserAddressResponses{
		Statuscode:  200,
		Status:      status,
		Value:     value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("Error in Marshal UserAddressUpdate Response:", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}}

func (UserAddress *UserAddressController) UserAddressDelete(w http.ResponseWriter, r *http.Request) {
	request := models.UserAddress{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Println("Error in Decoding UserAddressDelete Request :", err)
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
	repo := masterrepo.UserAddressInterface(&masterrepo.UserAddressStruct{})
   status, descreption := repo.UserAddressDelete(&request)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("Error in Marshal UserAddressDelete Response:", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
}

func (UserAddress *UserAddressController) UserAddressGetById(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["id"], 10, 64)
	// useraddressid := strconv.FormatInt(id, 10)
	if err != nil {
		fmt.Println("Error in Decoding  UserAddressGetById Request :", err)
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
	

	UserAddressStruct := models.UserAddress{
		Id: id,
	}
	if err != nil {
		fmt.Println(err)
	}
	repo := masterrepo.UserAddressInterface(&masterrepo.UserAddressStruct{})
	value, status, descreption := repo.UserAddressGetById(&UserAddressStruct)
	response := models.UserAddressResponses{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal UserAddressGetById Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
}

func (UserAddress *UserAddressController) UserAddressGetAll(w http.ResponseWriter, r *http.Request) {
	
	repo := masterrepo.UserAddressInterface(&masterrepo.UserAddressStruct{})
	value, status, descreption := repo.UserAddressGetAll()
	response := models.GetAllUserAddressResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal UserAddressGetAll Response :",err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}

func (UserAddress *UserAddressController) UserAddressGetAllCustomer(w http.ResponseWriter, r *http.Request) {
	request := mux.Vars(r)
	id, err := strconv.ParseInt(request["customerid"], 10, 64)
	// useraddressid := strconv.FormatInt(id, 10)
	if err != nil {
		fmt.Println("Error in Decoding  UserAddressGetById Request :", err)
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
	repo := masterrepo.UserAddressInterface(&masterrepo.UserAddressStruct{})
	value, status, descreption := repo.UserAddressGetAllCustomer(&id)
	response := models.GetAllUserAddressResponse{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal UserAddressGetAll Response :",err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
}