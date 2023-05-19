package repos

import (
	"OnlineShop/models"
	"OnlineShop/repos/masterRepo"
	// "sync"

	// "sync"

	// "OnlineShop/repos/masterRepo"
	"OnlineShop/utls"
	"fmt"
	"log"
)

// var wg sync.WaitGroup

type ProductInterface interface {
	ProductCreate(obj *models.Product) (string, bool)
	ProductUpdate(obj *models.Product) (string, bool)
	ProductDelete(obj *models.Product) (bool, string)

	GetProductById(obj *int64) (models.ProductAll, bool, string)
	ProductGetAll() ([]models.ProductAll, bool, string)
	ProductSearchBar(obj string) ([]models.ProductAll, bool)

	ProductGetAllByUnit(obj *int64) ([]models.ProductAll, bool, string)
	ProductGetAllByCategory(obj *int64) ([]models.ProductAll, bool, string)
}
type ProductStruct struct {
}

func (product *ProductStruct) ProductCreate(obj *models.Product) (string, bool) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Product Create")
	}
	if obj.Image == "" {
		obj.Image = "https://www.shutterstock.com/image-vector/blank-avatar-photo-place-holder-600w-1114445501.jpg"
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
		ProductId: obj.Id,
		Mrp:       obj.Mrp,
		Nop:       obj.Nop,
	}
	status, descreption, value := priceRepo.CreatePrice(&pricesStruct)

	if !status {
		fmt.Println(descreption)
		return descreption, false
	}
	_, err = Db.Query(`UPDATE "product" SET price=$2 WHERE id=$1 and isdeleted=0`, &value.ProductId, &value.Id)

	if err != nil {
		fmt.Println("Error in Price Update QueryRow :", err)
		return "Update Failed", false
	}
	defer func() {
		Db.Close()
	}()
	return "Product Created Successfully", true
} 

func (product *ProductStruct) ProductUpdate(obj *models.Product) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Product Update")
	}
	priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	pricesStruct := models.Price{
		ProductId: obj.Id,
		Mrp:       obj.Mrp,
		Nop:       obj.Nop,
	}
	status, descreption, value := priceRepo.CreatePrice(&pricesStruct)

	if !status {
		fmt.Println("Error in product update price create QueryRow :", descreption)
		return descreption, false
	}
	query := `UPDATE "product" SET image=$2,name=$3,category=$4,unit=$5,price=$6,quantity=$7 WHERE id=$1 and isdeleted=0`

	_, err := Db.Exec(query, &obj.Id, &obj.Image, &obj.Name, &obj.Category, obj.Unit, &value.Id, &obj.Quantity)

	if err != nil {
		fmt.Println("Error in Price Update QueryRow :", err)
		return "Update Failed", false
	}

	if err != nil {
		fmt.Println("Error in Product Update QueryRow :", err)
		return "Update Failed", false
	}
	defer func() {
		Db.Close()
	}()
	return "Successfully Updated", true
}

func (product *ProductStruct) ProductDelete(obj *models.Product) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in Product Delete ")
	}
	query := `UPDATE "product" SET isdeleted=1 WHERE id=$1 and isdeleted=0`
	_, err := Db.Exec(query, &obj.Id)

	if err != nil {
		fmt.Println("Error in product Delete QueryRow :", err)
		return false, "Product Delete Failed"
	}
	defer Db.Close()
	return true, "Product Deleted Successfully Completed"
}

func (product *ProductStruct) GetProductById(obj *int64) (models.ProductAll, bool, string) {

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
								   createdon from "product" where id=$1 and isdeleted=0`)
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
	// var productchannel chan models.ProductAll
	for query.Next() {
		err := query.Scan(&productStruct.Id)
		value, _, _ :=product.GetProductById(&productStruct.Id)
		if err != nil {
			fmt.Println("Error in Product GetAll QueryRow :", err)
			return result, false, "failed to  Get All Product Data"
		}
		result = append(result, value)
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "successfully Completed"
}

func (product *ProductStruct) ProductGetAllByCategory(obj *int64) ([]models.ProductAll, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Product GetAllBY Category ")
	}
	result := []models.ProductAll{}
	productStruct := models.Product{}

	query, err := Db.Query(`SELECT id FROM "product" WHERE category=$1`, obj)
	if err != nil {
		fmt.Println("Error in Product GetAll By Category  QueryRow :", err)
		return result, false, "failed"
	}
	// var productchannel chan models.ProductAll
	for query.Next() {
		err := query.Scan(&productStruct.Id)
		value, _,_:= product.GetProductById(&productStruct.Id)

		if err != nil {
			fmt.Println("Error in Product GetAll Category QueryRow :", err)
			return result, false, "failed"
		}

		result = append(result, value)
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "successfully Completed"
}

func (product *ProductStruct) ProductGetAllByUnit(obj *int64) ([]models.ProductAll, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Product GetAllBY Unit ")
	}
	result := []models.ProductAll{}
	productStruct := models.Product{}

	query, err := Db.Query(`SELECT id FROM "product" WHERE unit=$1`, obj)
	if err != nil {
		fmt.Println("Error in Product GetAll By Unit QueryRow :", err)
		return result, false, "failed"
	}
	// var Result chan models.ProductAll
	for query.Next() {
		err := query.Scan(&productStruct.Id)
		value, status, descreption := product.GetProductById(&productStruct.Id)

		if err != nil {
			fmt.Println("Error in Product GetAll By Unit  QueryRow Scan:", err)
			return result, false, "failed"
		}
		if !status {
			fmt.Println(descreption)
			return result, false, descreption
		}
		result = append(result, value)
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "successfully Completed"
}

func (product *ProductStruct) ProductSearchBar(obj string) ([]models.ProductAll, bool) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Product SearchBar")
	}
	var result []models.ProductAll
	productStruct:=models.ProductAll{}
	query, err := Db.Query(`SELECT id,
	image,
	name,
	category,
	coalesce( (select name from category where id = category) ) as category,
	quantity,
	unit,
	coalesce( (select item from unit where id = unit) ) as unit,
	price,
	createdon from "product" where LOWER(name) like $1`, "%"+obj+"%")
	fmt.Println("",obj)
	for query.Next() {
	err := query.Scan(&productStruct.Id,
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
		fmt.Println("Error in Product SearchBar QueryRow :", err)
		return result, false
	}
	priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	value, status, descreption := priceRepo.PriceById(&productStruct.Price)

	productStruct.Price = value
	if !status {
		fmt.Println("Error in Product GetbyId price ById QueryRow :", descreption)
		return result, false
	}
		// fmt.Println("",result)
		result = append(result, productStruct)
	}
	if err != nil {
		return result, false
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true
}
