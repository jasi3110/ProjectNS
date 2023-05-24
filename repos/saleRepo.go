package repos

import (
	"OnlineShop/models"
	"OnlineShop/repos/masterRepo"
	"OnlineShop/utls"
	"fmt"
	"log"
	"strconv"
	// "github.com/spf13/pflag"
)

type SaleInterface interface {
	CreateSale(Obj *models.Invoice) (bool, string, models.Invoice)

	// INVOICE
	InvoiceGetall() ([]models.Invoice, bool, string)
	InvoiceGetallByCustomerid(obj *int64) ([]models.Invoice, bool)

	// SALE ENTRY
	SaleGetByBillid(obj *int64) (models.InvoiceBillById, bool, string)
	SaleGetByCustomerid(obj *int64) (models.InvoiceBillById, bool, string)
	GetUserReportByDateRange(obj *models.GetUserReportByDateRange) ([]models.InvoiceBillById, bool, string)
	SaleGetByDate(obj *string) ([]models.InvoiceBillById, bool, string)
	SaleDelete(obj *models.Invoice) (bool, string)
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
		return false, "CREATE SALE FAILED", *obj
	}

	//write invoice
	err = Txn.QueryRow(`INSERT into "invoice"(billamount,customerid,createdon,items)values($1,$2,$3,$4,$5)RETURNING id`,
		obj.BillAmount,
		obj.CustomerId,
		utls.GetCurrentDate(),
		obj.CreatedBy,
		obj.Items,
	).Scan(&obj.Id)

	log.Println("invoice created:", obj.Id)

	if err != nil {
		fmt.Println("Error in CreateSale Invoice QueryRow :", err)
		err := Txn.Rollback()
		if err != nil {
			fmt.Println("Error in CreateSale Rollback in Invoice :", err)
		}
		return false, "CREATE SALE FAILED", *obj
	}
	//write sales entry data from array

	for _, productItem := range obj.Products {

		err := Txn.QueryRow(`INSERT into "saleentry"(customerid,
												billid,
												invoiceid,
			     								productid,
												productprice,
												quantity,
												createdon,
												createdby,
												isdeleted)
												values($1,$2,$3,$4,$5,$6,$7,$8,$9)RETURNING id`,
			obj.CustomerId,
			obj.Id,
			obj.Id,
			productItem.Id,
			productItem.Price,
			productItem.Quantity,
			utls.GetCurrentDate(),
			obj.CreatedBy,
			0,
		).Scan(&a)
		fmt.Println("", a)
		if err != nil {
			fmt.Println("Error in CreateSale SaleEntry QueryRow Scan :", err)
			err := Txn.Rollback()
			if err != nil {
				fmt.Println("Error in CreateSale Rollback in SaleEntry :", err)
			}
			return false, "CREATE SALE FAILED", *obj
		}
		// var productchannel chan models.ProductAll
		productRepo := ProductInterface(&ProductStruct{})
		value, _,_:=productRepo.GetProductById(&productItem.Id)
		// value := <-productchannel
		productqty, _ := strconv.ParseFloat(value.Quantity, 32)
		productqty1, _ := strconv.ParseFloat(productItem.Quantity, 32)
		if productqty != 0 || true {
			//reduce stock quantity from product table
			productqty = productqty - productqty1
			quatity := strconv.FormatFloat(productqty, 'E', -1, 64)
			updateQueryqty, err := Txn.Query(`UPDATE  "product" SET  quantity=$1 WHERE id=$2 and isdeleted=1`,
				quatity, productItem.Id)

			if err != nil {
				fmt.Println("Error in CreateSale in Product Update QueryRow  :", err)
				err := Txn.Rollback()
				if err != nil {
					fmt.Println("Error in CreateSale Rollback in Product Update QueryRow :", err)
				}
				return false, "CREATE SALE FAILED", *obj
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
			return false, "CREATE SALE FAILED", *obj
		}
	}
	defer func() {
		Db.Close()
	}()
	return true, "CREATE SALE SUCCESSFULLY COMPLETED", *obj
}

func (sale *SaleStruct) InvoiceGetall() ([]models.Invoice, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in invoice Getall")
	}
	invoiceStruct := models.Invoice{}
	result := []models.Invoice{}

	query, err := Db.Query(`SELECT id,billamount,customerid,createdon,createdby FROM "invoice" where isdeleted=1`)
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
			return result, false, "Failed"
		}
		result = append(result, invoiceStruct)
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "Sucessfully Compeleted"
}

func (sale *SaleStruct) InvoiceGetallByCustomerid(obj *int64) ([]models.Invoice, bool) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		fmt.Println("DB Disconnceted in invoice Getall")
	}
	invoiceStruct := models.Invoice{}
	result := []models.Invoice{}

	query, err := Db.Query(`SELECT id,billamount,customerid,createdon,items createdby FROM "invoice" WHERE customerid=$1  and isdeleted=0`, obj)
	if err != nil {
		log.Println(err)
		return result, false
	}
	for query.Next() {
		err = query.Scan(
			&invoiceStruct.Id,
			&invoiceStruct.BillAmount,
			&invoiceStruct.CustomerId,
			&invoiceStruct.CreatedOn,
			&invoiceStruct.Items,
		)
		fmt.Println("", invoiceStruct)
		if err != nil {
			fmt.Println("Error in User GetAll QueryRow :", err)
			return result, false
		}
		result = append(result, invoiceStruct)
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true
}

// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (sale *SaleStruct) SaleGetByCustomerid(obj *int64) (models.InvoiceBillById, bool, string) {
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
												WHERE customerid=$1 and isdeleted=0`, obj)
	if err != nil {
		fmt.Println("Error in SaleGetByBillid QueryRow :", err)
		return result, false, "Failed"
	}
	// var productchannel chan models.ProductAll
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
		value, _,_:= productRepo.GetProductById(&productStruct.Id)
		// value := <-productchannel
		priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
		valueprice, statusprice, descreptionprice := priceRepo.PriceById(&productStruct.Price)
		if !statusprice {
			fmt.Println(descreptionprice)
			return result, false, "Failed"
		}
		value.Price.Mrp = valueprice.Mrp
		value.Price.Nop = valueprice.Nop
		value.Quantity = productStruct.Quantity
		result.Products = append(result.Products, value)

	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "Sucessfully Completed"
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
												WHERE billid = $1 and isdeleted=0`, obj)
	if err != nil {
		fmt.Println("Error in SaleGetByBillid QueryRow :", err)
		return result, false, "Failed"
	}
	// var productchannel chan models.ProductAll
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
		value, _,_:= productRepo.GetProductById(&productStruct.Id)

		priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
		valueprice, statusprice, descreptionprice := priceRepo.PriceById(&productStruct.Price)
		if !statusprice {
			fmt.Println(descreptionprice)
			return result, false, "Failed"
		}
		// value := <-productchannel
		value.Price.Mrp = valueprice.Mrp
		value.Price.Nop = valueprice.Nop
		value.Quantity = productStruct.Quantity
		result.Products = append(result.Products, value)

	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return result, true, "Successfully Completed"
}

func (sale *SaleStruct) GetUserReportByDateRange(obj *models.GetUserReportByDateRange) ([]models.InvoiceBillById, bool, string) {

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
		return res, false, "Failed"
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
			return res, false, "Failed"
		}
		billid := 100001 + result.Id
		value, status, descreption := sale.SaleGetByBillid(&billid)
		if !status {
			fmt.Println(descreption)
			return res, false, "Failed"
		}
		result.Products = append(result.Products, value.Products...)
		res = append(res, result)
	}
	defer func() {
		Db.Close()
		query.Close()
	}()
	return res, true, "Successfully Completed"
}

func (sale *SaleStruct) SaleGetByDate(obj *string) ([]models.InvoiceBillById, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	res := []models.InvoiceBillById{}
	result := models.InvoiceBillById{}
	productStruct := models.ProductAll{}

	if !isconnceted {
		fmt.Println("DB is Disconnceted in SaleGetByBillid")
		return res, false, "Failed"
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
												WHERE createdon = $1`, obj)
	if err != nil {
		fmt.Println("Error in SaleGetByBillid QueryRow :", err)
		return res, false, "Failed"
	}
	// var productchannel chan models.ProductAll
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
			return res, false, "Failed"
		}
		productRepo := ProductInterface(&ProductStruct{})
		value, _,_:= productRepo.GetProductById(&productStruct.Id)
		// value := <-productchannel
		priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
		valueprice, statusprice, descreptionprice := priceRepo.PriceById(&productStruct.Price)
		if !statusprice {
			fmt.Println(descreptionprice)
			return res, false, "Failed"
		}
		value.Price.Mrp = valueprice.Mrp
		value.Price.Nop = valueprice.Nop
		value.Quantity = productStruct.Quantity
		result.Products = append(result.Products, value)
		res = append(res, result)
	}

	defer func() {
		Db.Close()
		query.Close()
	}()
	return res, true, "Sucessfully Completed"
}

func (sale *SaleStruct) SaleDelete(obj *models.Invoice) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		fmt.Println("DB disconnceted in SaleDelete ")
	}
	txn, err := Db.Begin()
	if err != nil {
		fmt.Println("Error in SaleDelete Transacation in DB :", err)
		return false, "DELETETSALE FAILED"
	}
	err = txn.QueryRow(`UPDATE "invoice" SET isdeleted=1  WHERE id=$1 `, obj.Id).Scan(&obj.BillId)

	if err != nil {
		fmt.Println("Error in Delete Invoice QueryRow  :", err)
		err := txn.Rollback()
		if err != nil {
			fmt.Println("Error in Delete Invioce Rollback :", err)
		}
		return false, "CREATESALE FAILED"
	}

	if err != nil {
		fmt.Println("Error in Delete Invioce QueryRow Close :", err)
	}

	DeleteQuerySaleEntry, err := txn.Query(`UPDATE "saleentry" SET isdeleted=1  WHERE billid=$1`, obj.BillId)
	for DeleteQuerySaleEntry.Next() {
		if err != nil {
			fmt.Println("Error in Delete SaleEntry QueryRow  :", err)
			err = txn.Rollback()
			if err != nil {
				fmt.Println("Error in Delete SaleEntry Rollback :", err)
			}
			return false, "CREATESALE FAILED"
		}
		err = DeleteQuerySaleEntry.Close()
		if err != nil {
			fmt.Println("Error in Delete SaleEntry QueryRow Close :", err)
		}
	}
	txn.Commit()
	defer func() {
		Db.Close()
	}()
	return true, "User SaleEntry Deleted Sucessfully Completed"
}
