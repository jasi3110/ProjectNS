package models

import (
	// "OnlineShop/utls"
	// "fmt"
	// "strconv"
	// "fmt"
	"fmt"
	"strings"
	"time"
	"unicode"
)

type CommanRespones struct {
	Statuscode  int64  `json:"statuscode"`
	Status      bool   `json:"status"`
	Descreption string `json:"desc"`
}

type Status struct {
	Result bool `json:"result"`
}

func VerifyMobileno(a string) bool {
	length := len(a)
	return length == 10
}
func VerifyEmail(a string) bool {
	return strings.Contains(a, "@") && strings.Contains(a, ".")
}
func VerifyPassword(s string) bool {
	letters := 0
	special := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c) || unicode.IsPunct(c) || unicode.IsSymbol(c) || c == ' ':
			special++
		case unicode.IsLetter(c):
			letters++
		}
	}
	res := letters + special
	return res >= 8 && special >= 1 && letters >= 1
}


func GenerateOtp()string{
	genotp := fmt.Sprint(time.Now().Nanosecond())
	otp:=genotp[0:4]
	return otp
	}



// func main(){
// 	// date :=utls.GetCurrentDate()

// fmt.Println("",GenerateTrxNumber())
// fmt.Println("",GenerateTrxNumber())
// }