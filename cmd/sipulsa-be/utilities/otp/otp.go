package otp

import (
	"crypto/rand"
	"io"
)

func GenerateOtp() string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	max := 6
	otp := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, otp, max)
	if n != max {
		panic(err)
	}

	for i := 0; i < len(otp); i++ {
		otp[i] = table[int(otp[i])%len(table)]
	}

	return string(otp)

}
