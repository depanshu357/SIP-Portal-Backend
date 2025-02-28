package services

import (
	"crypto/rand"
	"math/big"
	"net/smtp"
	"os"
)

func GenerateOTP(otp_size int) (string, error) {
	otp := ""
	for i := 0; i < otp_size; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(10)) // Generates a number between 0 and 9
		if err != nil {
			return "", err
		}
		otp += num.String()
	}
	return otp, nil
}

func SendMail(email string, otp string) error {
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("SMTP_HOST"))

	from := os.Getenv("EMAIL")
	subject := "Subject: OTP for Verification\n"
	toHeader := "To: " + email + "\n"
	fromHeader := "From: " + from + "\n"
	mimeHeader := "MIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n"
	htmlMessage := `
<html>
  <head>
    <style>
      body { font-family: Arial, sans-serif; background-color: #f9f9f9; margin: 0; padding: 0; }
      .container { max-width: 600px; margin: 20px auto; background-color: #ffffff; padding: 20px; border: 1px solid #ddd; }
      .header { text-align: center; padding: 10px 0; background-color: #4CAF50; color: #ffffff; }
      .content { padding: 20px; font-size: 16px; color: #333333; }
      .otp { font-size: 28px; font-weight: bold; color: #4CAF50; text-align: center; margin: 20px 0; }
      .footer { text-align: center; font-size: 12px; color: #777777; margin-top: 20px; }
    </style>
  </head>
  <body>
    <div class="container">
      <div class="header">
        <h1>SIP Portal</h1>
      </div>
      <div class="content">
        <p>Hello,</p>
        <p>Your One Time Password (OTP) is:</p>
        <p class="otp">` + otp + `</p>
        <p>This OTP will expire in 10 minutes. Please use it promptly.</p>
        <p>If you did not request this, please ignore this email.</p>
      </div>
      <div class="footer">
        &copy; 2025 SIP Portal. All rights reserved.
      </div>
    </div>
  </body>
</html>
`
	body := fromHeader + toHeader + subject + mimeHeader + htmlMessage

	// Send the email.
	err := smtp.SendMail(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		auth,
		from,
		[]string{email},
		[]byte(body),
	)

	return err
}
