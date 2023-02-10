package repos

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
	
)

type CartInterface interface {
	Createcart(obj *models.RCart) (bool, string)
	// CartById(obj *models.Cart) (models.Cart, bool, string)
	CartGetAll(obj *int64) ([]models.Cart, bool, string)
	// CartUpdate(obj *models.Cart) (models.Cart, string, bool)
}
type CartStruct struct {
}

func (cart *CartStruct) Createcart(obj *models.RCart) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Create cart")
	}

	err := Db.QueryRow(`INSERT INTO "cart" (customerid,productid,quantity)values($1,$2,$3)RETURNING id`,
	 obj.Id,obj.Productid,obj.Quantity).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Create cart QueryRow :", err)
		return false, " Create cart Failed "
	}
	return true, "cart Successfully Created"
}

// func (cart *CartStruct) cartUpdate(obj *models.RCart) (string, bool) {
// 	Db, isconnceted := utls.OpenDbConnection()
// 	if !isconnceted {
// 		fmt.Println("DB Disconnceted in cart Update")
// 	}

// 	query:= `UPDATE "cart" SET quantity=$2 WHERE productid=$1`
// 	_, err := Db.Exec(query, &obj.Productid, &obj.Quantity)

// 	if err != nil {
// 		fmt.Println("Error in cart Upadte QueryRow :", err)
// 		return  "Update Failed", false
// 	}
// 	return "Successfully Updated", true
// }

// func (cart *CartStruct) cartById(obj *models.RCart) (models.Cart, bool, string) {
	// Db, isconnceted := utls.OpenDbConnection()
	// if !isconnceted {
	// 	fmt.Println("DB Disconnceted in cart GetById")
	// }
	// cartStruct := models.Cart{}

	// query, _ := Db.Prepare(`SELECT customerid,productid,quantity, from "cart" where id=$1`)

	// err := query.QueryRow(obj.Id).Scan(&obj.Id, &obj.Productid,obj.Quantity)

	// if err != nil {
	// 	fmt.Println("Error in cart GetById QueryRow :", err)
	// 	return cartStruct, false, "Error is founded in cart get by id"
	// }
	// return cartStruct, true, "cart get id successfully"
// }

func (cart *CartStruct) CartGetAll(obj *int64) ([]models.Cart, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in cart GetAll")
	}
	result := []models.Cart{}
	cartStruct:=models.Cart{}

	 query,err := Db.Query(`SELECT id,productid,quantity FROM "cart" WHERE customerid=$1`,obj)
	if err != nil {
		fmt.Println("Error in cart GetAll QueryRow :", err)
		return result, false, "failed"
	}
	for query.Next() {
		err = query.Scan(
			&cartStruct.Id,
			&cartStruct.Product.Id,
			&cartStruct.Product.Quantity)
		if err != nil {
			fmt.Println("Error in cart GetAll QueryRow Scan :", err)
			return result, false, "failed"
		}
		productRepo:=ProductInterface(&ProductStruct{})
value,status,descreption:=productRepo.GetProductById(&cartStruct.Product.Id)
value.Quantity = cartStruct.Product.Quantity
cartStruct.Product = value
// quantity := value.Quantity

cartStruct.Total = cartStruct.Total + value.Price.Nop
if !status {
	fmt.Println(descreption)
	return result, false, descreption
}
		result = append(result, cartStruct)
	}
	return result, true, "sucessfully Completed"
}


// func (cart *CartStruct) cartGetAllbyid() ([]models.Cart, bool, string) {
// 	Db, isConnected := utls.CreateDbConnection()
// 	if !isConnected {
// 		fmt.Println("DB Disconnceted in cart GetAll")
// 	}
// 	result := []models.Cart{}
// 	cartStruct:=models.Cart{}

// 	query, err := Db.Query(`SELECT id,name FROM "cart"`)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	for query.Next() {
// 		err := query.Scan(
// 			&cartStruct.Id,
// 			&cartStruct.Name,
// 		)
// 		if err != nil {
// 			fmt.Println("Error in cart GetAll QueryRow :", err)
// 			return result, false, "failed to  Get All cart Data"
// 		}
// 		result = append(result, cartStruct )
// 	}
// 	return result, true, "sucessfully Completed"
// }
