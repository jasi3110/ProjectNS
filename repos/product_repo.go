package repos

import (
	"OnlineShop/models"
	"OnlineShop/repos/masterRepo"
	// "OnlineShop/repos/masterRepo"
	"OnlineShop/utls"
	"fmt"
	"log"
)

type ProductInterface interface {
	ProductCreate(obj *models.Product) (string, bool)
	GetProductById(obj *int64) (models.ProductAll, bool, string)
	ProductUpdate(obj *models.Product) (string, bool)
	ProductGetAll() ([]models.ProductAll, bool, string)
	ProductGetAllByUnit(obj *int64) ([]models.ProductAll, bool, string)
	ProductGetAllByCategory(obj *int64) ([]models.ProductAll, bool, string)
}
type ProductStruct struct {
}

func (product *ProductStruct) ProductCreate(obj *models.Product) (string, bool) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in ProductCreate")
	}
if obj.Image==""{
obj.Image="https://www.shutterstock.com/image-vector/blank-avatar-photo-place-holder-600w-1114445501.jpg"
}
	err := Db.QueryRow(`INSERT into "product"(
		image,
		name,
		category,
		quantity,
		unit,
		price,
		createdon
		)values($1,$2,$3,$4,$5,$6,$7)RETURNING id`,
		obj.Image,
		obj.Name,
		obj.Category,
		obj.Quantity,
		obj.Unit,
		0,
		utls.GetCurrentDate()).Scan(&obj.Id)

	if err != nil {
		fmt.Println("Error in Product Create QueryRow :", err)
		return "Create Product Failed", false
	}
	priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	pricesStruct := models.Price{
		ProductId:    obj.Id,
		Mrp: obj.Mrp,
		Nop: obj.Nop,
	}
	status, descreption, value := priceRepo.CreatePrice(&pricesStruct)

	if !status {
		fmt.Println(descreption)
		return descreption, false
	}
	fmt.Println("not updated")
	pricequery := `UPDATE "product" SET price=$2 WHERE id=$1`
	a, err := Db.Exec(pricequery, &value.ProductId, &value.Id)

	if err != nil {
		fmt.Println("Error in Price Update QueryRow :", a, err)
		return "Update Failed", false
	}
	return "Product Created Sucessfully", true
}

func (product *ProductStruct) ProductUpdate(obj *models.Product) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Product Update")
	}
	priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	pricesStruct := models.Price{
		ProductId:    obj.Id,
		Mrp: obj.Mrp,
		Nop: obj.Nop,
	}
	status, descreption, value := priceRepo.CreatePrice(&pricesStruct)

	if !status {
		fmt.Println(descreption)
		return descreption, false
	}

	pricequery := `UPDATE "product" SET image=$2,name=$3,category=$4,unit=$5,price=$6,quantity=$7 WHERE id=$1`
	a, err := Db.Exec(pricequery, &obj.Id,&obj.Image,&obj.Name,&obj.Category,obj.Unit, &value.Id, &obj.Quantity)
	if err != nil {
		fmt.Println("Error in Price Update QueryRow :", a, err)
		return "Update Failed", false
	}

	if err != nil {
		fmt.Println("Error in Product Update QueryRow :", err)
		return "Update Failed", false
	}
	return "Sucessfully Updated", true
}

func (product *ProductStruct) GetProductById(obj *int64) (models.ProductAll, bool, string) {

	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnected in ProductGetBy ID")
	}
	// catid:=3
	productStruct := models.ProductAll{}
	query, _ := Db.Prepare(`SELECT id,
								   image,
								   name,
								   category,
								   coalesce( (select name from category where id = category) ) as category,
								   quantity,
								   unit,
								   coalesce( (select item from unit where id = unit) ) as unit,
								   price,
								   createdon from "product" where id=$1`)
	err := query.QueryRow(obj).Scan(&productStruct.Id,
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
		fmt.Println("Error in Product GetById QueryRow :", err)
		return productStruct, false, "Failed"
	}
	priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	value,status, descreption := priceRepo.PriceById(&productStruct.Price)

	productStruct.Price = value
	if !status {
		fmt.Println(descreption)
		return productStruct,false,descreption
	}
	
	return productStruct, true, "Get Product Sucessfully Completed"
}

func (product *ProductStruct) ProductGetAll() ([]models.ProductAll, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Product GetAll ")
	}
	result := []models.ProductAll{}
	productStruct := models.Product{}

	query, err := Db.Query(`SELECT id FROM "product" WHERE isdiscount=0 and isdeleted=0`)
	if err != nil {
		log.Println(err)
	}
	for query.Next() {
		err := query.Scan(&productStruct.Id)
		value, status, descreption := product.GetProductById(&productStruct.Id)
		
		if err != nil {
			fmt.Println("Error in Product GetAll QueryRow :", err)
			return result, false, "failed to  Get All Product Data"
		}
		if !status {
			fmt.Println(descreption)
			return result, false,descreption
		}
		result = append(result, value)
	}
	return result, true, "sucessfully Completed"
}

func (product *ProductStruct) ProductGetAllByCategory(obj *int64) ([]models.ProductAll, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Product GetAllBY Category ")
	}
	result := []models.ProductAll{}
	productStruct := models.Product{}

	query, err := Db.Query(`SELECT id FROM "product" WHERE category=$1`,obj)
	if err != nil {
		log.Println(err)
	}
	for query.Next() {
		err := query.Scan(&productStruct.Id)
		value, status, descreption := product.GetProductById(&productStruct.Id)
		
		if err != nil {
			fmt.Println("Error in Product GetAll Category QueryRow :", err)
			return result, false, "failed to  Get All Product Data"
		}
		if !status {
			fmt.Println(descreption)
			return result, false,descreption
		}
		result = append(result, value)
	}
	return result, true, "sucessfully Completed"
}

func (product *ProductStruct) ProductGetAllByUnit(obj *int64) ([]models.ProductAll, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Product GetAllBY Category ")
	}
	result := []models.ProductAll{}
	productStruct := models.Product{}

	query, err := Db.Query(`SELECT id FROM "product" WHERE unit=$1`,obj)
	if err != nil {
		log.Println(err)
	}
	for query.Next() {
		err := query.Scan(&productStruct.Id)
		value, status, descreption := product.GetProductById(&productStruct.Id)
		
		if err != nil {
			fmt.Println("Error in Product GetAll Category QueryRow :", err)
			return result, false, "failed to  Get All Product Data"
		}
		if !status {
			fmt.Println(descreption)
			return result, false,descreption
		}
		result = append(result, value)
	}
	return result, true, "sucessfully Completed"
}
