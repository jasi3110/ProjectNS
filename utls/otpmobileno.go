// Download the helper library from https://www.twilio.com/docs/go/install
package utls

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/api/v2010"
)
var(
	accountSid string
	authToken  string
	fromPhone  string
	client 	   *twilio.RestClient
)
func MobilenoOTP(toPhone string,UserName string,otp string) {
	// client := twilio.NewRestClient()

	msg:="Hello "+UserName+"\n"+"NATIONAL STORE VERFICATION OTP :"+otp 
	params := verify.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetTo(fromPhone)
	params.SetBody(msg)

	resp, err := client.Api.CreateMessage(&params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Sid != nil {
			fmt.Println(*resp.Sid)
		} else {
			fmt.Println(resp.Sid)
		}
	}
}

func init(){
	err:=godotenv.Load("cred.env")
	if err !=nil{
		fmt.Println("error loading .env :",err)
		os.Exit(1)
	}
	accountSid = os.Getenv("ACCOUNT_SID")
	authToken  = os.Getenv("AUTH_TOKEN")
	fromPhone  = os.Getenv("FROM_PHONE")
	client	   = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   accountSid,
		Password:   authToken,
		AccountSid: accountSid,
		Client:     nil,
	})
	
}