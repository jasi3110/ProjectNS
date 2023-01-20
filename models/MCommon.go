package models

import (
	// "OnlineShop/utls"
	// "fmt"
	// "strconv"
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


// func GenerateTrxNumber(userID string)int64{
// // date:=utls.GetCurrentDate()

// 	 trxnum:=""
// 	if len(userID)==1{
// 		trxnum="00000"+userID
// 	}else if len(userID) == 2 {
// 		trxnum="0000"+userID
// 	}else if len(userID) == 3 {
// 		trxnum="000"+userID
// 	}else if len(userID) == 4 {
// 		trxnum="00"+userID
// 	}else {
// 		trxnum="0"+userID
// 	}
// 	// i, err := strconv.Atoi(trxnum)
// 	// if err !=nil {
// 		// fmt.Println("Error in Genrating ",date)
// 	// }
// 	// return int64(i)


// }

// func main(){
// 	// date :=utls.GetCurrentDate()
// // fmt.Println("",date)
// }