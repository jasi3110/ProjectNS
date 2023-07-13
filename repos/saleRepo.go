package repos

import (
	"OnlineShop/models"
	"OnlineShop/repos/masterRepo"
	"OnlineShop/utls"
	"log"
	"strconv"
)

type SaleInterface interface {
	CreateSale(Obj *models.Invoice) (bool, string, models.Invoice)
	InvoiceGetall() ([]models.Invoice, bool, string)
	InvoiceGetallByUserid(obj *int64) ([]models.Invoice, bool,string) 
	GetSaleByInvoiceid(obj *int64) (models.InvoiceSaleById, bool, string)
	InvoiceByDateRange(obj *models.InvoiceByDateRange) ([]models.InvoiceSaleById, bool, string)
	SaleDelete(obj *models.Invoice) (bool, string)
}

type SaleStruct struct {
}

func (sale *SaleStruct) CreateSale(obj *models.Invoice) (bool, string, models.Invoice) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		log.Panic("DB is Disconnceted in CreateSale ")
	}

	// create transaction
	Txn, err := Db.Begin()
	if err != nil {
		log.Panic("Error in CreateSale Transacation in DB :", err)
		return false, "Something Went Wrong", *obj
	}

	//write invoice
	err = Txn.QueryRow(`INSERT into "invoice"(billamount,userid,createdon,items)values($1,$2,$3,$4)RETURNING id`,
		obj.BillAmount,
		obj.UserId,
		utls.GetCurrentDate(),
		obj.Items,
	).Scan(&obj.Id)

	log.Println("invoice created:", obj.Id)

	if err != nil {
		log.Panic("Error in CreateSale Invoice QueryRow :", err)
		err := Txn.Rollback()
		if err != nil {
			log.Panic("Error in CreateSale Rollback in Invoice :", err)
		}
		return false, "Something Went Wrong", *obj
	}

	//write sales entry data from array
	id := 0
	for _, productItem := range obj.Products {

		err := Txn.QueryRow(`INSERT into "saleentry"(userid,invoiceid,productid,productprice,quantity,
			createdon,isdeleted)values($1,$2,$3,$4,$5,$6,$7) RETURNING id`,
			obj.UserId,
			obj.Id,
			productItem.Id,
			productItem.Price.Id,
			productItem.Quantity,
			utls.GetCurrentDate(),
			0,
		).Scan(&id)

		log.Println("Saleentry created:", id)

		if err != nil {
			log.Panic("Error in CreateSale SaleEntry QueryRow Scan :", err)
			err := Txn.Rollback()
			if err != nil {
				log.Panic("Error in CreateSale Rollback in SaleEntry :", err)
			}
			return false,"Something Went Wrong", *obj
		}

		productRepo := ProductInterface(&ProductStruct{})
		value, proStatus, _ := productRepo.GetProductById(&productItem.Id)

		if !proStatus {
			log.Panic("Error in CreateSale Getting Product Data :", err)
			err := Txn.Rollback()
			if err != nil {
				log.Panic("Error CreateSale Getting Product Data in Rollback :", err)

				return false,"Something Went Wrong", *obj
			}

		}

		productqty, _ := strconv.ParseFloat(value.Quantity, 32)
		productqty1, _ := strconv.ParseFloat(productItem.Quantity, 32)

		if productqty != 0 && productqty1 != 0 {
			productqty = productqty - productqty1
			quatity := strconv.FormatFloat(productqty, 'f', -1, 64)
			log.Println("", productItem.Quantity, quatity, productItem.Id)

			if productqty >= 0 {
				_, err := Txn.Exec(`UPDATE  "product" SET  quantity=$1 WHERE id=$2 and isdeleted=0`,
					quatity, productItem.Id)

				if err != nil {
					log.Panic("Error in CreateSale in Product Update QueryRow  :", err)
					err := Txn.Rollback()
					if err != nil {
						log.Panic("Error in CreateSale Rollback in Product Update QueryRow :", err)
					}
					return false, "Something Went Wrong", *obj
				}
			}

		}
	}
	err = Txn.Commit()
	log.Println("transcration commited in Created Sale")
	if err != nil {
		log.Panic("transcration commit Failed")
		err := Txn.Rollback()
		if err != nil {
			log.Panic("Error in CreateSale Rollback in Product Update :", err)

			return false, "Something Went Wrong", *obj
		}
	}
	defer func() {
		Db.Close()
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	return true, "CREATE SALE SUCCESSFULLY COMPLETED", *obj
}

func (sale *SaleStruct) InvoiceGetall() ([]models.Invoice, bool, string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		log.Panic("DB Disconnceted in invoice Getall")
	}
	invoiceStruct := models.Invoice{}
	result := []models.Invoice{}

	query, err := Db.Query(`SELECT id,billamount,userid,createdon,createdby FROM "invoice" where isdeleted=1`)
	if err != nil {
		log.Println("Error in Invoice GetAll Query : ", err)
	}

	for query.Next() {
		err := query.Scan(
			&invoiceStruct.Id,
			&invoiceStruct.BillAmount,
			&invoiceStruct.UserId,
			&invoiceStruct.CreatedOn,
		)
		if err != nil {
			log.Panic("Error in User GetAll QueryRow :", err)
			return result, false, "Something Went Wrong"
		}
		result = append(result, invoiceStruct)
	}
	defer func() {
		Db.Close()
		query.Close()

		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()
	return result, true, "Sucessfully Compeleted"
}

func (sale *SaleStruct) InvoiceGetallByUserid(obj *int64) ([]models.Invoice, bool,string) {
	Db, isConnected := utls.OpenDbConnection()
	if !isConnected {
		log.Panic("DB Disconnceted in invoice Getall")
	}
	invoiceStruct := models.Invoice{}
	result := []models.Invoice{}

	query, err := Db.Query(`SELECT id,billamount,userid,createdon,items FROM "invoice" WHERE customerid=$1  and isdeleted=0`, obj)
	if err != nil {
		log.Panic("Error in Invoice GetAll By Customerid QueryRow : ", err)
		return result, false,"Something Went Wrong"
	}
	for query.Next() {
		err = query.Scan(
			&invoiceStruct.Id,
			&invoiceStruct.BillAmount,
			&invoiceStruct.UserId,
			&invoiceStruct.CreatedOn,
			&invoiceStruct.Items,
		)
		if err != nil {
			log.Panic("Error in User GetAll QueryRow Scan :", err)
			return result, false,"Something Went Wrong"
		}
		result = append(result, invoiceStruct)
	}
	defer func() {
		Db.Close()
		query.Close()
		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()
	return result, true,"Successfully Completed"
}

func (sale *SaleStruct) GetSaleByInvoiceid(obj *int64) (models.InvoiceSaleById, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		log.Panic("DB is Disconnceted in SaleGetByInvoiceid")
	}

	result := models.InvoiceSaleById{}
	productStruct := models.ProductAll{}

	query, err := Db.Query(`SELECT   id,userid,invoiceid,productid,productprice,quantity,createdon
	FROM "saleentry" WHERE invoiceid=$1 and isdeleted=0`, obj)

	if err != nil {
		log.Panic("Error in SaleGetByCustomerid QueryRow :", err)
		return result, false,"Something Went Wrong"
	}

	for query.Next() {
		err := query.Scan(&result.Id,
			&result.UserId,
			&result.InvoiceId,
			&productStruct.Id,
			&productStruct.Price.Id,
			&productStruct.Quantity,
			&result.CreatedOn,
		)
		if err != nil {
			log.Panic("Error in SaleGetByInvoiceid QueryRow Scan : ", err)
			return result, false, "Something Went Wrong"
		}

		productRepo := ProductInterface(&ProductStruct{})
		value, status, _ := productRepo.GetProductById(&productStruct.Id)
		if !status {
			log.Panic("Error Getting Product Data in SaleGetByinvoiceid ")
			return result, false, "Something Went Wrong"
		}

		priceRepo := masterRepo.PriceInterface(&masterRepo.PriceStruct{})
		valueprice, statusprice, _ := priceRepo.PriceById(&productStruct.Price)
		if !statusprice {
			log.Panic("Error Getting Price Data in SaleGetByInvoiceid ")
			return result, false, "Something Went Wrong"
		}
		value.Price.Id = valueprice.Id
		value.Price.Mrp = valueprice.Mrp
		value.Price.Nop = valueprice.Nop
		value.Quantity = productStruct.Quantity
		result.Products = append(result.Products, value)

	}
	defer func() {
		Db.Close()
		query.Close()

		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()
	return result, true, "Successfully Completed"
}

func (sale *SaleStruct) InvoiceByDateRange(obj *models.InvoiceByDateRange) ([]models.InvoiceSaleById, bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		log.Panic("DB Disconnceted in  InvoiceByDateRange ")
	}

	result := []models.InvoiceSaleById{}
	invoiceStruct := models.InvoiceSaleById{}

	query, err := Db.Query(`SELECT id,billamount,userid,createdon,items
							FROM "invoice" where createdon between $1 and $2`, obj.FromDate, obj.ToDate)
	if err != nil {
		log.Panic("Error in QueryRow of  InvoiceByDateRange :", err)
		return result, false, "Something Went Wrong"
	}

	for query.Next() {
		err = query.Scan(
			&invoiceStruct.Id,
			&invoiceStruct.BillAmount,
			&invoiceStruct.UserId,
			&invoiceStruct.CreatedOn,
			&invoiceStruct.Items,
		)
		if err != nil {
			log.Panic("Error in QueryRow Scan in  InvoiceByDateRange :", err)
			return result, false,"Something Went Wrong"
		}

		result = append(result, invoiceStruct)
	}

	defer func() {
		Db.Close()
		query.Close()

		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}
	}()

	return result, true, "Successfully Completed"
}

func (sale *SaleStruct) SaleDelete(obj *models.Invoice) (bool, string) {
	Db, isconnceted := utls.OpenDbConnection()
	if !isconnceted {
		log.Panic("DB disconnceted in Sale Delete ")
	}

	txn, err := Db.Begin()
	if err != nil {
		log.Panic("Error in Sale Delete Transacation :", err)
		return false, "Something Went Wrong"
	}

	_,err = txn.Exec(`UPDATE "invoice" SET isdeleted=1  WHERE id=$1 `, obj.Id)
	if err != nil {
		log.Panic("Error in Sale Delete Invoice QueryRow  :", err)

		err := txn.Rollback()
		if err != nil {
			log.Panic("Error in Sale Delete Invioce Rollback :", err)
		}
		return false, "Something Went Wrong"
	}

	if err != nil {
		log.Panic("Error in Delete Invioce QueryRow Close :", err)
	}

	query, err := txn.Query(`UPDATE "saleentry" SET isdeleted=1  WHERE billid=$1`, obj.Id)
	for query.Next() {
		if err != nil {
			log.Panic("Error in Delete SaleEntry QueryRow  :", err)
			err = txn.Rollback()
			if err != nil {
				log.Panic("Error in Delete SaleEntry Rollback :", err)
			}
			return false, "Something Went Wrong"
		}
	}

	txn.Commit()

	defer func() {
		Db.Close()
		query.Close()

		if r := recover(); r != nil {
			log.Panic("Recovered from panic condition : ", r)
		}

	}()
	return true, "Sucessfully Completed"
}
