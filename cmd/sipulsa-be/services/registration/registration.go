package registration

import (
	"errors"
	models "sipulsa-be/models/pg"
	"sipulsa-be/models/requests"
	"sipulsa-be/repositories"
	mailservice "sipulsa-be/utilities/mail"
	"sipulsa-be/utilities/otp"
	"time"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context, ut *requests.NewUserTemp) error {

	isExistsByUsername, err := repositories.IsExistsByUsername(ut.Username)

	if isExistsByUsername {
		return errors.New(err.Error())
	}

	isExistsByPhoneNumber, err := repositories.IsExistsByPhoneNumber(ut.PhoneNumber)

	if isExistsByPhoneNumber {
		return errors.New(err.Error())
	}

	isExistsByEmail, err := repositories.IsExistsByEmail(ut.Email)

	if isExistsByEmail {
		return errors.New(err.Error())
	}

	otp, otpExpired := otp.GenerateOtp()

	userTemp := models.UserTemp{
		Name:                     ut.Name,
		Username:                 ut.Username,
		Email:                    ut.Email,
		Otp:                      otp,
		OtpExpired:               otpExpired,
		CreatedAt:                time.Now(),
		RegisteredByReferralCode: ut.RegisteredByReferralCode,
		PhoneNumber:              ut.PhoneNumber,
		CreatedBy:                "System",
	}

	var m = mailservice.MailTemplate{
		MailTo:  []string{ut.Email},
		MailCc:  []string{},
		Subject: "Test",
		Body:    otp + " expired pada " + otpExpired.String(),
	}

	go mailservice.Mail(m)

	time.Sleep(time.Millisecond * 100)

	errSave := repositories.AddUserTemp(userTemp)

	if errSave != nil {
		return errSave
	}

	return nil
}
