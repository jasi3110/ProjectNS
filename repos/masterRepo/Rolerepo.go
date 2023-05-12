package masterRepo

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
)

type RoleInterface interface {
	CreateRole(obj *models.Role) (bool, string)
	RoleById(obj *models.Role) (models.Role, bool, string)
	RoleGetAll() ([]models.Role, bool, string)
	RoleUpdate(obj *models.Role) (string, bool)
}
type RoleStruct struct {
}

func (role *RoleStruct) CreateRole(obj *models.Role) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Create Role ")
	}

	err := Db.QueryRow(`INSERT INTO "role" (type)values($1)RETURNING id`, obj.Type).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Create Role QueryRow:", err)
		return false, " Create Role Failed "
	}
	defer func() {
		Db.Close()
	}()
	return true, "Role  Sucessfully Created"
}

func (role *RoleStruct) RoleUpdate(obj *models.Role) (string, bool) {
	Db, isconncet := utls.OpenDbConnection()
	if !isconncet {
		fmt.Println("DB Disconnceted in Role Update")
	}

	query := `UPDATE "role" SET type=$2 WHERE id=$1`
	_, err := Db.Exec(query, &obj.Id, &obj.Type)

	if err != nil {
		fmt.Println("Error in Role Update QueryRow :", err)
		return "Update Failed", false
	}
	defer func() {
		Db.Close()
	}()
	return "Sucessfully Updated", true
}

func (role *RoleStruct) RoleById(obj *models.Role) (models.Role, bool, string) {
	Db, conncet := utls.OpenDbConnection()
	if !conncet {
		fmt.Println("DB Disconnceted in RoleById")
	}
	roleStruct := models.Role{}

	query, _ := Db.Prepare(`SELECT id,type from "role" where id=$1`)

	err := query.QueryRow(obj.Id).Scan(&roleStruct.Id, &roleStruct.Type)

	if err != nil {
		fmt.Println("Error in RoleById QueryRow :", err)
		return roleStruct, false, "Role Get By ID failed"
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return roleStruct, true, "Sucessfully Completed"
}

func (role *RoleStruct) RoleGetAll() ([]models.Role, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Role GetAll")
	}
	result := []models.Role{}
	roleStruct := models.Role{}

	query, err := Db.Query(`SELECT id,type FROM "role"`)
	if err != nil {
		fmt.Println("Error in Role GetAll Queryrow :", err)
	}

	for query.Next() {
		err := query.Scan(
			&roleStruct.Id,
			&roleStruct.Type,
		)
		if err != nil {
			fmt.Println("Error is founded :", err)
			return result, false, "failed to  Get All Role Data"
		}
		result = append(result, roleStruct)
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "sucessfully Completed"
}
