package main

import (
	"fmt"
	"time"

	"github.com/BourgeoisBear/email.v2"
)

func main() {

	oCfg := email.SMTPClientConfig{
		Server:      "smtp.gmail.com",
		Port:        587,
		Username:    "rushikesh.swami.kanaka@gmail.com",
		Password:    "fdtkaoauhuqwdxwp",
		Mode:        email.ModeSTARTTLS, //    "STARTTLS",
		TimeoutMsec: 10000,
		Proto:       "tcp4",
		SMTPLog:     "-", // note: uncomment to log SMTP session to STDOUT
	}

	oEmail := email.NewEmail()
	oEmail.From = "rushikesh.swami.kanaka@gmail.com"
	oEmail.To = []string{"mahadev.godbole@kanakasoftware.com"}
	oEmail.Subject = "Test Message"
	time := time.Now()

	oEmail.Text = []byte("Test message body !!" + time.String())

	E := oCfg.SimpleSend(oEmail)
	if E != nil {
		fmt.Println("\n\n", time, "Unable to send email", E)
		return
	}
	fmt.Println("\n\n", time, "Successfully mail sent to send email", E)
}
