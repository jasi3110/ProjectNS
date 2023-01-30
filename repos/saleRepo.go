package repos

import (
	"OnlineShop/models"
	"OnlineShop/repos/masterRepo"
	"OnlineShop/utls"
	"fmt"
	"log"
	// "github.com/spf13/pflag"
)

type SaleInterface interface {
	CreateSale(Obj *models.Invoice) (bool, string, models.Invoice)
	InvoiceGetall() ([]models.Invoice, bool)
	SaleGetByBillid(obj *int64) (models.InvoiceBillById, bool, string)
	GetUserReportByDateRange(obj *models.GetUserReportByDateRange)  ([]models.InvoiceBillById, bool, string)
}

type SaleStruct struct {
}

func (sale *SaleStruct) CreateSale(obj *models.Invoice) (bool, string, models.Invoice) {

	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB is Disconnceted in CreateSale ")
	}
	a := 1
	//create transaction
	Txn, err := Db.Begin()
	if err != nil {
		fmt.Println("Error in CreateSale Transacation in DB :", err)
		return false, "CREATESALE FAILED", *obj
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
		utls.GetCurrentDate(),
		obj.CreatedBy,
	).Scan(&obj.Id)

	log.Println("invoice created:", obj.Id)

	if err != nil {
		fmt.Println("Error in CreateSale Invoice QueryRow :", err)
		err := Txn.Rollback()
		if err != nil {
			fmt.Println("Error in CreateSale Rollback in Invoice :", err)
		}
		return false, "CREATESALE FAILED", *obj
	}
	//write sales entry data from array

	for _, productItem := range obj.Items {

		err := Txn.QueryRow(`INSERT into "saleentry"(customerid,
												billid,
												invoiceid,
			     								productid,
												productprice,
												quantity,
												createdon,
												createdby)
												values($1,$2,$3,$4,$5,$6,$7,$8)RETURNING id`,
			obj.CustomerId,
			100001+obj.Id,
			obj.Id,
			productItem.Id,
			productItem.Price,
			productItem.Quantity,
			obj.CreatedOn,
			obj.CreatedBy,
		).Scan(&a)
		fmt.Println("", a)
		if err != nil {
			fmt.Println("Error in CreateSale SaleEntry QueryRow :", err)
			err := Txn.Rollback()
			if err != nil {
				fmt.Println("Error in CreateSale Rollback in SaleEntry :", err)
			}
			return false, "CREATESALE FAILED", *obj
		}

		log.Println("Sales Entry Added", productItem.Id)
		productRepo := ProductInterface(&ProductStruct{})
		value, status, _ := productRepo.GetProductById(&productItem.Id)
		productqty := value.Quantity
		if productqty != 0 || status {
			//reduce stock quantity from product table
			productqty = productqty - productItem.Quantity
			updateQueryqty, err := Txn.Query(`UPDATE  "product" SET  quantity=$1 WHERE id=$2`, productqty, productItem.Id)

			if err != nil {
				fmt.Println("Error in CreateSale in Product Update QueryRow  :", err)
				err := Txn.Rollback()
				if err != nil {
					fmt.Println("Error in CreateSale Rollback in Product Update QueryRow :", err)
				}
				return false, "CREATESALE FAILED", *obj
			}
			err = updateQueryqty.Close()
			if err != nil {
				fmt.Println("Error in CreateSale Product Update Close :", err)
			}
		}

		err = Txn.Commit()
		fmt.Println("transcration commited ")
		if err != nil {
			fmt.Println("transcration commit Failed")
			err := Txn.Rollback()
			if err != nil {
				fmt.Println("Error in CreateSale Rollback in Product Update :", err)
			}
			return false, "CREATESALE FAILED", *obj
		}
	}
	fmt.Println("last error :", err)
	return true, "CREATESALE SUCESSFULLY COMPLETED", *obj
}

func (sale *SaleStruct) InvoiceGetall() ([]models.Invoice, bool) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in invoice Getall")
	}
	invoiceStruct := models.Invoice{}
	result := []models.Invoice{}

	query, err := Db.Query(`SELECT id,billamount,customerid,createdon,createdby FROM "invoice"`)
	if err != nil {
		log.Println(err)
	}

	for query.Next() {
		err := query.Scan(
			&invoiceStruct.Id,
			&invoiceStruct.BillAmount,
			&invoiceStruct.CustomerId,
			&invoiceStruct.CreatedOn,
			&invoiceStruct.CreatedBy,
		)
		if err != nil {
			fmt.Println("Error in User GetAll QueryRow :", err)
			return result, false
		}
		result = append(result, invoiceStruct)
	}

	return result, true
}

func (sale *SaleStruct) SaleGetByBillid(obj *int64) (models.InvoiceBillById, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()

	result := models.InvoiceBillById{}
	productStruct := models.ProductAll{}

	if !isconnceted {
		fmt.Println("DB is Disconnceted in SaleGetByBillid")
		return result, false, "Failed"
	}

	query, err := Db.Query(`SELECT          id,
											customerid,
											billid,
											productid,
											productprice,
											quantity,
											createdon,
											createdby
												FROM "saleentry"
												WHERE billid = $1`, obj)
	if err != nil {
		fmt.Println("Error in SaleGetByBillid QueryRow :", err)
		return result, false, "Failed"
	}
	for query.Next() {
		err := query.Scan(&result.Id,
			&result.CustomerId,
			&result.BillId,
			&productStruct.Id,
			&productStruct.Price.Id,
			&productStruct.Quantity,
			&result.CreatedOn,
			&result.CreatedBy,
		)
		if err != nil {
			fmt.Println("Error in SaleGetByBillid QueryRow :", err)
			return result, false, "Failed"
		}
		productRepo := ProductInterface(&ProductStruct{})
		value, status, descreption := productRepo.GetProductById(&productStruct.Id)
		if !status {
			fmt.Println(descreption)
			return result, false, "Failed"
		}
		priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
		valueprice, statusprice, descreptionprice := priceRepo.PriceById(&productStruct.Price)
		if !statusprice {
			fmt.Println(descreptionprice)
			return result, false, "Failed"
		}
		value.Price = valueprice.ProductPrice
		value.Quantity = productStruct.Quantity
		result.Items = append(result.Items, value)

	}
	fmt.Println("", result)
	return result, true, "Sucessfully Completed"
}

func (sale *SaleStruct) GetUserReportByDateRange(obj *models.GetUserReportByDateRange)  ([]models.InvoiceBillById, bool, string) {

	Db, isconnceted := utls.OpenDbConnection()

	res := []models.InvoiceBillById{}
	result := models.InvoiceBillById{}

	if !isconnceted {
		fmt.Println("DB Disconnceted in  GetUserReportByDateRange ")
	}

	query, err := Db.Query(`SELECT id,
										billamount,
										customerid,
										createdon,
										createdby
									FROM invoice
									where createdon between $1 and $2`, obj.FromDate, obj.ToDate)
	if err != nil {
		fmt.Println("Error in QueryRow of  GetUserReportByDateRange :", err)
		return res,false,"Failed"
	}

	// myObj := domain.Sales{}

	for query.Next() {
		err = query.Scan(
			&result.Id,
			&result.BillAmount,
			&result.CustomerId,
			&result.CreatedOn,
			&result.CreatedBy,
		)
		if err != nil {
			fmt.Println("Error in QueryRow Scan in  GetUserReportByDateRange :", err)
			return res,false,"Failed"
		}
		billid := 100001 + result.Id
		value, status, descreption := sale.SaleGetByBillid(&billid)
		if !status {
			fmt.Println(descreption)
			return res,false,"Failed"
		}
		result.Items = append(result.Items, value.Items...)
		res = append(res, result)
	}
	return res,true,"Successfully Completed"
}
