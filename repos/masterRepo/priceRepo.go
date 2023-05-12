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

	err := Db.QueryRow(`INSERT INTO "price" (
		productid,
		mrp,
		nop,
		createdon)values($1,$2,$3,$4)RETURNING id `,
		obj.ProductId, obj.Mrp,obj.Nop, utls.GetCurrentDate()).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Createprice QueryRow:", err)
		return false, "Failed", *obj
	}
	defer Db.Close()
	return true, "price  Successfully Created", *obj
}

func (price *PriceStruct) PriceUpdate(obj *models.Price) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in PriceUpdate")
	}

	query := `UPDATE "price" SET mrp=$2,nop=$3,productid=$4 WHERE id=$1`
	_, err := Db.Exec(query, &obj.Id, &obj.Mrp,&obj.Nop, &obj.ProductId)

	if err != nil {
		fmt.Println("Error in PriceUpdate QueryRow :", err)
		return "Failed", false
	}
	defer Db.Close()
	return "Successfully Updated", true
}

func (price *PriceStruct) PriceByDate(obj *models.Price) (models.Price, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in PriceByDate")
	}
	priceStruct := models.Price{}
	query, err := Db.Prepare(`SELECT id,productid,mrp,nop,createdon from "price" where productid=$1 and createdon=$2`)
	if err != nil {
		fmt.Println("Error in PriceByDate QueryRow :", err)
		return priceStruct, false, "Failed"
	}
	
	err = query.QueryRow(obj.ProductId, obj.Createdon).Scan(
		&priceStruct.Id,
		&priceStruct.ProductId,
		&priceStruct.Mrp,
		&priceStruct.Nop,
		&priceStruct.Createdon)

	if err != nil {
		fmt.Println("Error in PriceByDate QueryRow Scan :", err)
		return priceStruct, false, "Failed"
	}
	defer Db.Close()
	return priceStruct, true, "Successfully Completed"
}

func (price *PriceStruct) PriceById(obj *models.Price) (models.Price, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in PriceById")
	}
	priceStruct := models.Price{}

	query, err := Db.Prepare(`SELECT id,productid,mrp,nop,createdon from "price" where id=$1`)

	if err != nil {
		fmt.Println("Error in PriceById QueryRow :", err)
		return priceStruct, false, "Failed"
	}
	err = query.QueryRow(obj.Id).Scan(&priceStruct.Id, &priceStruct.ProductId, &priceStruct.Mrp,&priceStruct.Nop, &priceStruct.Createdon)
	percentage:= 100 - ((float64(priceStruct.Nop) / float64(priceStruct.Mrp)) * 100) 
	
	priceStruct.Percentage=int64(percentage)
	if err != nil {
		fmt.Println("Error in PriceById QueryRow Scan:", err)
		return priceStruct, false, "Failed"
	}
	defer Db.Close()
	return priceStruct, true, "Successfully Completed"
}

func (price *PriceStruct) PriceGetAll() ([]models.Price, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Price GetAll")
	}
	result := []models.Price{}
	priceStruct := models.Price{}

	query, err := Db.Query(`SELECT id,productid,mrp,nop,createdon FROM "price"`)
	if err != nil {
		fmt.Println("Error in Price GetAll Queryrow :", err)
		return result, false, "Failed"
	}

	for query.Next() {
		err := query.Scan(
			&priceStruct.Id,
			&priceStruct.ProductId,
			&priceStruct.Mrp,
			&priceStruct.Nop,
			&priceStruct.Createdon,
		)
		if err != nil {
			fmt.Println("Error is founded :", err)
			return result, false, "Failed"
		}
		result = append(result, priceStruct)
	}
	defer Db.Close()
	return result, true, "successfully Completed"
}

func (price *PriceStruct) PriceProductGetAll(obj *models.Price) ([]models.Price, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Price GetAll")
	}
	result := []models.Price{}
	priceStruct := models.Price{}

	query, err := Db.Query(`SELECT id,productid,mrp,nop,createdon FROM "price" WHERE productid=$1`, obj.ProductId)
	if err != nil {
		fmt.Println("Error in Price GetAll Queryrow :", err)
	}

	for query.Next() {
		err := query.Scan(
			&priceStruct.Id,
			&priceStruct.ProductId,
			&priceStruct.Mrp,
			&priceStruct.Nop,
			&priceStruct.Createdon,
		)
		if err != nil {
			fmt.Println("Error in Price GetAll Queryrow Scan :", err)
			return result, false, "Failed"
		}
		result = append(result, priceStruct)
	}
	defer Db.Close()
	return result, true, "successfully Completed"
}
