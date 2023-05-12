package masterRepo

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
)

type UnitInterface interface {
	CreateUnit(obj *models.Unit) (bool, string)
	UnityById(obj *models.Unit) (models.Unit, bool, string)
	UnitByAll() ([]models.Unit, bool, string)
	UnitUpdate(obj *models.Unit) (string, bool)
}
type UnitStruct struct {
}

func (unit *UnitStruct) CreateUnit(obj *models.Unit) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Create Unit")
	}
	err := Db.QueryRow(`INSERT INTO "unit"(item)values($1) RETURNING id`, obj.Item).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in  Create Unit QueryRow :", err)
		return false, "Create unit is Failed"
	}
	defer func() {
		Db.Close()
	}()
	return true, "Create Unit Sucessfully"
}

func (unit *UnitStruct) UnitUpdate(obj *models.Unit) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Unit Update")
	}

	query := `UPDATE "unit" SET item=$2 WHERE id=$1`
	_, err := Db.Exec(query,&obj.Id,&obj.Item)

	if err != nil {
		fmt.Println("Error in Unit Update QueryRow :", err)
		return "Update Failed", false
	}
	defer func() {
		Db.Close()
	}()
	return "Sucessfully Updated", true
}

func (unit *UnitStruct) UnityById(obj *models.Unit) (models.Unit, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in UnitById")
	}
	UnitStruct:= models.Unit{}

	query, err := Db.Prepare(`SELECT id,item from "unit" where id=$1`)

	if err != nil {
		fmt.Println("Error in UnitById QueryRow :", err)
		return UnitStruct, false, "Error is founded on get by id on unit"
	}

	err = query.QueryRow(obj.Id).Scan(&UnitStruct.Id, &UnitStruct.Item)
	if err != nil {
		fmt.Println("Error in UnitById QueryRow :", err)
		return UnitStruct, false, "Error is founded on get by id on unit"
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return UnitStruct, true, "sucessfully completed"
}

func (unit *UnitStruct) UnitByAll() ([]models.Unit, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Unit UnitByAll ")
	}
	result := []models.Unit{}
	unitStruct:=models.Unit{}

	query, err := Db.Query(`SELECT id,item FROM "unit"`)
	if err != nil {
		fmt.Println(err)
	}

	for query.Next() {
		err := query.Scan(
			&unitStruct.Id,
			&unitStruct.Item,
		)
		if err != nil {
			fmt.Println("Error in Unit GetbyALL QueryRow :", err)
			return result, false, "failed to  Get All Unit Data"
		}
		result = append(result, unitStruct)
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "sucessfully Completed"
}
  