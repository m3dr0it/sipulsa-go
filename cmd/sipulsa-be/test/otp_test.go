package test

import (
	"log"
	"sipulsa-be/utilities/otp"
	"testing"
	"time"
)

func TestGenerateOtp(t *testing.T) {
	var otps = []string{}

	for i := 0; i < 10; i++ {
		otp, otpExpired := otp.GenerateOtp()
		if checkSliceHasIt(otps, otp) {
			panic("THERE IS SAME OTP")
		}
		if time.Now().After(otpExpired) {
			t.Fail()
		}
		otps = append(otps, otp)
	}

	log.Println(otps)
}

func checkSliceHasIt(sl []string, s string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false
}
