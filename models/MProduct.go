package models

type Product struct {
	Id        int64  `json:"id"`
	Image     string `json:"image"`
	Name      string `json:"name"`
	Category  int64  `json:"category"`
	Quantity  string `json:"quantity"`
	Unit      int64  `json:"unit"`
	Price     int64  `json:"price"`
	Mrp       float64  `json:"mrp"`
	Nop       float64  `json:"nop"`
	CreatedOn string `json:"createdon"`
}
type ProductAll struct {
	Id        int64    `json:"id"`
	Image     string   `json:"image"`
	Name      string   `json:"name"`
	Category  Category `json:"category"`
	Quantity  string   `json:"quantity"`
	Unit      Unit     `json:"unit"`
	Percentage int64    `json:"percentage"`
	Price     Price    `json:"price"`
	CreatedOn string   `json:"createdon"`
}
type ProductResponses struct {
	Statuscode  int64      `json:"statuscode"`
	Status      bool       `json:"result"`
	Value       ProductAll `json:"data"`
	Descreption string     `json:"desc"`
}

type GetAllProductResponse struct {
	Statuscode  int64        `json:"statuscode"`
	Status      bool         `json:"result"`
	Value       []ProductAll `json:"data"`
	Descreption string       `json:"desc"`
}

type DiscountProductAll struct {
	Id         int64    `json:"id"`
	Image      string   `json:"image"`
	Name       string   `json:"name"`
	Category   Category `json:"category"`
	Quantity   string   `json:"quantity"`
	Unit       Unit     `json:"unit"`
	Percentage int64    `json:"precentage"`
	Price      Price    `json:"price"`
	CreatedOn  string   `json:"createdon"`
}

type IDobj struct {
	Id int64 `json:"id"`
}
