package models

type Price struct {
	Id        int64  `json:"id"`
	ProductId int64  `json:"productid"`
	Mrp       int64  `json:"mrp"`
	Nop       int64  `json:"nop"`
	Createdon string `json:"createdon"`
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
