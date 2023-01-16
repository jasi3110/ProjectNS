package models

// CATEGORY
type Category struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CategoryResponses struct {
	Statuscode  int64    `json:"statuscode"`
	Status      bool     `json:"status"`
	Value       Category `json:"categorydata"`
	Descreption string   `json:"desc"`
}

type GetAllCategoryResponse struct {
	Statuscode  int64      `json:"statuscode"`
	Status     bool       `json:"status"`
	Value       []Category `json:"categorydata"`
	Descreption string     `json:"desc"`
}

/////////////////////////////////////////////////////////////////////////////////////////////////////

// UNIT

type Unit struct {
	Id   string `json:"id"`
	Item string `json:"item"`
}

type GetAllUnitResponse struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Value       []Unit `json:"unitdata"`
	Descreption string `json:"desc"`
}

type UnitResponses struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Value       Unit   `json:"unitdata"`
	Descreption string `json:"desc"`
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ROLE

type Role struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type RoleResponses struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Value     Role   `json:"roledata"`
	Descreption string `json:"desc"`
}

type GetAllRoleResponse struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Value       []Role `json:"roledata"`
	Descreption string `json:"desc"`
}
