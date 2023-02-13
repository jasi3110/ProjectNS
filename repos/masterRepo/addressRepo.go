package masterRepo

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
)

type UserAddressInterface interface {
	UserAddressCreate(obj *models.UserAddress) (bool, string)
	UserAddressUpdate(obj *models.UserAddress) (models.UserAddress, string, bool)
	UserAddressDelete(obj *models.User) (bool, string)


	UserAddressGetById(obj *models.UserAddress) (models.UserAddress, bool, string)
	UserAddressGetAll() ([]models.UserAddress, bool, string)
	UserAddressGetAllCustomer(obj *int64) ([]models.UserAddress, bool, string)
}
type UserAddressStruct struct {
}

func (UserAddress *UserAddressStruct) UserAddressCreate(obj *models.UserAddress) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in UserAddress Create")
	}

	err := Db.QueryRow(`INSERT INTO "useraddress" (customerid,name,address) values($1,$2,$3)RETURNING id`, obj.Customerid,&obj.Name,obj.Address).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Create Category QueryRow :", err)
		return false, " Address Create Failed "
	}
	defer Db.Close()
	return true, "Address Create Successfully"
}




func (UserAddress *UserAddressStruct) UserAddressUpdate(obj *models.UserAddress) (models.UserAddress, string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in User Address Update")
	}
	err :=Db.QueryRow(`UPDATE "useraddress" SET name=$2,address=$3 WHERE id=$1 and isdeleted=0`,
	&obj.Id, &obj.Name,obj.Address)

	if err != nil {
		fmt.Println("Error in UserAddress Upadte QueryRow :", err)
		return *obj, "Update Failed", false
	}
	defer Db.Close()
	return *obj, "Address Successfully Updated", true
}



func (UserAddress *UserAddressStruct) UserAddressDelete(obj *models.User) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in Address Delete ")
	}
	err :=Db.QueryRow( `UPDATE "useraddress" SET isdeleted=1 WHERE id=$1 and isdeleted=0`,&obj.Id)
	
	if err != nil {
		fmt.Println("Error in Address Delete QueryRow :", err)
		return false, "Failed"
	}
	defer Db.Close()
	return true, "Address Deleted Successfully Completed"
}



func (UserAddress *UserAddressStruct) UserAddressGetById(obj *models.UserAddress) (models.UserAddress, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in User Address GetById")
	}
	UserAddressStruct := models.UserAddress{}

	query, _ := Db.Prepare(`SELECT id,customerid,name,address from "useraddress" where id=$1`)

	err := query.QueryRow(obj.Id).Scan(&UserAddressStruct.Id,&UserAddressStruct.Customerid, &UserAddressStruct.Name,&UserAddressStruct.Address)

	if err != nil {
		fmt.Println("Error in UserAddress GetById QueryRow :", err)
		return UserAddressStruct, false, "Error is founded in UserAddress GetById"
	}
	defer Db.Close()
	return UserAddressStruct, true, "successfully Completed"
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
			fmt.Println("Error in UserAddress GetAll QueryRow Scan:", err)
			return result, false, "Failed"
		}
		result = append(result, UserAddressStruct )
	}
	defer Db.Close()
	return result, true, "Successfully Completed"
}



func (UserAddress *UserAddressStruct) UserAddressGetAllCustomer(obj *int64) ([]models.UserAddress, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in UserAddress GetAll")
	}
	result := []models.UserAddress{}
	UserAddressStruct:=models.UserAddress{}

	query, err := Db.Query(`SELECT id,customerid,name,address FROM "useraddress" WHERE customerid=$1 and isdeleted=0`,obj)
	if err != nil {
		fmt.Println("Error in UserAddress GetAll QueryRow :", err)
		return result, false, "UserAddress GetAll Failed"
	}

	for query.Next() {
		err := query.Scan(
			&UserAddressStruct.Id,
			&UserAddressStruct.Customerid,
			&UserAddressStruct.Name,
			&UserAddressStruct.Address,
		)
		if err != nil {
			fmt.Println("Error in UserAddress GetAll QueryRow Scan :", err)
			return result, false, "Failed"
		}
		result = append(result, UserAddressStruct )
	}
	defer Db.Close()
	return result, true, "successfully Completed"
}


