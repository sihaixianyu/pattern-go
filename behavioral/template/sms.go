package template

import "fmt"

type SmsOPT struct{}

func (s *SmsOPT) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)

	return randomOTP
}

func (s *SmsOPT) saveOTP(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *SmsOPT) getMsg(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *SmsOPT) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)

	return nil
}
