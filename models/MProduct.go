package models

type Product struct {
	Id         int64   `json:"id"`
	Image      string  `json:"image"`
	Name       string  `json:"name"`
	Category   int64   `json:"category"`
	Quantity   string  `json:"quantity"`
	Unit       int64   `json:"unit"`
	Price      string  `json:"price"`
	Percentage int64   `json:"percentage"`
	Mrp        float64 `json:"mrp"`
	Nop        float64 `json:"nop"`
	CreatedOn  string  `json:"createdon"`
}


type ProductAll struct {
	Id         int64    `json:"id"`
	Image      string   `json:"image"`
	Name       string   `json:"name"`
	Category   Category `json:"category"`
	Quantity   string   `json:"quantity"`
	Unit       Unit     `json:"unit"`
	Percentage int64    `json:"percentage"`
	Price      Price    `json:"price"`
	CreatedOn  string   `json:"createdon"`
}
type ProductResponses struct {
	Statuscode  int64      `json:"statuscode"`
	Status      bool       `json:"result"`
	Value       ProductAll `json:"productdata"`
	Descreption string     `json:"desc"`
}

type GetAllProductResponse struct {
	Statuscode  int64        `json:"statuscode"`
	Status      bool         `json:"result"`
	Value       []ProductAll `json:"productdata"`
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

type ProductSearchResponses struct {
	Statuscode int64        `json:"statuscode"`
	Status     bool         `json:"result"`
	Value      []ProductAll `json:"searchdata"`
}

type ProductSearch struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
