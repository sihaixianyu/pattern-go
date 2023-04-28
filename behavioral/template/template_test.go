package template

import (
	"log"
	"testing"
)

func TestTemplate(t *testing.T) {
	smsOPT := &SmsOPT{}
	handler := &OTPHandler{opt: smsOPT}
	if err := handler.Process(7); err != nil {
		log.Fatal(err)
	}

	emailOPT := &EmailOPT{}
	handler.opt = emailOPT
	if err := handler.Process(7); err != nil {
		log.Fatal(err)
	}

}
