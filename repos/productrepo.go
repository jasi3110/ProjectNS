package repos

import (
	"OnlineShop/models"
	"OnlineShop/repos/masterRepo"
	"OnlineShop/utls"
	"log"
)

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
		log.Panic("DB Disconnceted in Product Create")
	}

	priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	pricesStruct := models.Price{
		ProductId: obj.Id,
		Mrp:       obj.Mrp,
		Nop:       obj.Nop,
	}

	status, _, value := priceRepo.CreatePrice(&pricesStruct)
	if !status {
		log.Panic("Error in Create Product Price Create")
		return "Somethong Went Wrong", false
	}

	err := Db.QueryRow(`INSERT into "product"(image,name,category,quantity,unit,price,createdon
		)values($1,$2,$3,$4,$5,$6,$7)RETURNING id`,
		obj.Image, obj.Name, obj.Category, obj.Quantity, obj.Unit, value.Id,
		utls.GetCurrentDate()).Scan(&obj.Id)

	if err != nil {
		log.Panic("Error in Product Create QueryRow :", err)
		return "Create Product Failed", false
	}

	defer func() {
		Db.Close()
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	return "Product Created Successfully", true
}

func (product *ProductStruct) ProductUpdate(obj *models.Product) (string, bool) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		log.Panic("DB Disconnceted in Product Update")
	}

	priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
	pricesStruct := models.Price{
		ProductId: obj.Id,
		Mrp:       obj.Mrp,
		Nop:       obj.Nop,
	}

	status, _, value := priceRepo.CreatePrice(&pricesStruct)
	if !status {
		log.Panic("Error in product update price create")
		return "Something Went Wrong", false
	}

	query := `UPDATE "product" SET image=$2,name=$3,category=$4,unit=$5,price=$6,quantity=$7 WHERE id=$1 and isdeleted=0`
	_, err := Db.Exec(query, &obj.Id, &obj.Image, &obj.Name, &obj.Category, obj.Unit, &value.Id, &obj.Quantity)

	if err != nil {
		log.Panic("Error in Product Update QueryRow :", err)
		return "Update Failed", false
	}

	defer func() {
		Db.Close()

		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	return "Successfully Updated", true
}

func (product *ProductStruct) ProductDelete(obj *models.Product) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		log.Panic("DB disconnceted in Product Delete ")
	}

	query := `UPDATE "product" SET isdeleted=1 WHERE id=$1 and isdeleted=0`
	_, err := Db.Exec(query, &obj.Id)

	if err != nil {
		log.Panic("Error in product Delete QueryRow :", err)
		return false, "Product Delete Failed"
	}

	defer func() {
		Db.Close()

		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	return true, "Product Deleted Successfully"
}

func (product *ProductStruct) GetProductById(obj *int64) (models.ProductAll, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		log.Panic("DB Disconnected in Product GetById")
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
								   coalesce( (select mrp from price where id = price) ) as price,
								   coalesce( (select nop from price where id = price) ) as price,
								   createdon from "product" where id=$1 and isdeleted=0`)
	if err != nil {
		log.Panic("Error in Product GetById QueryRow :", err)
		return productStruct, false, "Something Went Wrong"
	}

	err = query.QueryRow(obj).Scan(
		&productStruct.Id,
		&productStruct.Image,
		&productStruct.Name,
		&productStruct.Category.Id,
		&productStruct.Category.Name,
		&productStruct.Quantity,
		&productStruct.Unit.Id,
		&productStruct.Unit.Item,
		&productStruct.Price.Id,
		&productStruct.Price.Mrp,
		&productStruct.Price.Nop,
		&productStruct.CreatedOn)

	if err != nil {
		log.Panic("Error in Product GetById QueryRow Scan :", err)
		return productStruct, false, "Something Went Wrong"
	}

	basicURL := "https://drive.google.com/uc?export=view&id="
	productStruct.Image = basicURL + productStruct.Image
	percentage := 100 - ((float64(productStruct.Price.Nop) / float64(productStruct.Price.Mrp)) * 100)
	productStruct.Price.Percentage = int64(percentage)
	defer func() {
		Db.Close()
		query.Close()
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()
	return productStruct, true, "Successfully Completed"
}

func (product *ProductStruct) GetProductHomePage(obj *int64) (models.Product, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		log.Panic("DB Disconnected in Get Product HomePage")
	}

	productStruct := models.Product{}
	query, err := Db.Prepare(`SELECT id,
								   imageid,
								   name,
								   price,
								   coalesce( (select mrp,nop from price where id = price) ) as price,
								   createdon from "product" where id=$1 and isdeleted=0`)

	if err != nil {
		log.Panic("Error in Product Get product HomePage QueryRow :", err)
		return productStruct, false, "Something Went Wrong"
	}

	err = query.QueryRow(obj).Scan(
		&productStruct.Id,
		&productStruct.Image,
		&productStruct.Name,
		&productStruct.Price,
		&productStruct.CreatedOn)

	if err != nil {
		log.Panic("Error in Product Get Product HomePage QueryRow Scan : ", err)
		return productStruct, false, "Something Went Wrong"
	}
	
	defer func() {
		Db.Close()
		query.Close()
		
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}

	}()

	return productStruct, true, "Successfully Completed"
}

func (product *ProductStruct) ProductGetAll() ([]models.ProductAll, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		log.Panic("DB Disconnceted in Product GetAll ")
	}

	result := []models.ProductAll{}
	productStruct := models.Product{}

	query, err := Db.Query(`SELECT id FROM "product" WHERE isdiscount=0 and isdeleted=0`)
	if err != nil {
		log.Panic("Error in Product GetAll QueryRow :",err)
	}
	
	for query.Next() {
		err := query.Scan(&productStruct.Id)
		if err != nil {
			log.Panic("Error in Product GetAll QueryRow Scan :", err)
			return result, false, "Something Went Wrong"
		}

		value, status, _ := product.GetProductById(&productStruct.Id)
		if !status {
			log.Panic("Error in Product GetAll GetProductById ")
			return result, false, "Something Went Wrong"
		}

		result = append(result, value)
	}

	defer func() {
		Db.Close()
		query.Close()

		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
		
	}()
	return result, true, "successfully Completed"
}

func (product *ProductStruct) ProductGetAllByCategory(obj *int64) ([]models.ProductAll, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		log.Panic("DB Disconnceted in Product GetAllBY Category ")
	}

	result := []models.ProductAll{}
	productStruct := models.Product{}

	query, err := Db.Query(`SELECT id FROM "product" WHERE category=$1`, obj)
	if err != nil {
		log.Panic("Error in Product GetAll By Category  QueryRow :", err)
		return result, false, "Something Went Wrong"
	}
	
	for query.Next() {
		err := query.Scan(&productStruct.Id)
		if err != nil {
			log.Panic("Error in Product GetAll Category QueryRow Scan:", err)
			return result, false, "failed"
		}

		value, status, _ := product.GetProductById(&productStruct.Id)

		if !status {
			log.Panic("Error in Product GetAll Category Getting Product  :")
			return result, false, "Something Went Wrong"
		}

		result = append(result, value)
	}
	defer func() {
		Db.Close()
		query.Close()

		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}

	}()

	return result, true, "successfully Completed"
}

func (product *ProductStruct) ProductGetAllByUnit(obj *int64) ([]models.ProductAll, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		log.Panic("DB Disconnceted in Product GetAllBY Unit ")
	}

	result := []models.ProductAll{}
	productStruct := models.Product{}

	query, err := Db.Query(`SELECT id FROM "product" WHERE unit=$1`, obj)
	if err != nil {
		log.Panic("Error in Product GetAll By Unit QueryRow :", err)
		return result, false, "Something Went Wrong"
	}
	
	for query.Next() {
		err := query.Scan(&productStruct.Id)
		if err != nil {
			log.Panic("Error in Product GetAll By Unit  QueryRow Scan:", err)
			return result, false, "failed"
		}

		value, status, _:= product.GetProductById(&productStruct.Id)
		if !status {
			log.Panic("Error in Product GetAll By Unit Getting Product")
			return result, false, "Something Went Wrong"
		}

		result = append(result, value)
	}

	defer func() {
		Db.Close()
		query.Close()
		
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}

	}()
	return result, true, "successfully Completed"
}

func (product *ProductStruct) ProductSearchBar(obj string) ([]models.ProductAll, bool) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		log.Panic("DB Disconnceted in Product SearchBar")
	}
	var result []models.ProductAll
	productStruct := models.ProductAll{}

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
	
	if err != nil {
		log.Panic("Error in Product SearchBar QueryRow : ", err)
		return result, false
	}

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
			log.Panic("Error in Product SearchBar QueryRow Scan:", err)
			return result, false
		}

		basicURL := "https://drive.google.com/uc?export=view&id="
		productStruct.Image = basicURL + productStruct.Image

		priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
		value, status, _ := priceRepo.PriceById(&productStruct.Price)

		productStruct.Price = value

		if !status {
			log.Panic("Error in Product GetbyId Geeting priceById ")
			return result, false
		}

		result = append(result, productStruct)
	}
	
	defer func() {
		Db.Close()
		query.Close()

		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}

	}()
	return result, true
}
