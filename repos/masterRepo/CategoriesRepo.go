package masterRepo

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
)

type CategoryInterface interface {
	CreateCategory(obj *models.Category) (bool, string)
	CategoryUpdate(obj *models.Category) (models.Category, string, bool)
	CategoryDelete(obj *models.User) (bool, string)


	CategoryById(obj *models.Category) (models.Category, bool, string)
	CategoryGetAll() ([]models.Category, bool, string)
	CategoryGetAllbyid() ([]models.Category, bool, string) 
	
}
type CategoryStruct struct {
}

func (category *CategoryStruct) CreateCategory(obj *models.Category) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Create Category")
	}
	imageurl:= models.Imageurl(obj.Image)
	err := Db.QueryRow(`INSERT INTO "category" (name,image)values($1,$2)RETURNING id`, obj.Name,imageurl).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Create Category QueryRow :", err)
		return false, " Create Category Failed "
	}
	defer func() {
		Db.Close()
	}()
	return true, "Category Successfully Created"
}

func (category *CategoryStruct) CategoryUpdate(obj *models.Category) (models.Category, string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Category Update")
	}
	imageurl:= models.Imageurl(obj.Image)
	query:=`UPDATE "category" SET name=$2,image=$3 WHERE id=$1 AND isdeleted=0`
	_, err :=Db.Exec(query,&obj.Id, &obj.Name,&imageurl)
	
	if err != nil {
		fmt.Println("Error in Category Upadte QueryRow :", err)
		return *obj, "Update Failed", false
	}
	defer func() {
		Db.Close()
	}()
	return *obj, "Successfully Updated", true
}

func (category *CategoryStruct) CategoryDelete(obj *models.User) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in Category Delete ")
	}
	query:=`UPDATE "category" SET isdeleted=1 WHERE id=$1 and isdeleted=0`
	_, err :=Db.Exec(query,&obj.Id)
	
	if err != nil {
		fmt.Println("Error in category Delete QueryRow :", err)
		return false, "Failed"
	}
	defer func() {
		Db.Close()
	}()
	return true, "Category Successfully Completed"
}

func (category *CategoryStruct) CategoryById(obj *models.Category) (models.Category, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Category GetById")
	}
	categoryStruct := models.Category{}

	query, err:= Db.Prepare(`SELECT id,name,image from "category" where id=$1`)
	if err != nil {
		fmt.Println("Error in Category GetById QueryRow :", err)
		return categoryStruct, false, "Failed"
	}
	
	err = query.QueryRow(obj.Id).Scan(&categoryStruct.Id, &categoryStruct.Name,&categoryStruct.Image)
	basicURL := "https://drive.google.com/uc?export=view&id="
	categoryStruct.Image = basicURL + categoryStruct.Image
	if err != nil {
		fmt.Println("Error in Category GetById QueryRow Scan:", err)
		return categoryStruct, false, "Failed"
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return categoryStruct, true, "successfully Completed"
}

func (category *CategoryStruct) CategoryGetAll() ([]models.Category, bool, string) {
	// CHECKING DATABASE 
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Category GetAll")
	}

    // CREATING  SLICE CATEGORY MODEL VARIABLE AND CATEGORY MODEL VARIABLE --
	// FOR APPEND EACH ROW OF CATEGORY TABLE

	result := []models.Category{}
	categoryStruct:=models.Category{}


	// A QUERY FOR  CATEGORY ID,NAME,IMAGE FROM CATEGORY TABLE

	query, err := Db.Query(`SELECT id,name,image FROM "category"WHERE isdeleted=0`)
	if err != nil {
		fmt.Println("Error in Category GetAll QueryRow :", err)
		return result, false, "Failed"
	}
	// SCANNING VALUES FROM CATEGORY TABLE
	for query.Next() {err := query.Scan(&categoryStruct.Id,&categoryStruct.Name,&categoryStruct.Image)
		if err != nil {
			fmt.Println("Error in Category GetAll QueryRow Scan :", err)
			return result, false, "Failed"
		}
		basicURL := "https://drive.google.com/uc?export=view&id="
		categoryStruct.Image = basicURL + categoryStruct.Image
	// APPENDING EACH CATEGORY TABLE ROW IN RESULT VARIABLE
		result = append(result, categoryStruct )
	}
    
	// DATABASE CLOSING IN DEFER FUNCTION
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "successfully Completed"
}

func (category *CategoryStruct) CategoryGetAllbyid() ([]models.Category, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Category GetAll")
	}
	result := []models.Category{}
	categoryStruct:=models.Category{}

	query, err := Db.Query(`SELECT id,name,image FROM "category" WHERE isdeleted=0`)
	if err != nil {
		fmt.Println("Error in Category GetAll QueryRow :", err)
		return result, false, "Failed"
	}

	for query.Next() {
		err := query.Scan(&categoryStruct.Id,&categoryStruct.Name,&categoryStruct.Image)
		if err != nil {
			fmt.Println("Error in Category GetAll QueryRow :", err)
			return result, false, "Failed"
		}
		basicURL := "https://drive.google.com/uc?export=view&id="
	categoryStruct.Image = basicURL + categoryStruct.Image
		result = append(result, categoryStruct )
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "successfully Completed"
}
