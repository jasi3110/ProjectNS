package utls

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)


func SendEmailOTP(username string,useremail string,otp string) {
	
	// Set up SendGrid API key
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		log.Fatal("SENDGRID_API_KEY environment variable is not set")
	}

	// Compose email
	from := mail.NewEmail("NATIONAL STORE", "nationalstore676@gmail.com")
	subject := " OTP For Verify Our Customer "
	to := mail.NewEmail(username,useremail)
	content := mail.NewContent("text/plain", fmt.Sprintf("Your OTP is: %s", otp))
	sentemail:= mail.NewV3MailInit(from, subject, to, content)

	// Send email using SendGrid
	request := sendgrid.GetRequest(apiKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = http.MethodPost
	request.Body = mail.GetRequestBody(sentemail)
	_, err := sendgrid.API(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully!")
}
