package repos

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
	"strconv"
)

type CartInterface interface {
	Createcart(obj *models.RCart) (bool, string)
	CartUpdate(obj *models.RCart) (string, bool)
	CartGetAll(obj *int64) ([]models.Cart, bool, string)
	CartDelete(obj *models.User) (bool, string)
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



func (cart *CartStruct) CartUpdate(obj *models.RCart) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in cart Update")
	}
	err :=Db.QueryRow(`UPDATE "cart" SET quantity=$3 WHERE productid=$2 AND customerid=$1`,
	&obj.Id,&obj.Productid, &obj.Quantity)

	if err != nil {
		fmt.Println("Error in cart Upadte QueryRow :", err)
		return  "Update Failed", false
	}
	defer Db.Close()
	return "Successfully Updated", true
}



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

productqty, _ := strconv.ParseFloat(value.Quantity, 32)


cartStruct.Total = cartStruct.Total + (value.Price.Nop*productqty)
cartStruct.Items  = cartStruct.Items + 1
cartStruct.Productdiscoiunt =cartStruct.Productdiscoiunt + (value.Price.Mrp + value.Price.Nop)
if !status {
	fmt.Println(descreption)
	return result, false, descreption
}
		result = append(result, cartStruct)
	}
	defer Db.Close()
	return result, true, "sucessfully Completed"
}


func (cart *CartStruct) CartDelete(obj *models.User) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in Product Delete ")
	}
	err :=Db.QueryRow( `DELETE "cart"  WHERE id=$1 `,&obj.Id)
	
	if err != nil {
		fmt.Println("Error in Cart Delete QueryRow :", err)
		return false, "Cart Delete Failed"
	}
	defer Db.Close()
	return true, "Cart Deleted Successfully "
}