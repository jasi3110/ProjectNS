package masterRepo

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
	"strconv"
)

type ProductImageInterface interface {
	ProductImageCreate(obj *models.ProductImage) (bool, string)
	ProductByImageId(obj *int64) (models.ProductImage, bool, string)
	ProductImageGetall() (bool, string)

	
}
type ProductImageStruct struct {
}

func (image *ProductImageStruct) ProductImageCreate(obj *models.ProductImage) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Create Product Image")
	}
	url:=models.Imageurl(obj.Imageurl)

	err := Db.QueryRow(`INSERT INTO "productimage" (imageurl,productid,createdon)values($1,$2,$3)RETURNING id`,
	url,obj.Productid,utls.GetCurrentDate()).Scan(&obj.Id)
	if err != nil {
		fmt.Println("Error in Create Product Image QueryRow :", err)
		return false, "Failed"
	}
	query:=`UPDATE "product" SET image=$1 WHERE id=$2 AND isdeleted=0`
	_, err =Db.Exec(query,&obj.Id,&obj.Productid)
	
	if err != nil {
		fmt.Println("Error in Category Upadte QueryRow :", err)
		return false,"Update Failed"
	}
	defer func() {
		Db.Close()
	}()
	return true, "Successfully Created"
}

func (image *ProductImageStruct) ProductImageGetall() (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Category GetById")
	}
	obj := models.ProductImage{}

	query, err:= Db.Query(`SELECT id,image from "product"`)
	if err != nil {
		fmt.Println("Error in Category GetById QueryRow :", err)
		return  false, "Failed"
	}
	for query.Next() {
		err = query.Scan(&obj.Id,&obj.Imageurl)
		if err != nil {
			fmt.Println("Error:", err)
			return false,"Failed"
		}

		num, err := strconv.ParseInt(obj.Imageurl, 10, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return false,"Failed"
		}
		value,_,_:=image.ProductByImageId(&num)
		a:=models.Imageurl(value.Imageurl)
		query:=`UPDATE "product" SET image=$2 WHERE id=$1 AND isdeleted=0`
		_, err=Db.Exec(query,&obj.Id,&a)
		
	if err != nil {
		fmt.Println("Error:", err)
		return false,"Failed"
	}
		
	}
	

	defer func() {
		Db.Close()
		query.Close()
	}()
	return true, "successfully Completed"
}


func (image *ProductImageStruct) ProductByImageId(obj *int64) (models.ProductImage, bool, string) {

	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnected in Product Image GetByID")
	}

	productStruct := models.ProductImage{}
	query, err := Db.Prepare(`SELECT id,
								   imageurl,
								   createdon from "productimage" where id=$1`)
	if err != nil {
		fmt.Println("Error in ProductImage GetById QueryRow :", err)
		return productStruct, false, "Failed"
	}
	err = query.QueryRow(obj).Scan(&productStruct.Id,
		&productStruct.Imageurl,
		&productStruct.Createdon)
	if err != nil {
		fmt.Println("Error in Product Image GetById QueryRow Scan :", err)
		return productStruct, false, "Failed"
	}
	
	defer func() {
		Db.Close()
		query.Close()
	}()
	return productStruct, true, "Successfully Completed"
}
 