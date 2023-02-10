package models

type Discount struct {
	Id         int64      `json:"id"`
	Product    ProductAll `json:"productid"`
	Percentage int64      `json:"percentage"`
	Mrp        int64      `json:"mrp"`
	Nop        int64      `json:"nop"`
	Startend   string     `json:"startend"`
	Enddate    string     `json:"enddate"`
}

type RDiscount struct {
	Id         int64  `json:"id"`
	Percentage int64  `json:"percentage"`
	Enddate    string `json:"enddate"`
}

type GetAllDiscountResponse struct {
	Statuscode  int64        `json:"statuscode"`
	Status      bool         `json:"status"`
	Value       [] ProductAll `json:"discountdata"`
	Descreption string       `json:"desc"`
}
