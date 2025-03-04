package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
	"net/smtp"
	"os"
	"strings"
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
	subject := "Subject: [E-Cell][SIP] noreply:OTP for Verification\n"
	toHeader := "To: " + email + "\n"
	fromHeader := "From: " + from + "\n"
	mimeHeader := "MIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n"
	imagePath := "images/E-Cell_logo.png"
	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		fmt.Println("Error reading image file:", err)
		return err
	}
	imageBase64 := base64.StdEncoding.EncodeToString(imageBytes)
	htmlMessageBytes, err := os.ReadFile("mail.html")
	if err != nil {
		return err
	}
	htmlMessage := string(htmlMessageBytes)
	fmt.Println(imageBase64)
	htmlMessage = strings.Replace(htmlMessage, "{{OTP}}", otp, -1)
	htmlMessage = strings.Replace(htmlMessage, "{{IMAGE}}", imageBase64, -1)
	body := fromHeader + toHeader + subject + mimeHeader + htmlMessage
	fmt.Println(htmlMessage)
	// Send the email.
	err = smtp.SendMail(
		os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"),
		auth,
		from,
		[]string{email},
		[]byte(body),
	)

	return err
}
