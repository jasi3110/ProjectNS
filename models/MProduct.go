package models

type Product struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	Quantity  string `json:"quantity"`
	Unit      string `json:"unit"`
	Price     string `json:"price"`
	CreatedOn string `json:"createdon"`
}
type ProductAll struct {
	Id        int64    `json:"id"`
	Name      string   `json:"name"`
	Category  Category `json:"category"`
	Quantity  string   `json:"quantity"`
	Unit      Unit     `json:"unit"`
	Price     string   `json:"price"`
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

type IDobj struct {
	Id int64 `json:"id"`
}
