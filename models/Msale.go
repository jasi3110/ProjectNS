package models

type Invoice struct {
	Id         int64     `json:"id"`
	Items      []Product `json:"product"`
	BillAmount string    `json:"billamount"`
	TrxId      int64     `json:"trxid,string"`
	CustomerId string    `json:"customerid"`
	CreatedOn  string    `json:"createdon"`
	CreatedBy  int64     `json:"createdby"`
}


