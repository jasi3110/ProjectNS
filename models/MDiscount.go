package models

type Discount struct {
	Id         int64      `json:"id"`
	Product    ProductAll `json:"productid"`
	Percentage float64    `json:"percentage"`
	Priceid    int64      `json:"priceid"`
	Oldpriceid int64      `json:"oldpriceid"`
	Startend   string     `json:"startend"`
	Enddate    string     `json:"enddate"`
}

type RDiscount struct {
	Id         int64   `json:"id"`
	Percentage float64 `json:"percentage"`
	Enddate    string  `json:"enddate"`
}

type GetAllDiscountResponse struct {
	Statuscode  int64        `json:"statuscode"`
	Status      bool         `json:"status"`
	Value       []ProductAll `json:"discountdata"`
	Descreption string       `json:"desc"`
}
