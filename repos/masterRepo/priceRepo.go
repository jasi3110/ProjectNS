package masterRepo

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
)

type PriceInterface interface {
	CreatePrices(obj *models.Price) (bool, string, models.Price)
	// RoleById(obj *models.Role) (models.Role, bool, string)
	// RoleGetAll() ([]models.Role, bool, string)
	// RoleUpdate(obj *models.Role) (string, bool)
}
type PriceStruct struct {
}

func (price *PriceStruct) CreatePrices(obj *models.Price) (bool, string, models.Price) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Create price ")
	}

	err:= Db.QueryRow(`INSERT INTO "price" (
		productid,
		productprice,
		createdon)values($1,$2,$3)RETURNING id `,
	obj.ProductId,obj.ProductPrice,utls.GetCurrentDateTime(),).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Create prices QueryRow:", err)
		return false, " Create price Failed ", *obj
	}
	return true, "price  Sucessfully Created", *obj
}

// func (role *RoleStruct) RoleUpdate(obj *models.Role) (string, bool) {
// 	myDb, isconncet := utls.OpenDbConnection()
// 	if !isconncet {
// 		fmt.Println("DB Disconnceted in Role Update")
// 	}

// 	query := `UPDATE "role" SET type=$2 WHERE id=$1`
// 	_, err := myDb.Exec(query, &obj.Id, &obj.Type)

// 	if err != nil {
// 		fmt.Println("Error in Role Update QueryRow :", err)
// 		return "Update Failed", false
// 	}
// 	return "Sucessfully Updated", true
// }

// func (role *RoleStruct) RoleById(obj *models.Role) (models.Role, bool, string) {
// 	mydb, conncet := utls.OpenDbConnection()
// 	if !conncet {
// 		fmt.Println("DB Disconnceted in RoleById")
// 	}
// 	roleStruct := models.Role{}

// 	query, _ := mydb.Prepare(`SELECT id,type from "role" where id=$1`)

// 	err := query.QueryRow(obj.Id).Scan(&roleStruct.Id, &roleStruct.Type)

// 	if err != nil {
// 		fmt.Println("Error in RoleById QueryRow :", err)
// 		return roleStruct, false, "Role Get By ID failed"
// 	}
// 	return roleStruct, true, "Sucessfully Completed"
// }

// func (role *RoleStruct) RoleGetAll() ([]models.Role, bool, string) {
// 	Db, isConnected := utls.OpenDbConnection()
// 	if !isConnected {
// 		fmt.Println("DB Disconnceted in Role GetAll")
// 	}
// 	result := []models.Role{}
// 	roleStruct := models.Role{}

// 	res, err := Db.Query(`SELECT id,type FROM "role"`)
// 	if err != nil {
// 		fmt.Println("Error in Role GetAll Queryrow :", err)
// 	}

// 	for res.Next() {
// 		err := res.Scan(
// 			&roleStruct.Id,
// 			&roleStruct.Type,
// 		)
// 		if err != nil {
// 			fmt.Println("Error is founded :", err)
// 			return result, false, "failed to  Get All Role Data"
// 		}
// 		result = append(result, roleStruct)
// 	}
// 	return result, true, "sucessfully Completed"
// }
