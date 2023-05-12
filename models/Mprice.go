package models

type Price struct {
	Id         int64   `json:"id"`
	ProductId  int64   `json:"productid"`
	Mrp        float64 `json:"mrp"`
	Nop        float64 `json:"nop"`
	Percentage int64 `json:"percentage"`
	Createdon  string  `json:"createdon"`
}

type PriceResponses struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Value       Price  `json:"pricedata"`
	Descreption string `json:"desc"`
}

type GetAllPriceResponse struct {
	Statuscode  int64   `json:"statuscode"`
	Status      bool    `json:"status"`
	Value       []Price `json:"pricedata"`
	Descreption string  `json:"desc"`
}
