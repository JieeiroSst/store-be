package otp

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jltorresm/otpgo"
	"github.com/jltorresm/otpgo/config"
)

type otp struct {
	serect string
}

type OTP interface {
	generate(username string) otpgo.TOTP
	CreateOtpByUser(username string) (string, error)
	Authorize(otp string, username string) error
}

func NewOtp(serect string) OTP {
	return &otp{
		serect: serect,
	}
}

func (o *otp) generate(username string) otpgo.TOTP {
	serect := fmt.Sprintf("%s%s", o.serect, strings.ToUpper(username))
	return otpgo.TOTP{
		Key:       serect,
		Period:    30,
		Delay:     1,
		Algorithm: config.HmacSHA1,
		Length:    6,
	}
}

func (o *otp) CreateOtpByUser(username string) (string, error) {
	totp := o.generate(username)
	token, err := totp.Generate()
	if err != nil {
		return "There was an error in sending the OTP. Please enter a valid email id or contact site Admin", err
	}
	return token, nil
}

func (o *otp) Authorize(otp string, username string) error {
	totp := o.generate(username)
	ok, err := totp.Validate(otp)
	if err != nil {
		return errors.New("There was an error in sending the OTP. Please enter a valid email id or contact site Admin")
	}
	if !ok {
		return errors.New("There was an error in sending the OTP. Please enter a valid email id or contact site Admin")
	}
	return nil
}
