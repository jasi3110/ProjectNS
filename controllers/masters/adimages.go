package masters

import (
	"OnlineShop/models"
	masterrepo "OnlineShop/repos/masterRepo"
	"encoding/json"
	"fmt"
	"net/http"
)

type AdImageController struct {
}

func (image *AdImageController) AdImageCreate(w http.ResponseWriter, r *http.Request) {
	requst := models.AdImages{}
	err := json.NewDecoder(r.Body).Decode(&requst)
	if err != nil {
		fmt.Println("Error in Decoding AdImage Create Request :", err)
		response := models.CommanRespones{
			Statuscode:  200,
			Status:      false,
			Descreption: "Failed",
		}
		resp, err := json.Marshal(&response)
		if err != nil {
			fmt.Println("Error in Marshal AdImage Create Response:", err)
		}
		w.Header().Set("Content-Type", "Application/json")
		w.Write(resp)
	}else{
	repo := masterrepo.AdImagesInterface(&masterrepo.AdImagesStruct{})
	status, descreption := repo.AdImagesCreate(&requst)
	response := models.CommanRespones{
		Statuscode:  200,
		Status:      status,
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal AdImage Create Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}
}


func (image *AdImageController) AdImageGetall(w http.ResponseWriter, r *http.Request) {
	
	repo := masterrepo.AdImagesInterface(&masterrepo.AdImagesStruct{})
	value, status, descreption := repo.AdImageGetAll()
	response := models.ProductImageResponses{
		Statuscode:  200,
		Status:      status,
		Value:       value,
		Descreption: descreption,
	}
	respone, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal AdImageGetall Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(respone)
}
