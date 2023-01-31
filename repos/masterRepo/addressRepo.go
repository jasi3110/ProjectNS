package masterRepo

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
)

type UserAddressInterface interface {
	UserAddressCreate(obj *models.UserAddress) (bool, string)
	UserAddressGetById(obj *models.UserAddress) (models.UserAddress, bool, string)
	UserAddressGetAll() ([]models.UserAddress, bool, string)
	UserAddressUpdate(obj *models.UserAddress) (models.UserAddress, string, bool)
}
type UserAddressStruct struct {
}

func (UserAddress *UserAddressStruct) UserAddressCreate(obj *models.UserAddress) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in UserAddress Create")
	}

	err := Db.QueryRow(`INSERT INTO "useraddress" customerid,name,address values($1,$2,$3)RETURNING id`, obj.Customerid,&obj.Name,obj.Address).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Create Category QueryRow :", err)
		return false, " UserAddressCreate Failed "
	}
	return true, "UserAddressCreate Sucessfully Created"
}

func (UserAddress *UserAddressStruct) UserAddressUpdate(obj *models.UserAddress) (models.UserAddress, string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in UserAddress Update")
	}

	query:= `UPDATE "useraddress" SET name=$2, address=$3 WHERE id=$1`
	_, err := Db.Exec(query, &obj.Id, &obj.Name,obj.Address)

	if err != nil {
		fmt.Println("Error in UserAddress Upadte QueryRow :", err)
		return *obj, "UserAddress Update Failed", false
	}
	return *obj, "UserAddress Sucessfully Updated", true
}

func (UserAddress *UserAddressStruct) UserAddressGetById(obj *models.UserAddress) (models.UserAddress, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in UserAddress GetById")
	}
	UserAddressStruct := models.UserAddress{}

	query, _ := Db.Prepare(`SELECT id,customerid,name,address from "useraddress" where id=$1`)

	err := query.QueryRow(obj.Id).Scan(&UserAddressStruct.Id,UserAddressStruct.Customerid, &UserAddressStruct.Name,&UserAddressStruct.Address)

	if err != nil {
		fmt.Println("Error in UserAddress GetById QueryRow :", err)
		return UserAddressStruct, false, "Error is founded in UserAddress GetById"
	}
	return UserAddressStruct, true, "UserAddressGetByid successfully"
}

func (UserAddress *UserAddressStruct) UserAddressGetAll() ([]models.UserAddress, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in UserAddress GetAll")
	}
	result := []models.UserAddress{}
	UserAddressStruct:=models.UserAddress{}

	query, err := Db.Query(`SELECT id,customerid,name,address FROM "useraddress"`)
	if err != nil {
		fmt.Println(err)
	}

	for query.Next() {
		err := query.Scan(
			&UserAddressStruct.Id,
			&UserAddressStruct.Customerid,
			&UserAddressStruct.Name,
			&UserAddressStruct.Address,
		)
		if err != nil {
			fmt.Println("Error in UserAddress GetAll QueryRow :", err)
			return result, false, "UserAddress GetAll Failed"
		}
		result = append(result, UserAddressStruct )
	}
	return result, true, "sucessfully Completed"
}
