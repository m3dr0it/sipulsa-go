package registration

import (
	"errors"
	"sipulsa-be/models/requests"
	"sipulsa-be/repositories"
	"time"
)

func verify(ut requests.UserOtpValidation) error {
	user, err := repositories.IsExistsByUsernameEmailPhoneOtp(
		ut.Username,
		ut.Email,
		ut.PhoneNumber,
		ut.Otp)

	if err != nil {
		return err
	}

	if time.Now().After(user.OtpExpired) {
		return errors.New("OTP expired")
	}

	return nil
}
