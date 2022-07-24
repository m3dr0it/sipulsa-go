package requests

type UserOtpValidation struct {
	NewUserTemp
	Otp string `json:"otp"`
}
