package dashboard

import (
	"OnlineShop/models"
	"OnlineShop/repos"
	"OnlineShop/repos/masterRepo"
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
	
	Adimagesrepo := masterRepo.AdImagesInterface(&masterRepo.AdImagesStruct{})
	value, status0, descreption0 := Adimagesrepo.AdImageGetAll()


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

if !status0{
	fmt.Println(descreption0)
}
val22:=models.Dashboard{
	Id: 2,
	ViewType: 2,
	Data: nums,
	AdData:value,
}
result=append(result, val22)

val12:=models.Dashboard{
	Id: 3,
	Type: "Category",
	ViewType: 3,
	Data: nums,
}
result=append(result, val12)


	
	if !status1{
		fmt.Println(descreption1)
	}
	val1:=models.Dashboard{
		Id: 4,
		ViewType: 4,
		Data: value1,
	}
	result=append(result, val1)

	val13:=models.Dashboard{
		Id: 5,
		Type: "Discount Product",
		ViewType: 3,
		Data: nums,
	}
	result=append(result, val13)
	


	if !status2{
		fmt.Println(descreption2)
	}
	val2:=models.Dashboard{
		Id: 6,
		ViewType: 5,
		Data: value2,
	}
	result=append(result, val2)

	

	if !status3{
		fmt.Println(descreption3)
	}

	val14:=models.Dashboard{
		Id: 7,
		Type: "Products",
		ViewType: 3,
		Data: nums,
	}
	result=append(result, val14)
	val3:=models.Dashboard{
		Id: 8,
		ViewType: 5,
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

