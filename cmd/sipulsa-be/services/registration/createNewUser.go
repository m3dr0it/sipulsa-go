package registration

import (
	"log"
	models "sipulsa-be/models/pg"
	"sipulsa-be/models/requests"
	"sipulsa-be/utilities/generator"
)

func createNewUser(ut requests.UserOtpValidation) {
	u := models.User{
		Username:     ut.Username,
		Name:         ut.Name,
		PhoneNumber:  ut.PhoneNumber,
		Email:        ut.Email,
		MemberCode:   generator.GenerateMemberCode(),
		ReferralCode: generator.GenerateReferralCode(),
	}

	log.Println(u)
}
