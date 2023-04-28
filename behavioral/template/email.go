package template

import "fmt"

type EmailOPT struct{}

func (s *EmailOPT) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *EmailOPT) saveOTP(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *EmailOPT) getMsg(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *EmailOPT) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}
