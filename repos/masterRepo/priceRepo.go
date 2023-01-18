package masterRepo

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
)

type PriceInterface interface {
	CreatePrice(obj *models.Price) (bool, string, models.Price)
	PriceById(obj *models.Price) (models.Price, bool, string)
	PriceByDate(obj *models.Price) (models.Price, bool, string)
	PriceGetAll() ([]models.Price, bool, string)
	PriceUpdate(obj *models.Price) (string, bool)
	PriceProductGetAll(obj *models.Price) ([]models.Price, bool, string)
}
type PriceStruct struct {
}

func (price *PriceStruct) CreatePrice(obj *models.Price) (bool, string, models.Price) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Create price ")
	}

	err:= Db.QueryRow(`INSERT INTO "price" (
		productid,
		productprice,
		createdon)values($1,$2,$3)RETURNING id `,
	obj.ProductId,obj.ProductPrice,utls.GetCurrentDate(),).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Createprice QueryRow:", err)
		return false, " Createprice Failed ", *obj
	}
	return true, "price  Sucessfully Created", *obj
}

func (price *PriceStruct) PriceUpdate(obj *models.Price) (string, bool) {
 Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in PriceUpdate")
	}

	query := `UPDATE "price" SET productprice=$2,productid=$3 WHERE id=$1`
	_, err := Db.Exec(query, &obj.Id, &obj.ProductPrice,&obj.ProductId)

	if err != nil {
		fmt.Println("Error in PriceUpdate QueryRow :", err)
		return "Price Update Failed", false
	}
	return "Sucessfully Updated", true
}

func (price *PriceStruct) PriceByDate(obj *models.Price) (models.Price, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in PriceByDate")
	}
	

	query, _ := Db.Prepare(`SELECT id,productprice from "price" where createdon=$1 AND productid=$2`)

	err := query.QueryRow(obj.ProductId, obj.Createdon).Scan(&obj.Id, &obj.ProductId,obj.ProductPrice)

	if err != nil {
		fmt.Println("Error in PriceByDate QueryRow :", err)
		return *obj, false, "PriceByDate failed"
	}
	return *obj, true, "Sucessfully Completed"
}

func (price *PriceStruct) PriceById(obj *models.Price) (models.Price, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in PriceById")
	}
	

	query, _ := Db.Prepare(`SELECT productid,productprice,createdon from "role" where id=$1`)

	err := query.QueryRow(obj.Id).Scan(&obj.ProductId,obj.ProductPrice,obj.Createdon)

	if err != nil {
		fmt.Println("Error in PriceById QueryRow :", err)
		return *obj, false, "PriceById failed"
	}
	return *obj, true, "Sucessfully Completed"
}

func (price *PriceStruct) PriceGetAll() ([]models.Price, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Price GetAll")
	}
	result := []models.Price{}
	priceStruct := models.Price{}

	query, err := Db.Query(`SELECT id,productid,productprice,createdon FROM "price" WHERE productid=$1`)
	if err != nil {
		fmt.Println("Error in Price GetAll Queryrow :", err)
	}

	for query.Next() {
		err := query.Scan(
			&priceStruct.Id,
			&priceStruct.ProductId,
			&priceStruct.ProductPrice,
		)
		if err != nil {
			fmt.Println("Error is founded :", err)
			return result, false, "failed to  Get All Price Data"
		}
		result = append(result, priceStruct)
	}
	return result, true, "sucessfully Completed"
}

func (price *PriceStruct) PriceProductGetAll(obj *models.Price) ([]models.Price, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Price GetAll")
	}
	result := []models.Price{}
	priceStruct := models.Price{}

	query, err := Db.Query(`SELECT id,productid,productprice,createdon FROM "price" WHERE productid=$1`,obj.ProductId)
	if err != nil {
		fmt.Println("Error in Price GetAll Queryrow :", err)
	}

	for query.Next() {
		err := query.Scan(
			&priceStruct.Id,
			&priceStruct.ProductId,
			&priceStruct.ProductPrice,
		)
		if err != nil {
			fmt.Println("Error is founded :", err)
			return result, false, "failed to  Get All Price Data"
		}
		result = append(result, priceStruct)
	}
	return result, true, "sucessfully Completed"
}
