package repos

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
	// "strconv"
)

type CartInterface interface {
	Createcart(obj *models.RCart) (bool, string)
	CartUpdate(obj *models.RCart) (string, bool)
	CartGetAll(obj *int64) (models.GetAllCart, bool, string)


	CartProductDelete(obj *models.RCart) (bool, string)
	CartDelete(obj *models.RCart) (bool, string)
}
type CartStruct struct {
}

func (cart *CartStruct) Createcart(obj *models.RCart) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Create cart")
	}
	query, err := Db.Query(`SELECT customerid,productid FROM "cart"`)
	if err != nil {
		fmt.Println("Error in CreateUser Checking User verfiy QueryRow :", err)
	}
	cartStruct := models.RCart{}
	for query.Next() {
		query.Scan(&cartStruct.Id, &cartStruct.Productid)
		if obj.Id == cartStruct.Id && obj.Productid == cartStruct.Productid {
			query:=`UPDATE "cart" SET quantity=$3 WHERE productid=$2 AND customerid=$1`
	_, err =Db.Exec(query,&obj.Id,&obj.Productid,&obj.Quantity)
	if err != nil {
		fmt.Println("Error in Create cart QueryRow :", err)
		return false, " Create cart Failed "
	}
	fmt.Println("ereeee")
	return true,  " Create cart Sucessfully "
		}
	}
	err = Db.QueryRow(`INSERT INTO "cart" (customerid,productid,quantity)values($1,$2,$3)RETURNING id`,
	 obj.Id,obj.Productid,obj.Quantity).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Create cart QueryRow :", err)
		return false, " Create cart Failed "
	}
	defer func() {
		Db.Close()
	}()
	fmt.Println("CART REPO")
	return true, " Create cart sucessfully "
}

func (cart *CartStruct) CartUpdate(obj *models.RCart) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in cart Update")
	}
	fmt.Println("",obj)
	query:=`UPDATE "cart" SET quantity=$3 WHERE productid=$2 AND customerid=$1`
	_, err :=Db.Exec(query,&obj.Id,&obj.Productid,&obj.Quantity)
	
	if err != nil {
		fmt.Println("Error in cart Update QueryRow :", err)
		return  "Update Failed", false
	}
	defer func() {

		Db.Close()
	}()
	return "Successfully Updated", true
}

func (cart *CartStruct) CartGetAll(obj *int64) (models.GetAllCart, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in cart GetAll")
	}
	result :=models.GetAllCart{}
	cartStruct:=models.Cart{}

	 query,err := Db.Query(`SELECT productid,quantity FROM "cart" WHERE customerid=$1`,obj)
	if err != nil {
		fmt.Println("Error in cart GetAll QueryRow :", err)
		return result, false, "failed"
	}
	// var productchannel chan models.ProductAll
	for query.Next() {
		err = query.Scan(
			&cartStruct.Product.Id,
			&cartStruct.Product.CartQuantity)
		if err != nil {
			fmt.Println("Error in cart GetAll QueryRow Scan :", err)
			return result, false, "failed"
		}
		
		productRepo:=ProductInterface(&ProductStruct{})
		value, _,_:= productRepo.GetProductById(&cartStruct.Product.Id)
cartStruct.Product.Id=value.Id
cartStruct.Product.Image=value.Image
cartStruct.Product.Name=value.Name
cartStruct.Product.Category=value.Category
cartStruct.Product.Quantity=value.Quantity
// cartStruct.Product.CartQuantity = cartStruct.Product.Quantity
cartStruct.Product.Unit=value.Unit
cartStruct.Product.Percentage=value.Percentage
cartStruct.Product.Price=value.Price
cartStruct.Product.CreatedOn=value.CreatedOn
// cartStruct.Product = value
 
// productqty, _ := strconv.ParseFloat(value.Quantity, 32)
// fmt.Println("",value.Price.Mrp*productqty)
// value.Price.Mrp =value.Price.Mrp*productqty
// value.Price.Nop=value.Price.Nop*productqty
result.Items  = result.Items + 1
// result.Productdiscoiunt =(result.Productdiscoiunt + (value.Price.Mrp - value.Price.Nop))*productqty
		result.Value = append(result.Value, cartStruct.Product)
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "successfully Completed"
}

func (cart *CartStruct) CartProductDelete(obj *models.RCart) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in Cart Product Delete ")
	}
	query,err :=Db.Query( "DELETE FROM cart  WHERE customerid=$1 and productid=$2 ",&obj.Id,&obj.Productid)
	
	for query.Next(){}
	if err != nil {
		fmt.Println("Error in Cart Delete QueryRow :", err)
		return false, "Cart Product Delete Failed"
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return true, "Cart Deleted Successfully "
}

func (cart *CartStruct) CartDelete(obj *models.RCart) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in Cart Delete ")
	}
	query,err :=Db.Query( `DELETE FROM "cart"  WHERE customerid=$1 `,&obj.Id)
	
	for query.Next(){}
	if err != nil {
		fmt.Println("Error in Cart Delete QueryRow :", err)
		return false, "Cart Delete Failed"
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return true, "Cart Deleted Successfully "
}