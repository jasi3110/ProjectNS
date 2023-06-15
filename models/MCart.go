package models

type Cart struct {
	Product ProductAll `json:"productdata"`
}

type RCart struct {
	Id        int64  `json:"id"`
	Productid int64  `json:"productid"`
	Quantity  string `json:"quantity"`
}
type GetAllCartResponse struct {
	Statuscode       int64        `json:"statuscode"`
	Status           bool         `json:"status"`
	Value            []CartProductAll `json:"cartdata"`
	Items            int64        `json:"items"`
	Productprice     float64      `json:"mrpprice"`
	Productdiscoiunt float64      `json:"discountprice"`
	Total            float64      `json:"total"`
	Descreption      string       `json:"desc"`
}

type CartProductAll struct {
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
type CartRespones struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Descreption string `json:"desc"`
	Total       string `json:"total"`
}

type GetAllCart struct {
	Value []CartProductAll `json:"cartdata"`
	Items int64        `json:"items"`
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
