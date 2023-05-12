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

type UserChangePassword struct {
	Id       int64 `json:"id"`
	Password string `json:"password"`
	NewPassword string `json:"newpassword"`
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
	Value       UserUpdate `json:"userdata"`
	Descreption string     `json:"desc"`
}

type UserUpdatePasswordOtp struct {
	Statuscode  int64      `json:"statuscode"`
	Status      bool       `json:"status"`
	Value       Userverify `json:"userdetails"`
	Descreption string     `json:"desc"`
}

type UserUpdatePassword struct {
	Statuscode  int64        `json:"statuscode"`
	Status      bool         `json:"status"`
	Value       UserPassword `json:"userdetails"`
	Descreption string       `json:"desc"`
}

type Userverify struct {
	Id         int64  `json:"id"`
	VerifyUser string `json:"verifyuser"`
	OTP        string `json:"otp"`
}

type UserverfiyOtp struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Value       int64  `json:"userid"`
	Descreption string `json:"desc"`
}
