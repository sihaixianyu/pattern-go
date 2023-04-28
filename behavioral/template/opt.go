package template

type OTP interface {
	genRandomOTP(int) string
	saveOTP(string)
	getMsg(string) string
	sendNotification(string) error
}

type OTPHandler struct {
	opt OTP
}

func (o *OTPHandler) Process(optLen int) error {
	passwd := o.opt.genRandomOTP(int(optLen))
	o.opt.saveOTP(passwd)
	msg := o.opt.getMsg(passwd)
	err := o.opt.sendNotification(msg)
	if err != nil {
		return err
	}

	return nil
}
