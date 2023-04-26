package repos

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
	"log"
)

type DiscountInterface interface {
	CreateDiscount(obj *models.RDiscount) (bool, string)
	DiscountById(obj *int64) (models.ProductAll, bool, string)
	DiscountGetAll() ([]models.ProductAll, bool, string)
	DiscountUpdate(obj *models.RDiscount) (bool,string)
}
type 	DiscountStruct struct {
}

func (discount *DiscountStruct) CreateDiscount(obj *models.RDiscount) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Create Discount")
	}
	productRepo:=ProductInterface(&ProductStruct{})
	value,status,descreption:=productRepo.GetProductById(&obj.Id)
	
	if !status{
		fmt.Println(descreption)
			return false,descreption
	}
	query, err := Db.Query(`SELECT productid FROM "discount" WHERE isdeleted=0`)
	if err != nil {
		log.Println("Error in Check product Discount Create   QueryRow :", err)
	}
	userStruct := models.RDiscount{}
	for query.Next() {
		query.Scan(&userStruct.Id)
		if obj.Id == userStruct.Id {
			fmt.Println("This Product already in Discount")
			return false,"This Product already in Discount"
		}
	}
	precentagevalue:=value.Price.Mrp * obj.Percentage/100  
	val:=value.Price.Mrp-precentagevalue
	err = Db.QueryRow(`INSERT INTO "discount" (productid,
											   percentage,
											   mrp,
											   nop,
											   startdate,
											   enddate) 
	values($1,$2,$3,$4,$5,$6)RETURNING id`, value.Id,
											obj.Percentage,
											value.Price.Mrp,
											val,
											utls.GetCurrentDate(),
											obj.Enddate,
											).Scan(&obj.Id)
	

	if err != nil {
		fmt.Println("Error in Create Discount QueryRow :", err)
		return false, " Create Discount Failed "
}

	_,err=Db.Query(`UPDATE "product" SET isdiscount=1 where id=$1`,value.Id)

	if err != nil {
		fmt.Println("Error in Create Discount QueryRow :", err)
		return false, " Create Discount Failed "
	}
	return true, " Create Discount Successfully Compeleted "
}

func (discount *DiscountStruct) DiscountUpdate(obj *models.RDiscount) (bool,string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Discount Product Update")
	}

	err :=Db.QueryRow( `UPDATE "discount" SET percentage = $2 ,enddate=$3 WHERE id=$1 and isdeleted=0`,&obj.Id,&obj.Percentage,&obj.Enddate)
	
	if err != nil {
		fmt.Println("Error in Discount Product Upadte QueryRow :", err) 
		return false, "Update Failed"
	}
	return true, "Sucessfully Updated"
}

func (discount *DiscountStruct) DiscountById(obj *int64) (models.ProductAll, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB Disconnceted in Discount Product GetById")
	}
	productStruct := models.ProductAll{}

	query, err := Db.Prepare(`SELECT productid,percentage,mrp,nop FROM "discount" WHERE productid=$1 and isdeleted=0`)
	if err != nil {
		fmt.Println("Error in discount product Getbyid QueryRow Scan :",err)
		return productStruct, false, "Failed to get discount Product"
	}
		err = query.QueryRow(obj).Scan(&productStruct.Id,&productStruct.Percentage,&productStruct.Price.Mrp,&productStruct.Price.Nop)
		if err != nil {
			fmt.Println("Error in discount product Getbyid QueryRow Scan :", err)
			return productStruct, false, "Failed to get discount Product"
		}
			productRepo:=ProductInterface(&ProductStruct{})
			value,status,descreption:=productRepo.GetProductById(&productStruct.Id)
	
			value.Percentage=productStruct.Percentage
			value.Price.Id=0000
			value.Price.Mrp=productStruct.Price.Mrp
			value.Price.Nop=productStruct.Price.Nop
			if !status{
				fmt.Println(descreption)
					return value,false,descreption
			}		
	
	
	return value, true, "sucessfully Completed"
}

func (discount *DiscountStruct) DiscountGetAll() ([]models.ProductAll, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Discount Product  GetAll")
	}
	result := []models.ProductAll{}
	discountStruct:=models.DiscountProductAll{}

	query, err := Db.Query(`SELECT productid,percentage,mrp,nop FROM "discount"`)
	if err != nil {
		fmt.Println("Error in Discount Product  GetAll QueryRow :", err)
		return result, false, "failed"
	}

	for query.Next() {
		err := query.Scan(
			&discountStruct.Id,
			&discountStruct.Percentage,
			&discountStruct.Price.Mrp,
			&discountStruct.Price.Nop)
		if err != nil {
			fmt.Println("Error in Discount Product  GetAll QueryRow Scan :", err)
			return result, false, "failed"
		}
			productRepo:=ProductInterface(&ProductStruct{})
			value,status,descreption:=productRepo.GetProductById(&discountStruct.Id)
			// fmt.Println(value,status)
			value.Percentage=discountStruct.Percentage
			value.Price.Id=0000
			value.Price.Mrp=discountStruct.Price.Mrp
			value.Price.Nop=discountStruct.Price.Nop
			if !status{
				fmt.Println(descreption)
					return result,false,descreption
			}
			
			
		result = append(result, value )
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "sucessfully Completed"
}
