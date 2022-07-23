package main

import (
	"log"
	"sipulsa-be/utilities/otp"
	"testing"
)

func TestGenerateOtp(t *testing.T) {
	var otps = []string{}

	for i := 0; i < 10; i++ {
		otp := otp.GenerateOtp()
		if checkSliceHasIt(otps, otp) {
			panic("THERE IS SAME OTP")
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
