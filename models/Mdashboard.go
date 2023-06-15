package models

type Dashboard struct {
	Id       int64       `json:"id"`
	Type     string      `json:"type"`
	ViewType int64       `json:"viewtype"`
	Data     interface{} `json:"data"`
	AdData    []string    `json:"addata"`
}

type CommanResponesDashboard struct {
	Statuscode  int64       `json:"statuscode"`
	Status      bool        `json:"status"`
	Descreption string      `json:"desc"`
	Value       []Dashboard `json:"result"`
}

type DashboardCart1 struct {
	Id       int64        `json:"id"`
	Type     string       `json:"type"`
	ViewType int64        `json:"viewtype"`
	Items    int64        `json:"items"`
	Data     []ProductAll `json:"cartdata"`
}

type CommanResponesDashboardCart struct {
	Statuscode  int64            `json:"statuscode"`
	Status      bool             `json:"status"`
	Descreption string           `json:"desc"`
	Value       []DashboardCart1 `json:"result"`
}
