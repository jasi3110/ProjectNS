package masterRepo

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
)

type CategoryInterface interface {
	CreateCategory(obj *models.Category) (bool, string)
	CategoryById(obj *models.Category) (models.Category, bool, string)
	CategoryGetAll() ([]models.Category, bool, string)
	CategoryUpdate(obj *models.Category) (models.Category, string, bool)
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
	return true, "Category Sucessfully Created"
}

func (category *CategoryStruct) CategoryUpdate(obj *models.Category) (models.Category, string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Category Update")
	}

	query:= `UPDATE "category" SET name=$2 WHERE id=$1`
	_, err := Db.Exec(query, &obj.Id, &obj.Name)

	if err != nil {
		fmt.Println("Error in Category Upadte QueryRow :", err)
		return *obj, "Update Failed", false
	}
	return *obj, "Sucessfully Updated", true
}

func (category *CategoryStruct) CategoryById(obj *models.Category) (models.Category, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Category GetById")
	}
	categoryStruct := models.Category{}

	query, _ := Db.Prepare(`SELECT id,name from "category" where id=$1`)

	err := query.QueryRow(obj.Id).Scan(&categoryStruct.Id, &categoryStruct.Name)

	if err != nil {
		fmt.Println("Error in Category GetById QueryRow :", err)
		return categoryStruct, false, "Error is founded in category get by id"
	}
	return categoryStruct, true, "category get id successfully"
}

func (category *CategoryStruct) CategoryGetAll() ([]models.Category, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Category GetAll")
	}
	result := []models.Category{}
	categoryStruct:=models.Category{}

	query, err := Db.Query(`SELECT id,name FROM "category"`)
	if err != nil {
		fmt.Println(err)
	}

	for query.Next() {
		err := query.Scan(
			&categoryStruct.Id,
			&categoryStruct.Name,
		)
		if err != nil {
			fmt.Println("Error in Category GetAll QueryRow :", err)
			return result, false, "failed to  Get All Category Data"
		}
		result = append(result, categoryStruct )
	}
	return result, true, "sucessfully Completed"
}


func (category *CategoryStruct) CategoryGetAllbyid() ([]models.Category, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Category GetAll")
	}
	result := []models.Category{}
	categoryStruct:=models.Category{}

	query, err := Db.Query(`SELECT id,name FROM "category"`)
	if err != nil {
		fmt.Println(err)
	}

	for query.Next() {
		err := query.Scan(
			&categoryStruct.Id,
			&categoryStruct.Name,
		)
		if err != nil {
			fmt.Println("Error in Category GetAll QueryRow :", err)
			return result, false, "failed to  Get All Category Data"
		}
		result = append(result, categoryStruct )
	}
	return result, true, "sucessfully Completed"
}
