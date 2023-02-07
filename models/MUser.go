package models

type User struct {
	Id        int64  `json:"id,string"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Mobileno  string `json:"mobileno"`
	Role      string `json:"role"`
	Password  string `json:"password"`
	CreatedOn string `json:"createdon"`
	Token     string `json:"token"`
}

type UserUpdate struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobileno string `json:"mobileno"`
	Role     string `json:"role"`
}

type UserPassword struct {
	Id       int64  `json:"id"`
	Mobileno string `json:"mobileno"`
	Password string `json:"password"`
}
type LoginUser struct {
	Mobileno string `json:"mobileno"`
	Password string `json:"password"`
}

type UserResponseModel struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Value       User   `json:"userdata"`
	Descreption string `json:"desc"`
}

type GetAllUserResponseModel struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Value       []User `json:"userdata"`
	Descreption string `json:"desc"`
}

type UserUpdateResponseModel struct {
	Statuscode  int64      `json:"statuscode"`
	Status      bool       `json:"status"`
	Value     UserUpdate `json:"userdata"`
	Descreption string     `json:"desc"`
}

type UserUpdatePassword struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Descreption string `json:"desc"`
}

type UserverfiyMobileno struct {
	Mobileno string `json:"mobileno"`
	OTP      int64  `json:"otp"`
}