package models

type Cart struct {
	// Id      int64      `json:"id"`
	Product ProductAll `json:"productdata"`
}

type RCart struct {
	Id        int64  `json:"id"`
	Productid int64  `json:"productid"`
	Quantity  string `json:"quantity"`

}
type GetAllCartResponse struct {
	Statuscode       int64   `json:"statuscode"`
	Status           bool    `json:"status"`
	Value            []ProductAll  `json:"cartdata"`
	Items            int64   `json:"items"`
	Productprice     float64 `json:"mrpprice"`
	Productdiscoiunt float64 `json:"discountprice"`
	Total            float64 `json:"total"`
	Descreption      string  `json:"desc"`
}

type CartRespones struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Descreption string `json:"desc"`
	Total       string `json:"total"`
}

type GetAllCart struct {
	Value            []ProductAll  `json:"cartdata"`
	Items            int64   `json:"items"`
	// Productprice     float64 `json:"mrpprice"`
	// Productdiscoiunt float64 `json:"discountprice"`
	// Total            float64 `json:"total"`
	// Descreption      string  `json:"desc"`
}