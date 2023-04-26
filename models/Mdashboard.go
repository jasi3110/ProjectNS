package models

type Dashboard struct {
	Id       int64      `json:"id"`
	Type     string     `json:"type"`
	ViewType int64   `json:"viewtype"`
	Data     interface{} `json:"data"`
}

type CommanResponesDashboard struct {
	Statuscode  int64       `json:"statuscode"`
	Status      bool        `json:"status"`
	Descreption string      `json:"desc"`
	Value       []Dashboard `json:"result"`
}
