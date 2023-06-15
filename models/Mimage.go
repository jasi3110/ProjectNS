package models

import (
	"strings"
)

type ProductImage struct {
	Id        int64  `json:"id"`
	Imageurl  string `json:"imageurl"`
	Productid int64  `json:"productid"`
	Createdon string `json:"createdon"`
}
// https://drive.google.com/file/d/17JlVUgdJkfR-Zm_d2PmbsIk0KPAujekM/view?usp=share_link
func Imageurl(obj string) string {
	obj = strings.Replace(obj, "https://drive.google.com/file/d/", "", 1)
	obj = strings.Replace(obj, "/view?usp=share_link", "", 1)
	return obj
}

type ProductImageResponses struct {
	Statuscode  int64    `json:"statuscode"`
	Status      bool     `json:"status"`
	Value       []string `json:"imagedata"`
	Descreption string   `json:"desc"`
}