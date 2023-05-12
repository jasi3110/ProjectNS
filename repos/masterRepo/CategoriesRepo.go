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

	err := Db.QueryRow(`INSERT INTO "category" (name)values($1)RETURNING id`, obj.Name).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Create Category QueryRow :", err)
		return false, " Create Category Failed "
	}
	defer Db.Close()
	return true, "Category Successfully Created"
}



func (category *CategoryStruct) CategoryUpdate(obj *models.Category) (models.Category, string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Category Update")
	}
	query:=`UPDATE "category" SET name=$2 WHERE id=$1 AND isdeleted=0`
	_, err :=Db.Exec(query,&obj.Id, &obj.Name)
	
	if err != nil {
		fmt.Println("Error in Category Upadte QueryRow :", err)
		return *obj, "Update Failed", false
	}
	defer Db.Close()
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
	defer Db.Close()
	return true, "Category Successfully Completed"
}



func (category *CategoryStruct) CategoryById(obj *models.Category) (models.Category, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Category GetById")
	}
	categoryStruct := models.Category{}

	query, err:= Db.Prepare(`SELECT id,name from "category" where id=$1`)
	if err != nil {
		fmt.Println("Error in Category GetById QueryRow :", err)
		return categoryStruct, false, "Failed"
	}
	err = query.QueryRow(obj.Id).Scan(&categoryStruct.Id, &categoryStruct.Name)

	if err != nil {
		fmt.Println("Error in Category GetById QueryRow Scan:", err)
		return categoryStruct, false, "Failed"
	}
	defer Db.Close()
	return categoryStruct, true, "successfully Completed"
}



func (category *CategoryStruct) CategoryGetAll() ([]models.Category, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Category GetAll")
	}
	result := []models.Category{}
	categoryStruct:=models.Category{}

	query, err := Db.Query(`SELECT id,name FROM "category"WHERE isdeleted=0`)
	if err != nil {
		fmt.Println("Error in Category GetAll QueryRow :", err)
		return result, false, "Failed"
	}

	for query.Next() {
		err := query.Scan(
			&categoryStruct.Id,
			&categoryStruct.Name,
		)
		if err != nil {
			fmt.Println("Error in Category GetAll QueryRow Scan :", err)
			return result, false, "Failed"
		}
		result = append(result, categoryStruct )
	}
	defer Db.Close()
	return result, true, "successfully Completed"
}



func (category *CategoryStruct) CategoryGetAllbyid() ([]models.Category, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Category GetAll")
	}
	result := []models.Category{}
	categoryStruct:=models.Category{}

	query, err := Db.Query(`SELECT id,name FROM "category" WHERE isdeleted=0`)
	if err != nil {
		fmt.Println("Error in Category GetAll QueryRow :", err)
		return result, false, "Failed"
	}

	for query.Next() {
		err := query.Scan(
			&categoryStruct.Id,
			&categoryStruct.Name,
		)
		if err != nil {
			fmt.Println("Error in Category GetAll QueryRow :", err)
			return result, false, "Failed"
		}
		result = append(result, categoryStruct )
	}
	defer Db.Close()
	return result, true, "successfully Completed"
}
