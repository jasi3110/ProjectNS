package models

type Invoice struct {
	Id         int64        `json:"id"`
	Products   []ProductAll `json:"productdata"`
	BillAmount int64        `json:"billamount"`
	CustomerId int64        `json:"customerid"`
	CreatedOn  string       `json:"createdon"`
	Items      int64        `json:"items"`
}

type InvoiceSaleById struct {
	Id         int64        `json:"id"`
	Products   []ProductAll `json:"productdata"`
	BillAmount int64        `json:"billamount"`
	InvoiceId  int64        `json:"Invoiceid"`
	CustomerId int64        `json:"customerid"`
	CreatedOn  string       `json:"createdon"`
	Items      int64        `json:"items"`
}

type SaleCommanRespones struct {
	Statuscode  int64   `json:"statuscode"`
	Status      bool    `json:"status"`
	Value       Invoice `json:"invoicedata"`
	Descreption string  `json:"desc"`
}

type GetAllInvoiceResponse struct {
	Statuscode  int64     `json:"statuscode"`
	Status      bool      `json:"result"`
	Value       []Invoice `json:"data"`
	Descreption string    `json:"desc"`
}

type GetSaleByInvoiceIdResponse struct {
	Statuscode  int64           `json:"statuscode"`
	Status      bool            `json:"result"`
	Value       InvoiceSaleById `json:"data"`
	Descreption string          `json:"desc"`
}

type InvoiceByDateRange struct {
	FromDate string `json:"fromdate"`
	ToDate   string `json:"todate"`
}

type GetAllSaleInvoiceByDateRangeResponse struct {
	Statuscode  int64             `json:"statuscode"`
	Status      bool              `json:"result"`
	Value       []InvoiceSaleById `json:"data"`
	Descreption string            `json:"desc"`
}
