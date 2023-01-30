package repos

import (
	"OnlineShop/models"
	"OnlineShop/repos/masterRepo"
	"OnlineShop/utls"
	"fmt"
	"log"
)

type ProductInterface interface {
	ProductCreate(obj *models.Product) (string, bool)
	GetProductById(obj *int64) (models.Product, bool, string)
	ProductUpdate(obj *models.Product) (string, bool)
	ProductGetAll() ([]models.Product, bool, string)
}
type ProductStruct struct {
}

func (product *ProductStruct) ProductCreate(obj *models.Product) (string, bool) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in ProductCreate")
	}

	err := Db.QueryRow(`INSERT into "product"(
		name,
		category,
		quantity,
		unit,
		price,
		createdon
		)values($1,$2,$3,$4,$5,$6)RETURNING id`,
		obj.Name,
		obj.Category,
		obj.Quantity,
		obj.Unit,
		obj.Price,
		utls.GetCurrentDate()).Scan(&obj.Id)

	if err != nil {
		fmt.Println("Error in Product Create QueryRow :", err)
		return "Create Product Failed", false
	}
	priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	pricesStruct := models.Price{
		ProductId:    obj.Id,
		ProductPrice: obj.Price,
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
		ProductPrice: obj.Price,
	}
	status, descreption, value := priceRepo.CreatePrice(&pricesStruct)

	if !status {
		fmt.Println(descreption)
		return descreption, false
	}
	fmt.Println("not updated")
	pricequery := `UPDATE "product" SET price=$2,quantity=$3 WHERE id=$1`
	a, err := Db.Exec(pricequery, &obj.Id, &value.Id, &obj.Quantity)

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

func (product *ProductStruct) GetProductById(obj *int64) (models.Product, bool, string) {

	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnected in ProductGetBy ID")
	}
	productStruct := models.Product{}
	query, _ := Db.Prepare(`SELECT id,
								   name,
								   coalesce(select name from category where id = category)as category,
								   quantity,
								   unit,
								   price,
								   createdon from "product" where id=$1`)
	err := query.QueryRow(obj).Scan(&productStruct.Id,
		&productStruct.Name,
		&productStruct.Category,
		&productStruct.Quantity,
		&productStruct.Unit,
		&productStruct.Price,
		&productStruct.CreatedOn)
	if err != nil {
		fmt.Println("Error in Product GetById QueryRow :", err)
		return productStruct, false, "Failed"
	}
	// categoryrepo := masterRepo.CategoryInterface(&masterRepo.CategoryStruct{})
	// unitrepo := masterRepo.UnitInterface(&masterRepo.UnitStruct{})
	// pricerepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})

	// category, statuscat, descreptioncat := categoryrepo.CategoryById(&productStruct.Category)
	// unit, statusunit, descreptionunit := unitrepo.UnityById(&productStruct.Unit)
	// price, statusprice, descreptionprice := pricerepo.PriceById(&productStruct.Price)
	// if !statuscat {
	// 	fmt.Println(descreptioncat)
	// 	return productStruct, false, descreptioncat
	// }
	// if !statusunit {
	// 	fmt.Println(descreptionunit)
	// 	return productStruct, false, descreptionunit
	// }
	// if !statusprice {
	// 	fmt.Println(descreptionunit)
	// 	return productStruct, false, descreptionprice

	// }
	// result := models.ProductAll{
	// 	Id:        productStruct.Id,
	// 	Name:      productStruct.Name,
	// 	Category:  category,
	// 	Quantity:  productStruct.Quantity,
	// 	Unit:      unit,
	// 	Price:     price,
	// 	CreatedOn: productStruct.CreatedOn,
	// }
	return productStruct, true, "Get Product Sucessfully Completed"
}

func (product *ProductStruct) ProductGetAll() ([]models.Product, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Product GetAll ")
	}
	result := []models.Product{}
	productStruct := models.Product{}

	query, err := Db.Query(`SELECT id FROM "product"`)
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
