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
	to := "To: " + email + "\n"
	fromHeader := "From: " + from + "\n"
	message := "Your OTP is " + otp + "\n"
	body := fromHeader + to + subject + "\n" + message

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
