package repos

import (
	"OnlineShop/models"
	"OnlineShop/repos/masterRepo"
	"OnlineShop/utls"
	"fmt"
	"log"
)

type DiscountInterface interface {
	CreateDiscount(obj *models.RDiscount) (bool, string)
	DiscountProductById(obj *int64) (models.ProductAll, bool, string)
	DiscountGetAll() ([]models.ProductAll, bool, string)
	DiscountUpdate(obj *models.RDiscount) (bool, string)
}
type DiscountStruct struct {
}

func (discount *DiscountStruct) CreateDiscount(obj *models.RDiscount) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Create Discount")
	}
	// txn,err:=Db.Begin()
	// CHECKING PRODUCT IS ALREADY IN DISCOUNT PRODUCT
	query, err := Db.Query(`SELECT productid FROM "discount" WHERE isdeleted=0`)
	if err != nil {
		log.Println("Error in Check product Discount Create   QueryRow :", err)
	}

	userStruct := models.RDiscount{}
	for query.Next() {
		query.Scan(&userStruct.Id)
		if obj.Id == userStruct.Id {
			fmt.Println("This Product already in Discount")
			return false, "This Product already in Discount"
		}
	}
	// TAKING PRODUCT PRICE ID VALUE IN PRODUCT GET BY ID
	fmt.Println("in product getbyid")
	productRepo := ProductInterface(&ProductStruct{})
	productValue, productStatus, desc := productRepo.GetProductById(&obj.Id)
	if !productStatus {
		fmt.Println(desc)
		return false, desc
	}

	precentagevalue := productValue.Price.Mrp * obj.Percentage / 100
	nop := productValue.Price.Mrp - precentagevalue
	priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	pricesStruct := models.Price{
		ProductId: obj.Id,
		Mrp:       productValue.Price.Mrp,
		Nop:       nop,
	}
	
	status, descreption, priceValue := priceRepo.CreatePrice(&pricesStruct)

	if !status {
		fmt.Println(descreption)
		return false, "Failed"
	}
	err = Db.QueryRow(`INSERT INTO "discount" (productid,
											   percentage,
											   priceid,
											   oldpriceid,
											   startdate,
											   enddate)
	values($1,$2,$3,$4,$5,$6)RETURNING id`,
		obj.Id,
		obj.Percentage,
		priceValue.Id,
		productValue.Price.Id,
		utls.GetCurrentDate(),
		obj.Enddate,
	).Scan(&obj.Id)

	if err != nil {
		fmt.Println("rollback in insert into query")
		fmt.Println("Error in Create Discount QueryRow :", err)
		return false, " Create Discount Failed "
	}
	
	err= Db.QueryRow( `UPDATE "product" SET price=$2 ,isdiscount=1 WHERE id=$1 and isdeleted=0 RETURNING id`,
	productValue.Id,&priceValue.Id).Scan(&priceValue.Id)

	if err != nil {
		fmt.Println("Error in Price Update QueryRow in cREATE Discount Product  :", err)
		return false, " Failed "
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return true, " Create Discount Product Successfully Compeleted "
}

func (discount *DiscountStruct) DiscountUpdate(obj *models.RDiscount) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Discount Product Update")
	}
	Txn, _ := Db.Begin()
	// TAKING PRODUCT PRICE ID VALUE IN PRODUCT GET BY ID

	productValue, _, _ := discount.DiscountProductById(&obj.Id)

	precentagevalue := productValue.Price.Mrp * obj.Percentage / 100
	nop := productValue.Price.Mrp - precentagevalue
	priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	pricesStruct := models.Price{
		ProductId: obj.Id,
		Mrp:       productValue.Price.Mrp,
		Nop:       nop,
	}
	status, descreption, priceValue := priceRepo.CreatePrice(&pricesStruct)

	query := `UPDATE "discount" SET percentage = $2,priceid=$3,enddate=$4 WHERE id=$1 and isdeleted=0`
	_, err := Txn.Exec(query, &obj.Id, &obj.Percentage, &priceValue.Id, &obj.Enddate)

	if err != nil {
		err := Txn.Rollback()
		if err != nil {
			fmt.Println("Error in Update Discount Rollback in Discount :", err)
		}
		fmt.Println("Error in Discount Product Upadte QueryRow :", err)
		return false, "Update Failed"
	}
	err = Txn.QueryRow(`UPDATE "product" SET price =$2 , isdiscount=1 where id=$1 RETURNING id `, productValue.Id, priceValue.Id).Scan(&obj.Id)
	if err != nil {
		err := Txn.Rollback()
		if err != nil {
			fmt.Println("Error in Update Product Rollback in Create Discount :", err)
		}
		fmt.Println("Error in Update Discount QueryRow :", err)
		return false, " Update Discount Failed "
	}

	if !status {
		fmt.Println(descreption)
		return false, descreption
	}
	err = Txn.Commit()
		fmt.Println("transcration commited ")
		if err != nil {
			fmt.Println("transcration commit Failed")
			err := Txn.Rollback()
			if err != nil {
				fmt.Println("Error in CreateSale Rollback in Product Update :", err)
			}
			return false, " Update Discount Failed "
		}

	defer func() {
		Db.Close()
	}()
	return true, "Sucessfully Updated"
}

func (discount *DiscountStruct) DiscountProductById(obj *int64) (models.ProductAll, bool, string) {

	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnected in Product GetByID")
	}

	productStruct := models.ProductAll{}
	query, err := Db.Prepare(`SELECT id,
								   image,
								   name,
								   category,
								   coalesce( (select name from category where id = category) ) as category,
								   quantity,
								   unit,
								   coalesce( (select item from unit where id = unit) ) as unit,
								   price,
								   createdon from "product" where id=$1 and isdeleted=0 and isdiscount=1`)
	if err != nil {
		fmt.Println("Error in Product GetById QueryRow :", err)
		return productStruct, false, "Failed"
	}
	err = query.QueryRow(obj).Scan(&productStruct.Id,
		&productStruct.Image,
		&productStruct.Name,
		&productStruct.Category.Id,
		&productStruct.Category.Name,
		&productStruct.Quantity,
		&productStruct.Unit.Id,
		&productStruct.Unit.Item,
		&productStruct.Price.Id,
		&productStruct.CreatedOn)
	if err != nil {
		fmt.Println("Error in Product GetById QueryRow Scan :", err)
		return productStruct, false, "Failed"
	}
	priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	value, status, descreption := priceRepo.PriceById(&productStruct.Price)

	productStruct.Price = value
	if !status {
		fmt.Println("Error in Product GetbyId price ById QueryRow :", descreption)
		return productStruct, false, descreption
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return productStruct, true, "Successfully Completed"
}

func (discount *DiscountStruct) DiscountGetAll() ([]models.ProductAll, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Discount Product  GetAll")
	}
	result := []models.ProductAll{}
	discountStruct := models.DiscountProductAll{}

	query, err := Db.Query(`SELECT productid FROM "discount"`)
	if err != nil {
		fmt.Println("Error in Discount Product  GetAll QueryRow :", err)
		return result, false, "failed"
	}

	for query.Next() {
		err := query.Scan(&discountStruct.Id)
		if err != nil {
			fmt.Println("Error in Discount Product  GetAll QueryRow Scan :", err)
			return result, false, "failed"
		}
		// var productchannel chan models.ProductAll
		productRepo := ProductInterface(&ProductStruct{})
		value, _, _ := productRepo.GetProductById(&discountStruct.Id)
		// fmt.Println(value,status)
		result = append(result, value)
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "sucessfully Completed"
}
