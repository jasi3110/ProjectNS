package models

type Invoice struct {
	Id         int64     `json:"id"`
	Items      []Product `json:"product"`
	BillAmount int64     `json:"billamount"`
	TrxId      int64     `json:"trxid,string"`
	CustomerId int64     `json:"customerid"`
	CreatedOn  string    `json:"createdon"`
	CreatedBy  int64     `json:"createdby"`
}
