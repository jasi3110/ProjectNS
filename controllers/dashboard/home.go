package dashboard

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	"OnlineShop/repos/masterRepo"
	// "crypto/rand"bv 
	"encoding/json"
	"fmt"
	"net/http"
)

type Dashboard struct {
}

func (dashboard *Dashboard) Homepage(w http.ResponseWriter, r *http.Request){
	// CREATING A DASHBOARD SLICE 
	result:=[]models.Dashboard{}
	nums := []models.ProductAll{}

	categoryrepo:= masterRepo.CategoryInterface(&masterRepo.CategoryStruct{})
	value1,status1,descreption1:=categoryrepo.CategoryGetAll()

	discountRepo := repos.DiscountInterface(&repos.DiscountStruct{})
	value2, status2,descreption2 := discountRepo.DiscountGetAll()

	productrepo := repos.ProductInterface(&repos.ProductStruct{})
	value3,status3,descreption3:= productrepo.ProductGetAll()

val:=models.Dashboard{
	Id: 1,
	ViewType: 1,
	Data:nums,
}
result=append(result, val)

val12:=models.Dashboard{
	Id: 2,
	Type: "Category",
	ViewType: 2,
	Data: nums,
}
result=append(result, val12)


	
	if !status1{
		fmt.Println(descreption1)
	}
	val1:=models.Dashboard{
		Id: 3,
		ViewType: 3,
		Data: value1,
	}
	result=append(result, val1)

	val13:=models.Dashboard{
		Id: 4,
		Type: "Discount Product",
		ViewType: 2,
		Data: nums,
	}
	result=append(result, val13)
	


	if !status2{
		fmt.Println(descreption2)
	}
	val2:=models.Dashboard{
		Id: 5,
		ViewType: 4,
		Data: value2,
	}
	result=append(result, val2)

	

	if !status3{
		fmt.Println(descreption3)
	}

	val14:=models.Dashboard{
		Id: 6,
		Type: "Products",
		ViewType: 2,
		Data: nums,
	}
	result=append(result, val14)
	val3:=models.Dashboard{
		Id: 7,
		ViewType: 4,
		Data: value3,
	}
	result=append(result, val3)


	response := models.CommanResponesDashboard{
		Statuscode:  200,
		Status:      status3,
		Descreption: descreption3,
		Value: result,
	}

	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal Dashboard in Homepage Response :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}


func (dashboard *Dashboard) ProductImageGetAll(w http.ResponseWriter, r *http.Request) {

	repo := masterRepo.ProductImageInterface(&masterRepo.ProductImageStruct{})
	status,descreption:=repo.ProductImageGetall()
	response := models.GetAllProductResponse{
		Statuscode:   200,
		Status:      status,
		
		Descreption: descreption,
	}
	resp, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error in Marshal ProductGetAll Request :", err)
	}
	w.Header().Set("Content-Type", "Application/json")
	w.Write(resp)
}