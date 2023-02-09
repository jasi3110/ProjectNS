package repos


import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"
)

type DiscountInterface interface {
	CreateDiscount(obj *models.RDiscount) (bool, string)
	// DiscountById(obj *models.Category) (models.Category, bool, string)
	DiscountGetAll() ([]models.ProductAll, bool, string)
	// DiscountUpdate(obj *models.Category) (models.Category, string, bool)
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
	
	precentagevalue:=value.Price.Mrp * obj.Percentage/100  
	val:=value.Price.Mrp-precentagevalue
	err  	:= Db.QueryRow(`INSERT INTO "discount" (productid,
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
											obj.Enddate).Scan(&obj.Id)
	
	if err != nil {
		fmt.Println("Error in Create Discount QueryRow :", err)
		return false, " Create Discount Failed "
}
fmt.Println(obj)
	_,err=Db.Query(`UPDATE "product" SET isdiscount=1 where id=$1`,value.Id)

	if err != nil {
		fmt.Println("Error in Create Discount QueryRow :", err)
		return false, " Create Discount Failed "
	}
	return true, " Create Discount Successfully Compeleted "
}

// func (discount *DiscountStruct) DiscountUpdate(obj *models.Category) (models.Category, string, bool) {
// 	Db, isconnceted := utls.OpenDbConnection()
// 	if !isconnceted {
// 		fmt.Println("DB Disconnceted in Category Update")
// 	}

// 	query:= `UPDATE "category" SET name=$2 WHERE id=$1`
// 	_, err := Db.Exec(query, &obj.Id, &obj.Name)

// 	if err != nil {
// 		fmt.Println("Error in Category Upadte QueryRow :", err) 
// 		return *obj, "Update Failed", false
// 	}
// 	return *obj, "Sucessfully Updated", true
// }

// func (category *CategoryStruct) CategoryById(obj *models.Category) (models.Category, bool, string) {
// 	Db, isconnceted := utls.OpenDbConnection()
// 	if !isconnceted {
// 		fmt.Println("DB Disconnceted in Category GetById")
// 	}
// 	categoryStruct := models.Category{}

// 	query, _ := Db.Prepare(`SELECT id,name from "category" where id=$1`)

// 	err := query.QueryRow(obj.Id).Scan(&categoryStruct.Id, &categoryStruct.Name)

// 	if err != nil {
// 		fmt.Println("Error in Category GetById QueryRow :", err)
// 		return categoryStruct, false, "Error is founded in category get by id"
// 	}
// 	return categoryStruct, true, "category get id successfully"
// }

func (discount *DiscountStruct) DiscountGetAll() ([]models.ProductAll, bool, string) {
	Db, isConnected := utls.CreateDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in Category GetAll")
	}
	result := []models.ProductAll{}
	discountStruct:=models.DiscountProductAll{}

	query, err := Db.Query(`SELECT productid,percentage,mrp,nop FROM "discount"`)
	if err != nil {
		fmt.Println(err)
	}

	for query.Next() {
		err := query.Scan(
			&discountStruct.Id,
			&discountStruct.Percentage,
			&discountStruct.Price.Mrp,
			&discountStruct.Price.Nop)
		if err != nil {
			fmt.Println("Error in Category GetAll QueryRow :", err)
			return result, false, "failed to  Get All Category Data"
		}
			productRepo:=ProductInterface(&ProductStruct{})
			value,status,descreption:=productRepo.GetProductById(&discountStruct.Id)
			fmt.Println(value,status)
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
	return result, true, "sucessfully Completed"
}

