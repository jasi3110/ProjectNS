package models

type Invoice struct {
	Id         int64     `json:"id"`
	Products   []Product `json:"products"`
	BillAmount int64     `json:"billamount"`
	BillId     int64     `json:"billid"`
	CustomerId int64     `json:"customerid"`
	CreatedOn  string    `json:"createdon"`
	CreatedBy  int64     `json:"createdby"`
	Items       int64     `json:"items"`
}

type InvoiceBillById struct {
	Id         int64        `json:"id"`
	Products   []ProductAll `json:"products"`
	BillAmount int64        `json:"billamount"`
	BillId     int64        `json:"billid"`
	CustomerId int64        `json:"customerid"`
	CreatedOn  string       `json:"createdon"`
	CreatedBy  int64        `json:"createdby"`
	Items      int64        `json:"items"`
}

type SaleCommanRespones struct {
	Statuscode  int64   `json:"statuscode"`
	Status      bool    `json:"status"`
	Value       Invoice `json:"invoicedata"`
	Descreption string  `json:"desc"`
}

type GetAllSaleInvoiceResponse struct {
	Statuscode  int64     `json:"statuscode"`
	Status      bool      `json:"result"`
	Value       []Invoice `json:"data"`
	Descreption string    `json:"desc"`
}

type GetAllSaleInvoiceGetByBillIdResponse struct {
	Statuscode  int64           `json:"statuscode"`
	Status      bool            `json:"result"`
	Value       InvoiceBillById `json:"data"`
	Descreption string          `json:"desc"`
}

type GetUserReportByDateRange struct {
	FromDate string `json:"fromdate"`
	ToDate   string `json:"todate"`
}

type GetAllSaleInvoiceByDateRangeResponse struct {
	Statuscode  int64             `json:"statuscode"`
	Status      bool              `json:"result"`
	Value       []InvoiceBillById `json:"data"`
	Descreption string            `json:"desc"`
}
