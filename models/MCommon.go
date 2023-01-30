package models

import (
	// "OnlineShop/utls"
	// "fmt"
	// "strconv"
	// "fmt"
	"strings"
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


// func GenerateTrxNumber()int64{
// // date:=utls.GetCurrentDate()

// 	 *trxnum:=1000 +1
// 	 sum:=trxnum
// 	 trxnum=&sum
// 	return int64(sum)
// 	}



// func main(){
// 	// date :=utls.GetCurrentDate()

// fmt.Println("",GenerateTrxNumber())
// fmt.Println("",GenerateTrxNumber())
// }