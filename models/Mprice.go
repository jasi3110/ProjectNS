package models

type Price struct {
	Id          string `json:"id"`
	ProductId   int64 `json:"productid"`
	ProductPrice string   `json:"productprice"`
	Createdon     string   `json:"createdon"`
}
