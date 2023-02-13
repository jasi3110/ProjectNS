package models

type Cart struct {
	Id               int64      `json:"id"`
	Product          ProductAll `json:"productdata"`
	Items            int64      `json:"items"`
	Productdiscoiunt float64    `json:"productdiscount"`
	Total            float64    `json:"total"`
}

type RCart struct {
	Id        int64  `json:"id"`
	Productid int64  `json:"productid"`
	Quantity  string `json:"quantity"`
}

type GetAllCartResponse struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Value       []Cart `json:"discountdata"`
	Descreption string `json:"desc"`
}
