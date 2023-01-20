package repos

import (
	"OnlineShop/models"
	"OnlineShop/utls"
	"fmt"

	// "github.com/spf13/pflag"
)

type SaleInterface interface {
	CreateSale(Obj models.Invoice) (bool,string)
}

type SaleStruct struct {
}

func (sale *SaleStruct) CreateSale(obj models.Invoice) (bool,string) {

	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB is Disconnceted in CreateSale ")
	}
	//create transaction
	Txn, err := Db.Begin()
	if err != nil {
		fmt.Println("Error in CreateSale Transacation in DB :", err)
		return false,"CREATESALE FAILED"
	}

	//write invoice
	err = Txn.QueryRow(`INSERT into "invoice"(
												billamount,
												customerid,
												createdon,
												createdby)
												values($1,$2,$3,$4)RETURNING id`,
		obj.BillAmount,
		obj.CustomerId,
		obj.CreatedOn,
		obj.CreatedBy,
	).Scan(&obj.Id)

	if err != nil {
		fmt.Println("Error in CreateSale Invoice QueryRow :", err)
		err := Txn.Rollback()
		if err != nil {
			fmt.Println("Error in CreateSale Rollback in Invoice :", err)
		}
		return false,"CREATESALE FAILED"
	}
	//write sales entry data from array
	for _, productItem := range obj.Items {

		err := Txn.QueryRow(`INSERT into "salesEntry"(customerid,
												billid,
												invoiceid,
			     								productid,
												productprice,
												quantity,
												createdon)
												values($1,$2,$3,$4,$5,$6,$7)RETURNING id`,
			obj.CustomerId,
			2121222,
			obj.Id,
			productItem.Id,
			productItem.Price,
			productItem.Quantity,
			obj.CreatedOn,
		).Scan(&productItem.Id)
		if err != nil {
			fmt.Println("Error in CreateSale Saleentry QueryRow :", err)
			err := Txn.Rollback()
			if err != nil {
				fmt.Println("Error in CreateSale Rollback in SaleEntry :", err)
			}
			return false,"CREATESALE FAILED"
		}

		productRepo := ProductInterface(&ProductStruct{})
		value, _, _ := productRepo.GetProductById(&productItem.Id)
		productqty := value.Quantity
		if productqty != 0 {
			//reduce stock quantity from product table
			productqty = productqty - productItem.Quantity
			updateQueryqty, err := Txn.Query(`UPDATE  "product" SET  quantity=$1 WHERE id=$2`, productqty, productItem.Id)

			if err != nil {
				fmt.Println("Error in CreateSale  in Product Update QueryRow  :", err)
				err := Txn.Rollback()
				if err != nil {
					fmt.Println("Error in CreateSale Rollback in Product Update QueryRow :", err)
				}
				return false,"CREATESALE FAILED"
			}
			err = updateQueryqty.Close()
			if err != nil {
				fmt.Println("Error in CreateSale Product Update Close :")
			}
		}

		err = Txn.Commit()
		if err != nil {

			err := Txn.Rollback()
			if err != nil {
				fmt.Println("Error in CreateSale Rollback in Product Update :")
			}
			return false,"CREATESALE FAILED"
		}
	}
	fmt.Println(err)
	return true,"CREATESALE SUCESSFULLY COMPLETED"

}
