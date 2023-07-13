package repos

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
)

type CartInterface interface {
	Createcart(obj *models.RCart) (bool, string)
	CartUpdate(obj *models.GetCart) (string, bool)
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
	query, err := Db.Query(`SELECT productid FROM "cart" WHERE customerid=$1`, obj.Id)
	if err != nil {
		fmt.Println("Error in Create Cart Checking User verfiy QueryRow :", err)
		return false, " Create cart Failed "
	}
	cartStruct := models.RCart{}
	for query.Next() {
		query.Scan(&cartStruct.Productid)
		if obj.Productid == cartStruct.Productid {
			query := `UPDATE "cart" SET quantity=$3 WHERE productid=$2 AND customerid=$1`
			_, err = Db.Exec(query, &obj.Id, &obj.Productid, &obj.Quantity)
			if err != nil {
				fmt.Println("Error in Create cart QueryRow :", err)
				return false, " Create cart Failed "
			}
			return true, " Create cart Sucessfully "
		}
	}
	err = Db.QueryRow(`INSERT INTO "cart" (customerid,productid,quantity)values($1,$2,$3)RETURNING id`,
		obj.Id, obj.Productid, obj.Quantity).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Create cart QueryRow :", err)
		return false, " Create cart Failed "
	}
	defer func() {
		Db.Close()
	}()
	return true, " Create cart sucessfully "
}

func (cart *CartStruct) CartUpdate(obj *models.GetCart) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in cart Update")
	}
	fmt.Println("", obj)
	for _, product := range obj.Products {
	query := `UPDATE "cart" SET quantity=$3 WHERE productid=$2 AND customerid=$1`
	_, err := Db.Exec(query, &obj.Customerid,&product.Productid,&product.Quantity)

	if err != nil {
		fmt.Println("Error in cart Update QueryRow :", err)
		return "Update Failed", false
	}
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
	result := models.GetAllCart{}
	cartStruct := models.CartProductAll{}

	query, err := Db.Query(`SELECT productid,quantity FROM "cart" WHERE customerid=$1`, obj)
	if err != nil {
		fmt.Println("Error in cart GetAll QueryRow :", err)
		return result, false, "Failed"
	}

	for query.Next() {err = query.Scan(&cartStruct.Id,&cartStruct.CartQuantity)
		if err != nil {
			fmt.Println("Error in cart GetAll QueryRow Scan :", err)
			return result, false, "Failed"
		}

		productRepo := ProductInterface(&ProductStruct{})
		value, status, _ := productRepo.GetProductById(&cartStruct.Id)
		if !status {
			fmt.Println("Error in cart Product Gettting Data QueryRow Scan :", err)
			return result, false, "Failed"
		}
		cartStruct.Id=value.Id
		cartStruct.Image=value.Image
		cartStruct.Name=value.Name
		cartStruct.Category=value.Category
		cartStruct.Quantity=value.Quantity
		cartStruct.Unit=value.Unit
		cartStruct.Percentage=value.Percentage
		cartStruct.Price=value.Price
		cartStruct.CreatedOn=value.CreatedOn

		result.Items = result.Items + 1
		result.Value = append(result.Value,cartStruct)
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
	fmt.Println("details :",obj)
	query, err := Db.Query("DELETE FROM cart  WHERE customerid=$1 and productid=$2 ", &obj.Id, &obj.Productid)

	for query.Next() {
	}
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
	query, err := Db.Query(`DELETE FROM "cart"  WHERE customerid=$1 `, &obj.Id)

	for query.Next() {
	}
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
