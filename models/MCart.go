package models

type Cart struct {
	// Id      int64      `json:"id"`
	Product ProductAllCart `json:"productdata"`
}

type RCart struct {
	Id        int64  `json:"id"`
	Productid int64  `json:"productid"`
	Quantity  string `json:"quantity"`
}
type GetAllCartResponse struct {
	Statuscode       int64        `json:"statuscode"`
	Status           bool         `json:"status"`
	Value            []ProductAllCart `json:"cartdata"`
	Items            int64        `json:"items"`
	Productprice     float64      `json:"mrpprice"`
	Productdiscoiunt float64      `json:"discountprice"`
	Total            float64      `json:"total"`
	Descreption      string       `json:"desc"`
}

type CartRespones struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Descreption string `json:"desc"`
	Total       string `json:"total"`
}

type GetAllCart struct {
	Value []ProductAllCart `json:"cartdata"`
	Items int64        `json:"items"`
	// Productprice     float64 `json:"mrpprice"`
	// Productdiscoiunt float64 `json:"discountprice"`
	// Total            float64 `json:"total"`
	// Descreption      string  `json:"desc"`
}

type ProductAllCart struct {
	Id           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Category     Category `json:"category"`
	Quantity     string   `json:"quantity"`
	CartQuantity string   `json:"cartquantity"`
	Unit         Unit     `json:"unit"`
	Percentage   int64    `json:"percentage"`
	Price        Price    `json:"price"`
	CreatedOn    string   `json:"createdon"`
}
